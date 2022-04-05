package resolve

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/model/graph"
	"rxdrag.com/entity-engine/repository"
)

func PostOneResolveFn(entity *graph.Entity) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		object := p.Args[consts.ARG_OBJECT].(map[string]interface{})
		return repository.SaveOne(object, entity)
	}
}
