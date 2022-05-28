package common

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entify/consts"
)

type ContextValues struct {
	Token string
	Me    *User
}

func ParseContextValues(p graphql.ResolveParams) ContextValues {
	values := p.Context.Value(consts.CONTEXT_VALUES)
	if values == nil {
		panic("Not set CONTEXT_VALUES in context")
	}

	return values.(ContextValues)
}
