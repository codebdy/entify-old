package comparison

import (
	"github.com/graphql-go/graphql"
)

var DateTimeComparisonExp = graphql.InputObjectFieldConfig{
	Type: graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name: "DateTimeComparisonExp",
			Fields: graphql.InputObjectConfigFieldMap{
				"eq": &graphql.InputObjectFieldConfig{
					Type: graphql.DateTime,
				},
				"gt": &graphql.InputObjectFieldConfig{
					Type: graphql.DateTime,
				},
				"gte": &graphql.InputObjectFieldConfig{
					Type: graphql.DateTime,
				},
				"in": &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(graphql.DateTime),
				},
				"isNull": &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				"lt": &graphql.InputObjectFieldConfig{
					Type: graphql.DateTime,
				},
				"lte": &graphql.InputObjectFieldConfig{
					Type: graphql.DateTime,
				},
				"notEq": &graphql.InputObjectFieldConfig{
					Type: graphql.DateTime,
				},
				"notIn": &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(graphql.DateTime),
				},
			},
		},
	),
}
