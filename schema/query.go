package schema

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/utils"
)

const (
	DISTINCTEXP string = "DistinctExp"
)

func (entity *EntityMeta) AppendToQueryFields(feilds *graphql.Fields) {
	metaDistinctType := graphql.NewEnum(
		graphql.EnumConfig{
			Name: entity.Name + DISTINCTEXP,
			Values: graphql.EnumValueConfigMap{
				"name": &graphql.EnumValueConfig{
					Value: "name",
				},
			},
		},
	)

	(*feilds)[utils.FirstLower(entity.Name)] = &graphql.Field{
		Type: entity.toOutputType(),
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
				Type: entity.toOrderBy(),
			},
			"where": &graphql.ArgumentConfig{
				Type: entity.toWhereExp(),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			fmt.Println(p.Context.Value("data"))
			return "world", nil
		},
	}
}
