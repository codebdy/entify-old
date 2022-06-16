package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entify/config"
	"rxdrag.com/entify/consts"
	"rxdrag.com/entify/model"
	"rxdrag.com/entify/model/graph"
	"rxdrag.com/entify/resolve"
	"rxdrag.com/entify/scalars"
	"rxdrag.com/entify/utils"
)

func rootQuery() *graphql.Object {
	rootQueryConfig := graphql.ObjectConfig{
		Name:   consts.ROOT_QUERY_NAME,
		Fields: queryFields(),
	}

	return graphql.NewObject(rootQueryConfig)
}

func queryFields() graphql.Fields {
	queryFields := graphql.Fields{
		consts.SERVICE: serviceField(),
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
		consts.INSTALLED: &graphql.Field{
			Type: graphql.Boolean,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				defer utils.PrintErrorStack()
				return true, nil
			},
		},
	}

	for _, intf := range model.GlobalModel.Graph.RootInterfaces() {
		appendInterfaceToQueryFields(intf, queryFields)
	}

	for _, entity := range model.GlobalModel.Graph.RootEnities() {
		appendEntityToQueryFields(entity, queryFields)
	}

	for _, service := range model.GlobalModel.Graph.RootExternals() {
		appendServiceQueryFields(service, queryFields)
	}

	if config.AuthUrl() == "" {
		appendAuthToQuery(queryFields)
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

func quryeArgs(name string) graphql.FieldConfigArgument {
	config := graphql.FieldConfigArgument{
		consts.ARG_DISTINCTON: &graphql.ArgumentConfig{
			Type: Cache.DistinctOnEnum(name),
		},
		consts.ARG_LIMIT: &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		consts.ARG_OFFSET: &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		consts.ARG_WHERE: &graphql.ArgumentConfig{
			Type: Cache.WhereExp(name),
		},
	}
	orderByExp := Cache.OrderByExp(name)

	if orderByExp != nil {
		config[consts.ARG_ORDERBY] = &graphql.ArgumentConfig{
			Type: orderByExp,
		}
	}
	return config
}

func appendInterfaceToQueryFields(intf *graph.Interface, fields graphql.Fields) {
	(fields)[intf.QueryName()] = &graphql.Field{
		Type:    queryResponseType(intf),
		Args:    quryeArgs(intf.Name()),
		Resolve: resolve.QueryInterfaceResolveFn(intf),
	}
	(fields)[intf.QueryOneName()] = &graphql.Field{
		Type:    Cache.OutputType(intf.Name()),
		Args:    quryeArgs(intf.Name()),
		Resolve: resolve.QueryOneInterfaceResolveFn(intf),
	}

	(fields)[intf.QueryAggregateName()] = &graphql.Field{
		Type:    *AggregateType(intf),
		Args:    quryeArgs(intf.Name()),
		Resolve: resolve.QueryInterfaceResolveFn(intf),
	}
}

func appendEntityToQueryFields(entity *graph.Entity, fields graphql.Fields) {
	(fields)[entity.QueryName()] = &graphql.Field{
		Type:    queryResponseType(entity),
		Args:    quryeArgs(entity.Name()),
		Resolve: resolve.QueryEntityResolveFn(entity),
	}
	(fields)[entity.QueryOneName()] = &graphql.Field{
		Type:    Cache.OutputType(entity.Name()),
		Args:    quryeArgs(entity.Name()),
		Resolve: resolve.QueryOneEntityResolveFn(entity),
	}

	(fields)[entity.QueryAggregateName()] = &graphql.Field{
		Type:    *AggregateType(entity),
		Args:    quryeArgs(entity.Name()),
		Resolve: resolve.QueryEntityResolveFn(entity),
	}
}

func appendAuthToQuery(fields graphql.Fields) {
	fields[consts.ME] = &graphql.Field{
		Type:    baseUserType,
		Resolve: resolve.Me,
	}
}
