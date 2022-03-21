package schema

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"rxdrag.com/entity-engine/handler"

	"github.com/graphql-go/graphql"
)

type Feed struct {
	ID string `graphql:"id"`
}

var FeedType = graphql.NewObject(graphql.ObjectConfig{
	Name: "FeedType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
		},
	},
})

var RootSubscription = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootSubscription",
	Fields: graphql.Fields{
		"feed": &graphql.Field{
			Type: FeedType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return p.Source, nil
			},
			Subscribe: func(p graphql.ResolveParams) (interface{}, error) {
				c := make(chan interface{})

				go func() {
					var i int

					for {
						i++

						feed := Feed{ID: fmt.Sprintf("%d", i)}

						select {
						case <-p.Context.Done():
							log.Println("[RootSubscription] [Subscribe] subscription canceled")
							close(c)
							return
						default:
							c <- feed
						}

						time.Sleep(250 * time.Millisecond)

						if i == 21 {
							close(c)
							return
						}
					}
				}()

				return c, nil
			},
		},
	},
})

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"ping": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "ok", nil
			},
		},
	},
})

var schema graphql.Schema

func test() {
	schemaConfig := graphql.SchemaConfig{
		Query:        RootQuery,
		Subscription: RootSubscription,
	}

	s, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatal(err)
	}

	schema = s

	h := handler.New(&handler.Config{
		SchemaResolveFn: func() *graphql.Schema { return &s },
		Pretty:          true,
	})

	http.Handle("/graphql", h)
	http.HandleFunc("/subscriptions", SubscriptionsHandler)

	log.Println("GraphQL Server running on [POST]: localhost:8081/graphql")
	log.Println("GraphQL Playground running on [GET]: localhost:8081/graphql")

	log.Fatal(http.ListenAndServe(":8081", nil))
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	Subprotocols: []string{"graphql-ws"},
}

type ConnectionACKMessage struct {
	OperationID string `json:"id,omitempty"`
	Type        string `json:"type"`
	Payload     struct {
		Query string `json:"query"`
	} `json:"payload,omitempty"`
}

func SubscriptionsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("failed to do websocket upgrade: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	connectionACK, err := json.Marshal(map[string]string{
		"type": "connection_ack",
	})
	if err != nil {
		log.Printf("failed to marshal ws connection ack: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := conn.WriteMessage(websocket.TextMessage, connectionACK); err != nil {
		log.Printf("failed to write to ws connection: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	go handleSubscription(conn)
}

func handleSubscription(conn *websocket.Conn) {
	var subscriber *Subscriber
	subscriptionCtx, subscriptionCancelFn := context.WithCancel(context.Background())

	handleClosedConnection := func() {
		log.Println("[SubscriptionsHandler] subscriber closed connection")
		unsubscribe(subscriptionCancelFn, subscriber)
		return
	}

	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Printf("failed to read websocket message: %v", err)
			return
		}

		var msg ConnectionACKMessage
		if err := json.Unmarshal(p, &msg); err != nil {
			log.Printf("failed to unmarshal websocket message: %v", err)
			continue
		}

		if msg.Type == "stop" {
			handleClosedConnection()
			return
		}

		if msg.Type == "start" {
			subscriber = subscribe(subscriptionCtx, subscriptionCancelFn, conn, msg)
		}
	}
}

type Subscriber struct {
	UUID          string
	Conn          *websocket.Conn
	RequestString string
	OperationID   string
}

var subscribers sync.Map

func subscribersSize() uint64 {
	var size uint64
	subscribers.Range(func(_, _ interface{}) bool {
		size++
		return true
	})
	return size
}

func unsubscribe(subscriptionCancelFn context.CancelFunc, subscriber *Subscriber) {
	subscriptionCancelFn()
	if subscriber != nil {
		subscriber.Conn.Close()
		subscribers.Delete(subscriber.UUID)
	}
	log.Printf("[SubscriptionsHandler] subscribers size: %+v", subscribersSize())
}

func subscribe(ctx context.Context, subscriptionCancelFn context.CancelFunc, conn *websocket.Conn, msg ConnectionACKMessage) *Subscriber {
	subscriber := &Subscriber{
		UUID:          uuid.New().String(),
		Conn:          conn,
		RequestString: msg.Payload.Query,
		OperationID:   msg.OperationID,
	}
	subscribers.Store(subscriber.UUID, &subscriber)

	log.Printf("[SubscriptionsHandler] subscribers size: %+v", subscribersSize())

	sendMessage := func(r *graphql.Result) error {
		message, err := json.Marshal(map[string]interface{}{
			"type":    "data",
			"id":      subscriber.OperationID,
			"payload": r.Data,
		})
		if err != nil {
			return err
		}

		if err := subscriber.Conn.WriteMessage(websocket.TextMessage, message); err != nil {
			return err
		}

		return nil
	}

	go func() {
		subscribeParams := graphql.Params{
			Context:       ctx,
			RequestString: msg.Payload.Query,
			Schema:        schema,
		}

		subscribeChannel := graphql.Subscribe(subscribeParams)

		for {
			select {
			case <-ctx.Done():
				log.Printf("[SubscriptionsHandler] subscription ctx done")
				return
			case r, isOpen := <-subscribeChannel:
				if !isOpen {
					log.Printf("[SubscriptionsHandler] subscription channel closed")
					unsubscribe(subscriptionCancelFn, subscriber)
					return
				}
				if err := sendMessage(r); err != nil {
					if err == websocket.ErrCloseSent {
						unsubscribe(subscriptionCancelFn, subscriber)
					}
					log.Printf("failed to send message: %v", err)
				}
			}
		}
	}()

	return subscriber
}
