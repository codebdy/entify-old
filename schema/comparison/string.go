package comparison

import (
	"github.com/graphql-go/graphql"
)

var StringComparisonExp = graphql.InputObjectFieldConfig{
	Type: graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name: "StringComparisonExp",
			Fields: graphql.InputObjectConfigFieldMap{
				"eq": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"gt": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"gte": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"iLike": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"in": &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(graphql.String),
				},
				"iregex": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"isNull": &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				"like": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"lt": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"lte": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"notEq": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"notILike": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"notIn": &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(graphql.String),
				},
				"notIRegex": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"notLike": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"notRegex": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"notSimilar": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"regex": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"similar": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
			},
		},
	),
}
