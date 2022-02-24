package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/meta"
	"rxdrag.com/entity-engine/repository"
	"rxdrag.com/entity-engine/utils"
)

func AppendToQueryFields(entity *meta.Entity, feilds *graphql.Fields) {
	//如果是枚举
	if entity.EntityType == meta.Entity_ENUM {
		return
	}

	(*feilds)[utils.FirstLower(entity.Name)] = &graphql.Field{
		Type: &graphql.NonNull{
			OfType: &graphql.List{
				OfType: OutputType(entity),
			},
		},
		Args: graphql.FieldConfigArgument{
			"distinctOn": &graphql.ArgumentConfig{
				Type: DistinctOnEnum(entity),
			},
			"limit": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"offset": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"orderBy": &graphql.ArgumentConfig{
				Type: OrderBy(entity),
			},
			"where": &graphql.ArgumentConfig{
				Type: WhereExp(entity),
			},
		},
		Resolve: repository.QueryResolveFn(entity),
	}
	(*feilds)[utils.FirstLower(entity.Name)+"ById"] = &graphql.Field{
		Type: OutputType(entity),
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: repository.QueryResolveFn(entity),
	}

	(*feilds)[utils.FirstLower(entity.Name)+"Aggregate"] = &graphql.Field{
		Type: AggregateType(entity),
		Args: graphql.FieldConfigArgument{
			"distinctOn": &graphql.ArgumentConfig{
				Type: DistinctOnEnum(entity),
			},
			"limit": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"offset": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"orderBy": &graphql.ArgumentConfig{
				Type: OrderBy(entity),
			},
			"where": &graphql.ArgumentConfig{
				Type: WhereExp(entity),
			},
		},
		Resolve: repository.QueryResolveFn(entity),
	}
}
