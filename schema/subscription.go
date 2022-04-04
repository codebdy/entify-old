package schema

import (
	"fmt"
	"log"
	"time"

	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/model/graph"
	"rxdrag.com/entity-engine/utils"

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

func RootSubscription() *graphql.Object {
	subscriptionFields := graphql.Fields{"feed": &graphql.Field{
		Type: FeedType,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			fmt.Println("Resolve", p.Source)
			return p.Source, nil
		},
		Subscribe: func(p graphql.ResolveParams) (interface{}, error) {
			fmt.Println("Subscribe it")
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
	}}

	for _, entity := range Model.Graph.Entities {
		appendToSubscriptionFields(entity, &subscriptionFields)
	}

	return graphql.NewObject(graphql.ObjectConfig{
		Name:   consts.ROOT_SUBSCRIPTION_NAME,
		Fields: subscriptionFields,
	})

}

func appendToSubscriptionFields(node graph.Node, fields *graphql.Fields) {

	(*fields)[utils.FirstLower(node.Name())] = &graphql.Field{
		Type: queryResponseType(node),
		Args: quryeArgs(node),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return p.Source, nil
		},
	}
	(*fields)[consts.ONE+node.Name()] = &graphql.Field{
		Type: Cache.OutputType(node.Name()),
		Args: quryeArgs(node),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return p.Source, nil
		},
	}

	(*fields)[utils.FirstLower(node.Name())+utils.FirstUpper(consts.AGGREGATE)] = &graphql.Field{
		Type: *AggregateType(node),
		Args: quryeArgs(node),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return p.Source, nil
		},
	}
}
