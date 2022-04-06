package resolve

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"rxdrag.com/entity-engine/model/graph"
	"rxdrag.com/entity-engine/repository"
	"rxdrag.com/entity-engine/utils"
)

func QueryOneResolveFn(node graph.Node) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		defer utils.PrintErrorStack()
		return repository.QueryOne(node, p.Args)
	}
}

func QueryResolveFn(node graph.Node) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		defer utils.PrintErrorStack()
		// names := entity.ColumnNames()
		// queryStr := "select %s from %s "
		for _, iSelection := range p.Info.Operation.GetSelectionSet().Selections {
			switch selection := iSelection.(type) {
			case *ast.Field:
				fmt.Println(selection.Directives[len(selection.Directives)-1].Name.Value)
			case *ast.InlineFragment:
			case *ast.FragmentSpread:
			}
		}

		//err = db.Select(&instances, queryStr)
		return repository.Query(node, p.Args)
	}
}
