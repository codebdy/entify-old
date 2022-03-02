package comparison

import (
	"github.com/graphql-go/graphql"
)

var DateTimeComparisonExp = graphql.InputObjectFieldConfig{
	Type: graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name: "DateTimeComparisonExp",
			Fields: graphql.InputObjectConfigFieldMap{
				ARG_EQ: &graphql.InputObjectFieldConfig{
					Type: graphql.DateTime,
				},
				ARG_GT: &graphql.InputObjectFieldConfig{
					Type: graphql.DateTime,
				},
				ARG_GTE: &graphql.InputObjectFieldConfig{
					Type: graphql.DateTime,
				},
				ARG_IN: &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(graphql.DateTime),
				},
				ARG_ISNULL: &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				ARG_LT: &graphql.InputObjectFieldConfig{
					Type: graphql.DateTime,
				},
				ARG_LTE: &graphql.InputObjectFieldConfig{
					Type: graphql.DateTime,
				},
				ARG_NOTEQ: &graphql.InputObjectFieldConfig{
					Type: graphql.DateTime,
				},
				ARG_NOTIN: &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(graphql.DateTime),
				},
			},
		},
	),
}
