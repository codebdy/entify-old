package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entify/model/graph"
	"rxdrag.com/entify/model/meta"
	"rxdrag.com/entify/utils"
)

func appendServiceQueryFields(exClass *graph.Entity, fields graphql.Fields) {
	methods := exClass.MethodsByType(meta.QUERY)
	if len(methods) > 0 {
		(fields)[exClass.QueryName()] = &graphql.Field{
			Type: graphql.String,
			//Resolve: resolve.QueryResolveFn(node),
		}
	}

}

func appendServiceMutationFields(serviceClass *graph.Entity, fields graphql.Fields) {
	methods := serviceClass.MethodsByType(meta.MUTATION)
	if len(methods) > 0 {
		(fields)[utils.FirstLower(serviceClass.Name())] = &graphql.Field{
			Type: graphql.String,
			//Resolve: resolve.QueryResolveFn(node),
		}
	}
}
