package comparisons

import (
	"github.com/graphql-go/graphql"
)

const (
	STRING_COMPARISONEXP string = "StringComparisonExp"
)

var StringComparisonExp = graphql.InputObjectFieldConfig{
	Type: graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name: STRING_COMPARISONEXP,
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
			},
		},
	),
}
