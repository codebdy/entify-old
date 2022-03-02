package comparison

import (
	"github.com/graphql-go/graphql"
)

var BooleanComparisonExp = graphql.InputObjectFieldConfig{
	Type: graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name: "BooleanComparisonExp",
			Fields: graphql.InputObjectConfigFieldMap{
				ARG_EQ: &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				ARG_GT: &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				ARG_GTE: &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				ARG_IN: &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(graphql.Boolean),
				},
				ARG_ISNULL: &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				ARG_LT: &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				ARG_LTE: &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				ARG_NOTEQ: &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				ARG_NOTIN: &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(graphql.Boolean),
				},
			},
		},
	),
}
