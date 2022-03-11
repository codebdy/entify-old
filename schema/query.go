package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/meta"
	"rxdrag.com/entity-engine/repository"
)

func AppendToQueryFields(entity *meta.Entity, feilds *graphql.Fields) {
	//如果是枚举
	if entity.EntityType == meta.Entity_ENUM {
		return
	}

	(*feilds)[consts.QUERY+entity.Name] = &graphql.Field{
		Type: &graphql.NonNull{
			OfType: &graphql.List{
				OfType: OutputType(entity),
			},
		},
		Args: graphql.FieldConfigArgument{
			consts.ARG_DISTINCTON: &graphql.ArgumentConfig{
				Type: DistinctOnEnum(entity),
			},
			consts.ARG_LIMIT: &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			consts.ARG_OFFSET: &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			consts.ARG_ORDERBY: &graphql.ArgumentConfig{
				Type: OrderBy(entity),
			},
			consts.ARG_WHERE: &graphql.ArgumentConfig{
				Type: WhereExp(entity),
			},
		},
		Resolve: repository.QueryResolveFn(entity),
	}
	(*feilds)[consts.ONE+entity.Name] = &graphql.Field{
		Type: OutputType(entity),
		Args: graphql.FieldConfigArgument{
			consts.ARG_DISTINCTON: &graphql.ArgumentConfig{
				Type: DistinctOnEnum(entity),
			},
			consts.ARG_OFFSET: &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			consts.ARG_ORDERBY: &graphql.ArgumentConfig{
				Type: OrderBy(entity),
			},
			consts.ARG_WHERE: &graphql.ArgumentConfig{
				Type: WhereExp(entity),
			},
		},
		Resolve: repository.QueryOneResolveFn(entity),
	}

	(*feilds)[consts.AGGREGATE+entity.Name] = &graphql.Field{
		Type: AggregateType(entity),
		Args: graphql.FieldConfigArgument{
			consts.ARG_DISTINCTON: &graphql.ArgumentConfig{
				Type: DistinctOnEnum(entity),
			},
			consts.ARG_LIMIT: &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			consts.ARG_OFFSET: &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			consts.ARG_ORDERBY: &graphql.ArgumentConfig{
				Type: OrderBy(entity),
			},
			consts.ARG_WHERE: &graphql.ArgumentConfig{
				Type: WhereExp(entity),
			},
		},
		Resolve: repository.QueryResolveFn(entity),
	}
}
