package schema

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/model"
	"rxdrag.com/entity-engine/utils"
)

func AppendEntityToQueryFields(entity model.EntityMeta, feilds *graphql.Fields) {
	(*feilds)[utils.FirstLower(entity.Name)] = &graphql.Field{
		Type: graphql.String,
		// 	Args: graphql.FieldConfigArgument{
		// 		"distinctOn": &graphql.ArgumentConfig{
		// 			Type: metaDistinctType,
		// 		},
		// 		"limit": &graphql.ArgumentConfig{
		// 			Type: graphql.Int,
		// 		},
		// 		"offset": &graphql.ArgumentConfig{
		// 			Type: graphql.Int,
		// 		},
		// 		"orderBy": &graphql.ArgumentConfig{
		// 			Type: graphql.String,
		// 		},
		// 		"where": &graphql.ArgumentConfig{
		// 			Type: metaBoolExp,
		// 		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			fmt.Println(p.Context.Value("data"))
			return "world", nil
		},
	}
}
