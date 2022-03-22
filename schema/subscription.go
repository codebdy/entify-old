package schema

import (
	"fmt"
	"log"
	"time"

	"rxdrag.com/entity-engine/consts"

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
	return graphql.NewObject(graphql.ObjectConfig{
		Name: consts.ROOT_SUBSCRIPTION_NAME,
		Fields: graphql.Fields{
			"feed": &graphql.Field{
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
			},
		},
	})

}
