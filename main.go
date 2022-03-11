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
	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/meta"
	"rxdrag.com/entity-engine/migration"
	"rxdrag.com/entity-engine/repository"
	"rxdrag.com/entity-engine/schema"
)

func main() {
	queryFields := graphql.Fields{}

	for _, entity := range *repository.Entities {
		schema.AppendToQueryFields(entity, &queryFields)
	}

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
		"_publish": &graphql.Field{
			Type:    schema.OutputType(&meta.MetaEntity),
			Resolve: migration.PublishMetaResolve,
		},
		"_rollbackPublish": &graphql.Field{
			Type:    schema.OutputType(&meta.MetaEntity),
			Resolve: migration.SyncMetaResolve,
		},
		"_syncMeta": &graphql.Field{
			Type:    schema.OutputType(&meta.MetaEntity),
			Resolve: migration.SyncMetaResolve,
		},
	}

	for _, entity := range *repository.Entities {
		schema.AppendToMutationFields(entity, &mutationFields)
	}

	rootQuery := graphql.ObjectConfig{Name: consts.DefaultRootQueryName, Fields: queryFields}
	rootMutation := graphql.ObjectConfig{Name: consts.DefaultRootMutationName, Fields: mutationFields}
	schemaConfig := graphql.SchemaConfig{
		Query:    graphql.NewObject(rootQuery),
		Mutation: graphql.NewObject(rootMutation),
		Directives: []*graphql.Directive{
			graphql.NewDirective(graphql.DirectiveConfig{
				Name:      "testDirective",
				Locations: []string{graphql.DirectiveLocationField},
			}),
		},
	}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

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
