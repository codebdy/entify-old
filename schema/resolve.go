package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entify/resolve"
)

func publishResolve(p graphql.ResolveParams) (interface{}, error) {
	result, err := resolve.PublishMetaResolve(p)
	if err != nil {
		return result, err
	}

	MakeSchema()
	return result, nil
}

func installResolve(p graphql.ResolveParams) (interface{}, error) {
	result, err := resolve.InstallResolve(p)
	if err != nil {
		return result, err
	}
	MakeSchema()
	return result, err
}
