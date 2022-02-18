package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"rxdrag.com/entity-engine/authentication"
	"rxdrag.com/entity-engine/authentication/jwt"
	"rxdrag.com/entity-engine/schema"
)

func main() {
	// metaFields := graphql.Fields{
	// 	"id": &graphql.Field{
	// 		Type: graphql.String,
	// 		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
	// 			fmt.Println(p.Context.Value("data"))
	// 			return "world", nil
	// 		},
	// 	},
	// 	"metasContent": &graphql.Field{
	// 		Type: graphql.String,
	// 		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
	// 			return "world2", nil
	// 		},
	// 	},
	// }

	// metaType := graphql.NewObject(graphql.ObjectConfig{Name: "Meta", Fields: metaFields})
	// metaDistinctType := graphql.NewEnum(graphql.EnumConfig{
	// 	Name: "MetaDistinctExp",
	// 	Values: graphql.EnumValueConfigMap{
	// 		"name": &graphql.EnumValueConfig{
	// 			Value: "name",
	// 		},
	// 	},
	// })

	// Schema
	queryFields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				fmt.Println(p.Context.Value("data"))
				return "world", nil
			},
		},
		"hello2": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world2", nil
			},
		},
		// "_meta": &graphql.Field{
		// 	Type: graphql.NewList(metaType),
		// 	Args: graphql.FieldConfigArgument{
		// 		"distinctOn": &graphql.ArgumentConfig{
		// 			Type: metaDistinctType,
		// 		},
		// 		"limit": &graphql.ArgumentConfig{
		// 			Type: graphql.Int,
		// 		},
		// 		"offset": &graphql.ArgumentConfig{
		// 			Type: graphql.Int,
		// 		},
		// 		"orderBy": &graphql.ArgumentConfig{
		// 			Type: graphql.String,
		// 		},
		// 		"where": &graphql.ArgumentConfig{
		// 			Type: metaBoolExp,
		// 		},
		// 	},
		// 	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		// 		return "world2", nil
		// 	},
		// },
	}

	schema.MetaEntity.AppendToQueryFields(&queryFields)

	mutationFields := graphql.Fields{
		"login": &graphql.Field{
			Type: graphql.String,
			Args: graphql.FieldConfigArgument{
				"loginName": &graphql.ArgumentConfig{
					Type: &graphql.NonNull{OfType: graphql.String},
				},
				"password": &graphql.ArgumentConfig{
					Type: &graphql.NonNull{OfType: graphql.String},
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				loginName, err := authentication.Login(p.Args["loginName"].(string), p.Args["password"].(string))
				if err != nil {
					return "", err
				}
				return jwt.GenerateToken(loginName)
			},
		},
		"logout": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world2", nil
			},
		},
	}

	schema.MetaEntity.AppendToMutationFields(&mutationFields)

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: queryFields}
	rootMutation := graphql.ObjectConfig{Name: "RootMutation", Fields: mutationFields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery), Mutation: graphql.NewObject(rootMutation)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	// Query
	// query := `
	// 	{
	// 		hello
	// 	}
	// `
	// params := graphql.Params{Schema: schema, RequestString: query}
	// r := graphql.Do(params)
	// if len(r.Errors) > 0 {
	// 	log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	// }
	// rJSON, _ := json.Marshal(r)
	// fmt.Printf("%s \n", rJSON) // {"data":{"hello":"world"}}

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", authentication.CorsMiddleware(authentication.AuthMiddleware(h)))
	fmt.Println("Running a GraphQL API server at http://localhost:8080/graphql")
	err2 := http.ListenAndServe(":8080", nil)
	if err2 != nil {
		fmt.Printf("启动失败:%s", err2)
	}

}
