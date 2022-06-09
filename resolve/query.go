package resolve

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"rxdrag.com/entify/consts"
	"rxdrag.com/entify/model/graph"
	"rxdrag.com/entify/repository"
	"rxdrag.com/entify/utils"
)

func QueryOneInterfaceResolveFn(intf *graph.Interface) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		defer utils.PrintErrorStack()
		v := makeInterfaceAbilityVerifier(p, intf)
		instance := repository.QueryOneInterface(intf, p.Args, v)
		return instance, nil
	}
}

func QueryInterfaceResolveFn(intf *graph.Interface) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		defer utils.PrintErrorStack()
		v := makeInterfaceAbilityVerifier(p, intf)
		return repository.QueryInterface(intf, p.Args, v), nil
	}
}

func QueryOneEntityResolveFn(entity *graph.Entity) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		defer utils.PrintErrorStack()
		v := makeEntityAbilityVerifier(p, entity.Uuid())
		instance := repository.QueryOneEntity(entity, p.Args, v)
		return instance, nil
	}
}

func QueryEntityResolveFn(entity *graph.Entity) graphql.FieldResolveFn {
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
		v := makeEntityAbilityVerifier(p, entity.Uuid())
		return repository.QueryEntity(entity, p.Args, v), nil
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
		loader := loaders.GetLoader(p, asso, p.Args)
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
