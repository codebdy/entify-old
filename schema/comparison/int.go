package comparison

import (
	"github.com/graphql-go/graphql"
)

var IntComparisonExp = graphql.InputObjectFieldConfig{
	Type: graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name: "IntComparisonExp",
			Fields: graphql.InputObjectConfigFieldMap{
				ARG_EQ: &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				ARG_GT: &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				ARG_GTE: &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				ARG_IN: &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(graphql.Int),
				},
				ARG_ISNULL: &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				ARG_LT: &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				ARG_LTE: &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				ARG_NOTEQ: &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				ARG_NOTIN: &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(graphql.Int),
				},
			},
		},
	),
}
