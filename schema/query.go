package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/config"
	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/model"
	"rxdrag.com/entity-engine/model/graph"
	"rxdrag.com/entity-engine/resolve"
	"rxdrag.com/entity-engine/utils"
)

var serviceType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: utils.FirstUpper(consts.SERVICE),
		Fields: graphql.Fields{
			consts.ID: &graphql.Field{
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
					consts.ID: config.SERVICE_ID,
				}, nil
			},
		},
		consts.NODE: &graphql.Field{
			Type: NodeInterfaceType,
			Args: graphql.FieldConfigArgument{
				consts.ID: &graphql.ArgumentConfig{
					Type: graphql.ID,
				},
			},
		},
	}
	for _, intf := range model.GlobalModel.Graph.Interfaces {
		appendToQueryFields(intf, &queryFields)
	}

	for _, entity := range model.GlobalModel.Graph.Entities {
		appendToQueryFields(entity, &queryFields)
	}

	rootQueryConfig := graphql.ObjectConfig{Name: consts.ROOT_QUERY_NAME, Fields: queryFields}

	return graphql.NewObject(rootQueryConfig)
}

func queryResponseType(node graph.Node) graphql.Output {
	return &graphql.NonNull{
		OfType: &graphql.List{
			OfType: Cache.OutputType(node.Name()),
		},
	}
}

func quryeArgs(node graph.Node) graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		consts.ARG_DISTINCTON: &graphql.ArgumentConfig{
			Type: Cache.DistinctOnEnum(node.Name()),
		},
		consts.ARG_LIMIT: &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		consts.ARG_OFFSET: &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		consts.ARG_ORDERBY: &graphql.ArgumentConfig{
			Type: Cache.OrderByExp(node.Name()),
		},
		consts.ARG_WHERE: &graphql.ArgumentConfig{
			Type: Cache.WhereExp(node.Name()),
		},
	}
}

func appendToQueryFields(node graph.Node, fields *graphql.Fields) {
	(*fields)[utils.FirstLower(node.Name())] = &graphql.Field{
		Type:    queryResponseType(node),
		Args:    quryeArgs(node),
		Resolve: resolve.QueryResolveFn(node),
	}
	(*fields)[consts.ONE+node.Name()] = &graphql.Field{
		Type:    Cache.OutputType(node.Name()),
		Args:    quryeArgs(node),
		Resolve: resolve.QueryOneResolveFn(node),
	}

	(*fields)[utils.FirstLower(node.Name())+utils.FirstUpper(consts.AGGREGATE)] = &graphql.Field{
		Type:    *AggregateType(node),
		Args:    quryeArgs(node),
		Resolve: resolve.QueryResolveFn(node),
	}
}
