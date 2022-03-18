package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/meta"
	"rxdrag.com/entity-engine/resolve"
	"rxdrag.com/entity-engine/utils"
)

func AppendToQueryFields(entity *meta.Entity, fields *graphql.Fields) {
	//如果是枚举
	if entity.EntityType == meta.ENTITY_ENUM {
		return
	}

	(*fields)[consts.QUERY+entity.Name] = &graphql.Field{
		Type: &graphql.NonNull{
			OfType: &graphql.List{
				OfType: *OutputType(entity, []*meta.Entity{}),
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
		Resolve: resolve.QueryResolveFn(entity),
	}
	(*fields)[consts.ONE+entity.Name] = &graphql.Field{
		Type: *OutputType(entity, []*meta.Entity{}),
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
		Resolve: resolve.QueryOneResolveFn(entity),
	}

	(*fields)[utils.FirstLower(entity.Name)+utils.FirstUpper(consts.AGGREGATE)] = &graphql.Field{
		Type: *AggregateType(entity, []*meta.Entity{}),
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
		Resolve: resolve.QueryResolveFn(entity),
	}
}
