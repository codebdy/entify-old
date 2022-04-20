package resolve

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/model/graph"
	"rxdrag.com/entity-engine/repository"
	"rxdrag.com/entity-engine/utils"
)

func QueryOneResolveFn(node graph.Noder) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		defer utils.PrintErrorStack()
		return repository.QueryOne(node, p.Args)
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

		//err = db.Select(&instances, queryStr)
		return repository.Query(node, p.Args)
	}
}

func QueryAssociationFn(asso *graph.Association) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		defer utils.PrintErrorStack()
		// for _, iSelection := range p.Info.Operation.GetSelectionSet().Selections {
		// 	switch selection := iSelection.(type) {
		// 	case *ast.Field:
		// 		fmt.Println(selection.Directives[len(selection.Directives)-1].Name.Value)
		// 	case *ast.InlineFragment:
		// 	case *ast.FragmentSpread:
		// 	}
		// }
		loadersObj := p.Context.Value(consts.LOADERS)
		if loadersObj == nil {
			panic("Data loaders is nil")
		}
		loaders := loadersObj.(Loaders)
		loader := loaders.GetLoader(asso)
		loader.LoadMany(p.Context)
		fmt.Println("哈哈", asso)
		var retValue interface{}
		if asso.IsArray() {
			retValue = []map[string]interface{}{}
		} else {
			retValue = nil
		}
		//err = db.Select(&instances, queryStr)
		return retValue, nil //repository.Query(node, p.Args)
	}
}
