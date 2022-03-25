package schema

import (
	"fmt"
	"log"
	"time"

	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/meta"
	"rxdrag.com/entity-engine/model"
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

	for _, entity := range model.TheModel.Entities {
		appendToSubscriptionFields(entity, &subscriptionFields)
	}

	return graphql.NewObject(graphql.ObjectConfig{
		Name:   consts.ROOT_SUBSCRIPTION_NAME,
		Fields: subscriptionFields,
	})

}

func appendToSubscriptionFields(entity *model.Entity, fields *graphql.Fields) {
	//如果是枚举
	if entity.EntityType == meta.ENTITY_ENUM {
		return
	}

	(*fields)[utils.FirstLower(entity.Name)] = &graphql.Field{
		Type: queryResponseType(entity),
		Args: quryeArgs(entity),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return p.Source, nil
		},
	}
	(*fields)[consts.ONE+entity.Name] = &graphql.Field{
		Type: Cache.OutputType(entity.Name),
		Args: quryeArgs(entity),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return p.Source, nil
		},
	}

	(*fields)[utils.FirstLower(entity.Name)+utils.FirstUpper(consts.AGGREGATE)] = &graphql.Field{
		Type: *AggregateType(entity, []*model.Entity{}),
		Args: quryeArgs(entity),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return p.Source, nil
		},
	}
}
