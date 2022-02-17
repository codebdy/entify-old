package comparisons

import (
	"github.com/graphql-go/graphql"
)

var EnumComparisonExp = graphql.InputObjectFieldConfig{
	Type: graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name: "EnumComparisonExp",
			Fields: graphql.InputObjectConfigFieldMap{
				"eq": &graphql.InputObjectFieldConfig{
					Type: graphql.EnumValueType,
				},
				"in": &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(graphql.Boolean),
				},
				"isNull": &graphql.InputObjectFieldConfig{
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
