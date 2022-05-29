package resolve

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entify/authorization"
	"rxdrag.com/entify/utils"
)

func Me(p graphql.ResolveParams) (interface{}, error) {
	defer utils.PrintErrorStack()
	return authorization.ParseContextValues(p).Me, nil
}
