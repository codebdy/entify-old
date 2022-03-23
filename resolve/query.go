package resolve

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"rxdrag.com/entity-engine/meta"
	"rxdrag.com/entity-engine/repository"
)

func QueryOneResolveFn(entity *meta.EntityMeta) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		return repository.QueryOne(entity, p.Args)
	}
}

func QueryResolveFn(entity *meta.EntityMeta) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
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
		return repository.Query(entity, p.Args)
	}
}

func RelationResolveFn(relation *meta.EntityRelation) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		return nil, nil
	}
}
