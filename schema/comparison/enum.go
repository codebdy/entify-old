package comparison

import (
	"github.com/graphql-go/graphql"
)

var EnumComparisonExp = graphql.InputObjectFieldConfig{
	Type: graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name: "EnumComparisonExp",
			Fields: graphql.InputObjectConfigFieldMap{
				ARG_EQ: &graphql.InputObjectFieldConfig{
					Type: graphql.EnumValueType,
				},
				ARG_IN: &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(graphql.Boolean),
				},
				ARG_ISNULL: &graphql.InputObjectFieldConfig{
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
