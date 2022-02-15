package schema

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/model"
	"rxdrag.com/entity-engine/utils"
)

func AppendEntityToQueryFields(entity model.EntityMeta, feilds *graphql.Fields) {
	metaFields := graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				fmt.Println(p.Context.Value("data"))
				return "world", nil
			},
		},
		"metasContent": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world2", nil
			},
		},
	}
	metaType := graphql.NewObject(graphql.ObjectConfig{Name: "Meta", Fields: metaFields})
	metaDistinctType := graphql.NewEnum(graphql.EnumConfig{
		Name: "MetaDistinctExp",
		Values: graphql.EnumValueConfigMap{
			"name": &graphql.EnumValueConfig{
				Value: "name",
			},
		},
	})

	andExp := graphql.InputObjectFieldConfig{}
	notExp := graphql.InputObjectFieldConfig{}
	orExp := graphql.InputObjectFieldConfig{}
	intComparisonExp := graphql.InputObjectFieldConfig{
		Type: graphql.NewInputObject(
			graphql.InputObjectConfig{
				Name: "IntComparisonExp",
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

	metaBoolExp := graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name: "MetaBoolExp",
			Fields: graphql.InputObjectConfigFieldMap{
				"and": &andExp,
				"not": &notExp,
				"or":  &orExp,
				"id":  &intComparisonExp,
			},
		},
	)
	andExp.Type = metaBoolExp
	notExp.Type = metaBoolExp
	orExp.Type = metaBoolExp

	(*feilds)[utils.FirstLower(entity.Name)] = &graphql.Field{
		Type: metaType,
		Args: graphql.FieldConfigArgument{
			"distinctOn": &graphql.ArgumentConfig{
				Type: metaDistinctType,
			},
			"limit": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"offset": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"orderBy": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"where": &graphql.ArgumentConfig{
				Type: metaBoolExp,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			fmt.Println(p.Context.Value("data"))
			return "world", nil
		},
	}
}
