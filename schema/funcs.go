package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entify/consts"
	"rxdrag.com/entify/entity"
)

func contextValues(p graphql.ResolveParams) entity.ContextValues {
	values := p.Context.Value(consts.CONTEXT_VALUES)
	if values == nil {
		panic("Not set token and me values in context")
	}

	return values.(entity.ContextValues)
}
