package comparison

import (
	"github.com/graphql-go/graphql"
)

var IntComparisonExp = graphql.InputObjectFieldConfig{
	Type: graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name: "IntComparisonExp",
			Fields: graphql.InputObjectConfigFieldMap{
				"eq": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				"gt": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				"gte": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				"in": &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(graphql.Int),
				},
				"isNull": &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				"lt": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				"lte": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				"notEq": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				"notIn": &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(graphql.Int),
				},
			},
		},
	),
}
