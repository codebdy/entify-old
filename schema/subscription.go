package schema

import (
	"fmt"
	"log"
	"time"

	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/meta"
	"rxdrag.com/entity-engine/resolve"
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

	for _, entity := range meta.Metas.Entities {
		appendToSubscriptionFields(&entity, &subscriptionFields)
	}

	return graphql.NewObject(graphql.ObjectConfig{
		Name:   consts.ROOT_SUBSCRIPTION_NAME,
		Fields: subscriptionFields,
	})

}

func appendToSubscriptionFields(entity *meta.Entity, fields *graphql.Fields) {
	//如果是枚举
	if entity.EntityType == meta.ENTITY_ENUM {
		return
	}

	(*fields)[utils.FirstLower(entity.Name)] = &graphql.Field{
		Type: &graphql.NonNull{
			OfType: &graphql.List{
				OfType: Cache.OutputType(entity),
			},
		},
		Args: graphql.FieldConfigArgument{
			consts.ARG_DISTINCTON: &graphql.ArgumentConfig{
				Type: Cache.DistinctOnEnum(entity),
			},
			consts.ARG_LIMIT: &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			consts.ARG_OFFSET: &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			consts.ARG_ORDERBY: &graphql.ArgumentConfig{
				Type: Cache.OrderByExp(entity),
			},
			consts.ARG_WHERE: &graphql.ArgumentConfig{
				Type: Cache.WhereExp(entity),
			},
		},
		Resolve: resolve.QueryResolveFn(entity),
	}
	(*fields)[consts.ONE+entity.Name] = &graphql.Field{
		Type: Cache.OutputType(entity),
		Args: graphql.FieldConfigArgument{
			consts.ARG_DISTINCTON: &graphql.ArgumentConfig{
				Type: Cache.DistinctOnEnum(entity),
			},
			consts.ARG_OFFSET: &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			consts.ARG_ORDERBY: &graphql.ArgumentConfig{
				Type: Cache.OrderByExp(entity),
			},
			consts.ARG_WHERE: &graphql.ArgumentConfig{
				Type: Cache.WhereExp(entity),
			},
		},
		Resolve: resolve.QueryOneResolveFn(entity),
	}

	(*fields)[utils.FirstLower(entity.Name)+utils.FirstUpper(consts.AGGREGATE)] = &graphql.Field{
		Type: *AggregateType(entity, []*meta.Entity{}),
		Args: graphql.FieldConfigArgument{
			consts.ARG_DISTINCTON: &graphql.ArgumentConfig{
				Type: Cache.DistinctOnEnum(entity),
			},
			consts.ARG_LIMIT: &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			consts.ARG_OFFSET: &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			consts.ARG_ORDERBY: &graphql.ArgumentConfig{
				Type: Cache.OrderByExp(entity),
			},
			consts.ARG_WHERE: &graphql.ArgumentConfig{
				Type: Cache.WhereExp(entity),
			},
		},
		Resolve: resolve.QueryResolveFn(entity),
	}
}
