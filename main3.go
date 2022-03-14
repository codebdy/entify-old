package main

import (
	"log"

	"github.com/graphql-go/graphql"
)

func main3() {
	// Schema
	// Add Resource Type
	resourceType := graphql.NewObject(graphql.ObjectConfig{
		Name: "resourceA",
		Fields: graphql.Fields{
			"uuid": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The identifier of the resource.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return "aUUID", nil
				},
			},
		},
	})
	objectType := graphql.NewObject(graphql.ObjectConfig{
		Name: "objectA",
		Fields: graphql.Fields{
			"fieldA": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return "fieldA", nil
				},
			},
		},
	})
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
		"fieldA": &graphql.Field{
			Type: resourceType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "fieldA", nil
			},
		},
		"fieldB": &graphql.Field{
			Type: objectType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "fieldB", nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	_, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}
	_, err = graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

}
