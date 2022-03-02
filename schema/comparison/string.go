package comparison

import (
	"github.com/graphql-go/graphql"
)

var StringComparisonExp = graphql.InputObjectFieldConfig{
	Type: graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name: "StringComparisonExp",
			Fields: graphql.InputObjectConfigFieldMap{
				ARG_EQ: &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				ARG_GT: &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				ARG_GTE: &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				ARG_ILIKE: &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				ARG_IN: &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(graphql.String),
				},
				ARG_IREGEX: &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				ARG_ISNULL: &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				ARG_LIKE: &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				ARG_LT: &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				ARG_LTE: &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				ARG_NOTEQ: &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				ARG_NOTILIKE: &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				ARG_NOTIN: &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(graphql.String),
				},
				ARG_NOTIREGEX: &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				ARG_NOTLIKE: &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				ARG_NOTREGEX: &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				ARG_NOTSIMILAR: &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				ARG_REGEX: &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				ARG_SIMILAR: &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
			},
		},
	),
}
