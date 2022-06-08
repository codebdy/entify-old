package resolve

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entify/consts"
	"rxdrag.com/entify/model/data"
	"rxdrag.com/entify/model/graph"
	"rxdrag.com/entify/repository"
	"rxdrag.com/entify/utils"
)

func PostOneResolveFn(entity *graph.Entity) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		defer utils.PrintErrorStack()
		object := p.Args[consts.ARG_OBJECT].(map[string]interface{})
		ConvertId(object)
		v := makeEntityAbilityVerifier(p, entity.Uuid())
		instance := data.NewInstance(object, entity)
		return repository.SaveOne(instance, v)
	}
}
