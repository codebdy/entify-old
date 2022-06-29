package resolve

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"rxdrag.com/entify/utils"
)

func SingleUploadResolve(p graphql.ResolveParams) (interface{}, error) {
	defer utils.PrintErrorStack()
	fmt.Println("呵呵:SingleUploadResolve", p.Args["file"])
	return nil, nil
}
