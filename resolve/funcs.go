package resolve

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entify/consts"
	"rxdrag.com/entify/entity"
)

func ContextValues(p graphql.ResolveParams) entity.ContextValues {
	values := p.Context.Value(consts.CONTEXT_VALUES)
	if values == nil {
		panic("Not set CONTEXT_VALUES in context")
	}

	return values.(entity.ContextValues)
}
