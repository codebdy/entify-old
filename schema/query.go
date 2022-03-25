package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/config"
	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/model"
	"rxdrag.com/entity-engine/resolve"
	"rxdrag.com/entity-engine/utils"
)

var serviceType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: utils.FirstUpper(consts.SERVICE),
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
		},
		Description: "Micro service info",
	},
)

func rootQuery() *graphql.Object {
	queryFields := graphql.Fields{
		consts.SERVICE: &graphql.Field{
			Type: serviceType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return map[string]interface{}{
					"id": config.SERVICE_ID,
				}, nil
			},
		},
		consts.NODE: &graphql.Field{
			Type: NodeInterfaceType,
		},
	}

	for _, entity := range model.TheModel.Entities {
		appendToQueryFields(entity, &queryFields)
	}

	rootQueryConfig := graphql.ObjectConfig{Name: consts.ROOT_QUERY_NAME, Fields: queryFields}

	return graphql.NewObject(rootQueryConfig)
}

func queryResponseType(entity *model.Entity) graphql.Output {
	return &graphql.NonNull{
		OfType: &graphql.List{
			OfType: Cache.OutputObjectType(entity),
		},
	}
}

func quryeArgs(entity *model.Entity) graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		consts.ARG_DISTINCTON: &graphql.ArgumentConfig{
			Type: Cache.DistinctOnEnum(entity),
		},
		consts.ARG_LIMIT: &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		consts.ARG_OFFSET: &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		consts.ARG_ORDERBY: &graphql.ArgumentConfig{
			Type: Cache.OrderByExp(entity),
		},
		consts.ARG_WHERE: &graphql.ArgumentConfig{
			Type: Cache.WhereExp(entity),
		},
	}
}

func appendToQueryFields(entity *model.Entity, fields *graphql.Fields) {
	(*fields)[utils.FirstLower(entity.Name)] = &graphql.Field{
		Type:    queryResponseType(entity),
		Args:    quryeArgs(entity),
		Resolve: resolve.QueryResolveFn(entity),
	}
	(*fields)[consts.ONE+entity.Name] = &graphql.Field{
		Type:    Cache.OutputObjectType(entity),
		Args:    quryeArgs(entity),
		Resolve: resolve.QueryOneResolveFn(entity),
	}

	(*fields)[utils.FirstLower(entity.Name)+utils.FirstUpper(consts.AGGREGATE)] = &graphql.Field{
		Type:    *AggregateType(entity, []*model.Entity{}),
		Args:    quryeArgs(entity),
		Resolve: resolve.QueryResolveFn(entity),
	}
}
