package resolve

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/model/data"
	"rxdrag.com/entity-engine/model/graph"
	"rxdrag.com/entity-engine/repository"
	"rxdrag.com/entity-engine/utils"
)

func PostOneResolveFn(entity *graph.Entity) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		defer utils.PrintErrorStack()
		object := p.Args[consts.ARG_OBJECT].(map[string]interface{})
		ConvertId(object)
		instance := data.NewInstance(object, entity)
		return repository.SaveOne(instance)
	}
}
