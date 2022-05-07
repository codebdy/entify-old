package schema

import (
	"fmt"

	"rxdrag.com/entify/consts"

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
	// subscriptionFields := graphql.Fields{"feed": &graphql.Field{
	// 	Type: FeedType,
	// 	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
	// 		fmt.Println("Resolve", p.Source)
	// 		return p.Source, nil
	// 	},
	// 	Subscribe: func(p graphql.ResolveParams) (interface{}, error) {
	// 		fmt.Println("Subscribe it")
	// 		c := make(chan interface{})

	// 		go func() {
	// 			var i int
	// 			for {
	// 				i++
	// 				feed := Feed{ID: fmt.Sprintf("%d", i)}
	// 				select {
	// 				case <-p.Context.Done():
	// 					log.Println("[RootSubscription] [Subscribe] subscription canceled")
	// 					close(c)
	// 					return
	// 				default:
	// 					c <- feed
	// 				}

	// 				time.Sleep(250 * time.Millisecond)

	// 				if i == 21 {
	// 					close(c)
	// 					return
	// 				}
	// 			}
	// 		}()

	// 		return c, nil
	// 	},
	// }}

	subscriptionObj := graphql.NewObject(graphql.ObjectConfig{
		Name:   consts.ROOT_SUBSCRIPTION_NAME,
		Fields: queryFields(),
	})

	//添加订阅代码
	fields := subscriptionObj.Fields()
	for i := range fields {
		field := fields[i]
		field.Subscribe = func(p graphql.ResolveParams) (interface{}, error) {
			fmt.Println("Resolve", p.Source)
			return p.Source, nil
		}
	}

	return subscriptionObj
}

// func appendToSubscriptionFields(node graph.Node, fields *graphql.Fields) {

// 	(*fields)[utils.FirstLower(node.Name())] = &graphql.Field{
// 		Type: queryResponseType(node),
// 		Args: quryeArgs(node),
// 		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 			return p.Source, nil
// 		},
// 	}
// 	(*fields)[consts.ONE+node.Name()] = &graphql.Field{
// 		Type: Cache.OutputType(node.Name()),
// 		Args: quryeArgs(node),
// 		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 			return p.Source, nil
// 		},
// 	}

// 	(*fields)[utils.FirstLower(node.Name())+utils.FirstUpper(consts.AGGREGATE)] = &graphql.Field{
// 		Type: *AggregateType(node),
// 		Args: quryeArgs(node),
// 		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 			return p.Source, nil
// 		},
// 	}
// }
