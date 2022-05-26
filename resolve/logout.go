package resolve

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entify/authentication"
	"rxdrag.com/entify/utils"
)

func Logout(p graphql.ResolveParams) (interface{}, error) {
	defer utils.PrintErrorStack()
	token := ContextValues(p).Token
	if token != "" {
		authentication.Logout(token)
	}
	return true, nil
}
