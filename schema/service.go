package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entify/model/graph"
	"rxdrag.com/entify/model/meta"
	"rxdrag.com/entify/utils"
)

func appendServiceQueryFields(serviceClass *graph.Class, fields *graphql.Fields) {
	methods := serviceClass.MethodsByType(meta.QUERY)
	if len(methods) > 0 {
		(*fields)[utils.FirstLower(serviceClass.Name())] = &graphql.Field{
			Type: graphql.String,
			//Resolve: resolve.QueryResolveFn(node),
		}
	}

}

func appendServiceMutationFields(serviceClass *graph.Class, fields *graphql.Fields) {
	methods := serviceClass.MethodsByType(meta.MUTATION)
	if len(methods) > 0 {
		(*fields)[utils.FirstLower(serviceClass.Name())] = &graphql.Field{
			Type: graphql.String,
			//Resolve: resolve.QueryResolveFn(node),
		}
	}
}
