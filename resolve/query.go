package resolve

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"rxdrag.com/entify/authorization"
	"rxdrag.com/entify/consts"
	"rxdrag.com/entify/model/graph"
	"rxdrag.com/entify/repository"
	"rxdrag.com/entify/utils"
)

func QueryOneResolveFn(node graph.Noder) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		defer utils.PrintErrorStack()
		instance := repository.QueryOne(node, extractArgs(p))
		return instance, nil
	}
}

func QueryResolveFn(node graph.Noder) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		defer utils.PrintErrorStack()
		// for _, iSelection := range p.Info.Operation.GetSelectionSet().Selections {
		// 	switch selection := iSelection.(type) {
		// 	case *ast.Field:
		// 		//fmt.Println(selection.Directives[len(selection.Directives)-1].Name.Value)
		// 	case *ast.InlineFragment:
		// 	case *ast.FragmentSpread:
		// 	}
		// }

		return repository.Query(node, extractArgs(p)), nil
	}
}

func QueryAssociationFn(asso *graph.Association) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		var (
			source      = p.Source.(map[string]interface{})
			v           = p.Context.Value
			loaders     = v(consts.LOADERS).(*Loaders)
			handleError = func(err error) error {
				return fmt.Errorf(err.Error())
			}
		)
		defer utils.PrintErrorStack()

		if loaders == nil {
			panic("Data loaders is nil")
		}
		loader := loaders.GetLoader(asso)
		thunk := loader.Load(p.Context, NewKey(source[consts.ID].(uint64)))
		return func() (interface{}, error) {
			data, err := thunk()
			if err != nil {
				return nil, handleError(err)
			}

			var retValue interface{}
			if data == nil {
				if asso.IsArray() {
					retValue = []map[string]interface{}{}
				} else {
					retValue = nil
				}
			} else {
				retValue = data
			}
			return retValue, nil
		}, nil
	}
}

func extractArgs(p graphql.ResolveParams) map[string]interface{} {
	verifier := authorization.ParseAbilityVerifier(p)

	if verifier == nil {
		panic("Can not finde Ability Verifier")
	}

	inputArgs := map[string]interface{}{}
	if p.Args != nil {
		inputArgs = p.Args
	}
	args := verifier.WeaveAuthInArgs(inputArgs)

	return args
}
