package repository

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/meta"
)

func PostOneResolveFn(entity *meta.Entity) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		object := p.Args[consts.ARG_OBJECT].(map[string]interface{})
		return SaveOneEntity(object, entity)
	}
}