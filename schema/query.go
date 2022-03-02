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
			ARG_DISTINCTON: &graphql.ArgumentConfig{
				Type: DistinctOnEnum(entity),
			},
			ARG_LIMIT: &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			ARG_OFFSET: &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			ARG_ORDERBY: &graphql.ArgumentConfig{
				Type: OrderBy(entity),
			},
			ARG_WHERE: &graphql.ArgumentConfig{
				Type: WhereExp(entity),
			},
		},
		Resolve: repository.QueryResolveFn(entity),
	}
	(*feilds)[CONST_ONE+entity.Name] = &graphql.Field{
		Type: OutputType(entity),
		Args: graphql.FieldConfigArgument{
			ARG_DISTINCTON: &graphql.ArgumentConfig{
				Type: DistinctOnEnum(entity),
			},
			ARG_OFFSET: &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			ARG_ORDERBY: &graphql.ArgumentConfig{
				Type: OrderBy(entity),
			},
			ARG_WHERE: &graphql.ArgumentConfig{
				Type: WhereExp(entity),
			},
		},
		Resolve: repository.QueryResolveFn(entity),
	}

	(*feilds)[utils.FirstLower(entity.Name)+CONST_AGGREGATE] = &graphql.Field{
		Type: AggregateType(entity),
		Args: graphql.FieldConfigArgument{
			ARG_DISTINCTON: &graphql.ArgumentConfig{
				Type: DistinctOnEnum(entity),
			},
			ARG_LIMIT: &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			ARG_OFFSET: &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			ARG_ORDERBY: &graphql.ArgumentConfig{
				Type: OrderBy(entity),
			},
			ARG_WHERE: &graphql.ArgumentConfig{
				Type: WhereExp(entity),
			},
		},
		Resolve: repository.QueryResolveFn(entity),
	}
}
