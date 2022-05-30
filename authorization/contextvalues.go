package authorization

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entify/common"
	"rxdrag.com/entify/consts"
)

type ContextValues struct {
	Token           string
	Me              *common.User
	AbilityVerifier *AbilityVerifier
}

func ParseContextValues(p graphql.ResolveParams) ContextValues {
	values := p.Context.Value(consts.CONTEXT_VALUES)
	if values == nil {
		panic("Not set CONTEXT_VALUES in context")
	}

	return values.(ContextValues)
}

func ParseAbilityVerifier(p graphql.ResolveParams) *AbilityVerifier {
	return ParseContextValues(p).AbilityVerifier
}
