package resolve

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entify/common"
	"rxdrag.com/entify/utils"
)

func Me(p graphql.ResolveParams) (interface{}, error) {
	defer utils.PrintErrorStack()
	return common.ParseContextValues(p).Me, nil
}
