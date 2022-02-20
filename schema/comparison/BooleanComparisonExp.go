package comparison

import (
	"github.com/graphql-go/graphql"
)

var BooleanComparisonExp = graphql.InputObjectFieldConfig{
	Type: graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name: "BooleanComparisonExp",
			Fields: graphql.InputObjectConfigFieldMap{
				"eq": &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				"gt": &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				"gte": &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				"in": &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(graphql.Boolean),
				},
				"isNull": &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				"lt": &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				"lte": &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				"notEq": &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				"notIn": &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(graphql.Boolean),
				},
			},
		},
	),
}
