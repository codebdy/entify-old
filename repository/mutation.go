package repository

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/meta"
	"rxdrag.com/entity-engine/utils"
)

func PostOneResolveFn(entity *meta.Entity) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		object := p.Args[consts.ARG_OBJECT]
		return SaveOneEntity(object.(utils.SimpleJSON), entity)
	}
}
