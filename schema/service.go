package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/model/graph"
	"rxdrag.com/entity-engine/model/meta"
	"rxdrag.com/entity-engine/utils"
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
