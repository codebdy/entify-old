package schema

import "github.com/graphql-go/graphql"

var EnumOrderBy = graphql.NewEnum(
	graphql.EnumConfig{
		Name: "OrderBy",
		Values: graphql.EnumValueConfigMap{
			"asc": &graphql.EnumValueConfig{
				Value: "asc",
			},
			"ascNullsFirst": &graphql.EnumValueConfig{
				Value: "ascNullsFirst",
			},
			"ascNullsLast": &graphql.EnumValueConfig{
				Value: "ascNullsLast",
			},
			"desc": &graphql.EnumValueConfig{
				Value: "desc",
			},
			"descNullsFirst": &graphql.EnumValueConfig{
				Value: "descNullsFirst",
			},
			"descNullsLast": &graphql.EnumValueConfig{
				Value: "descNullsLast",
			},
		},
	},
)
