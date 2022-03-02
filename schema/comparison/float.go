package comparison

import (
	"github.com/graphql-go/graphql"
)

var FloatComparisonExp = graphql.InputObjectFieldConfig{
	Type: graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name: "FloatComparisonExp",
			Fields: graphql.InputObjectConfigFieldMap{
				ARG_EQ: &graphql.InputObjectFieldConfig{
					Type: graphql.Float,
				},
				ARG_GT: &graphql.InputObjectFieldConfig{
					Type: graphql.Float,
				},
				ARG_GTE: &graphql.InputObjectFieldConfig{
					Type: graphql.Float,
				},
				ARG_IN: &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(graphql.Float),
				},
				ARG_ISNULL: &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				ARG_LT: &graphql.InputObjectFieldConfig{
					Type: graphql.Float,
				},
				ARG_LTE: &graphql.InputObjectFieldConfig{
					Type: graphql.Float,
				},
				ARG_NOTEQ: &graphql.InputObjectFieldConfig{
					Type: graphql.Float,
				},
				ARG_NOTIN: &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(graphql.Float),
				},
			},
		},
	),
}
