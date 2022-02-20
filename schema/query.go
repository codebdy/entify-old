package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/utils"
)

func (entity *EntityMeta) AppendToQueryFields(feilds *graphql.Fields) {
	//如果是枚举
	if entity.EntityType == Entity_ENUM {
		return
	}

	(*feilds)[utils.FirstLower(entity.Name)] = &graphql.Field{
		Type: &graphql.NonNull{
			OfType: &graphql.List{
				OfType: entity.toOutputType(),
			},
		},
		Args: graphql.FieldConfigArgument{
			"distinctOn": &graphql.ArgumentConfig{
				Type: entity.toDistinctOnEnum(),
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
		Resolve: entity.QueryResolve(),
	}
	(*feilds)[utils.FirstLower(entity.Name)+"ById"] = &graphql.Field{
		Type: entity.toOutputType(),
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: entity.QueryResolve(),
	}

	(*feilds)[utils.FirstLower(entity.Name)+"Aggregate"] = &graphql.Field{
		Type: entity.toOutputType(),
		Args: graphql.FieldConfigArgument{
			"distinctOn": &graphql.ArgumentConfig{
				Type: entity.toDistinctOnEnum(),
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
		Resolve: entity.QueryResolve(),
	}
}
