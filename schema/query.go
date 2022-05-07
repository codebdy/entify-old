package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/config"
	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/model"
	"rxdrag.com/entity-engine/model/graph"
	"rxdrag.com/entity-engine/resolve"
	"rxdrag.com/entity-engine/scalars"
	"rxdrag.com/entity-engine/utils"
)

func resolveTypeFn(p graphql.ResolveTypeParams) *graphql.Object {
	if value, ok := p.Value.(map[string]interface{}); ok {
		if id, ok := value[consts.ID].(uint64); ok {
			entityInnerId := utils.DecodeEntityInnerId(id)
			return Cache.GetEntityTypeByInnerId(entityInnerId)
		}
	}
	return nil
}

func rootQuery() *graphql.Object {
	rootQueryConfig := graphql.ObjectConfig{
		Name:   consts.ROOT_QUERY_NAME,
		Fields: queryFields(),
	}

	return graphql.NewObject(rootQueryConfig)
}

func queryFields() graphql.Fields {
	queryFields := graphql.Fields{
		consts.SERVICE: &graphql.Field{
			Type: _ServiceType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return map[string]interface{}{
					consts.ID:  config.SERVICE_ID,
					consts.SDL: `query{}`,
				}, nil
			},
		},
		consts.ENTITIES: &graphql.Field{
			Type: &graphql.NonNull{
				OfType: &graphql.List{
					OfType: EntityType,
				},
			},
			Args: graphql.FieldConfigArgument{
				consts.REPRESENTATIONS: &graphql.ArgumentConfig{
					Type: &graphql.NonNull{
						OfType: &graphql.List{
							OfType: &graphql.NonNull{
								OfType: scalars.AnyType,
							},
						},
					},
				},
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
	for _, intf := range model.GlobalModel.Graph.RootInterfaces() {
		appendToQueryFields(intf, &queryFields)
	}

	for _, entity := range model.GlobalModel.Graph.RootEnities() {
		appendToQueryFields(entity, &queryFields)
	}

	for _, service := range model.GlobalModel.Graph.RootServices() {
		appendServiceQueryFields(service, &queryFields)
	}

	return queryFields
}

func queryResponseType(node graph.Noder) graphql.Output {
	return &graphql.NonNull{
		OfType: &graphql.List{
			OfType: Cache.OutputType(node.Name()),
		},
	}
}

func quryeArgs(node graph.Noder) graphql.FieldConfigArgument {
	config := graphql.FieldConfigArgument{
		consts.ARG_DISTINCTON: &graphql.ArgumentConfig{
			Type: Cache.DistinctOnEnum(node.Name()),
		},
		consts.ARG_LIMIT: &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		consts.ARG_OFFSET: &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		consts.ARG_WHERE: &graphql.ArgumentConfig{
			Type: Cache.WhereExp(node.Name()),
		},
	}
	orderByExp := Cache.OrderByExp(node.Name())

	if orderByExp != nil {
		config[consts.ARG_ORDERBY] = &graphql.ArgumentConfig{
			Type: orderByExp,
		}
	}
	return config
}

func appendToQueryFields(node graph.Noder, fields *graphql.Fields) {
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
