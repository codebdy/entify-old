package comparison

import (
	"github.com/graphql-go/graphql"
)

var FloatComparisonExp = graphql.InputObjectFieldConfig{
	Type: graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name: "FloatComparisonExp",
			Fields: graphql.InputObjectConfigFieldMap{
				"eq": &graphql.InputObjectFieldConfig{
					Type: graphql.Float,
				},
				"gt": &graphql.InputObjectFieldConfig{
					Type: graphql.Float,
				},
				"gte": &graphql.InputObjectFieldConfig{
					Type: graphql.Float,
				},
				"in": &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(graphql.Float),
				},
				"isNull": &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				"lt": &graphql.InputObjectFieldConfig{
					Type: graphql.Float,
				},
				"lte": &graphql.InputObjectFieldConfig{
					Type: graphql.Float,
				},
				"notEq": &graphql.InputObjectFieldConfig{
					Type: graphql.Float,
				},
				"notIn": &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(graphql.Float),
				},
			},
		},
	),
}
