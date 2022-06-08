package resolve

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entify/authcontext"
	"rxdrag.com/entify/utils"
)

func Me(p graphql.ResolveParams) (interface{}, error) {
	defer utils.PrintErrorStack()
	return authcontext.ParseContextValues(p).Me, nil
}
