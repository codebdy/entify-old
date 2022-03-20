package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/meta"
)

func ObjectType(entity *meta.Entity) *graphql.Object {
	name := entity.Name

	return graphql.NewObject(
		graphql.ObjectConfig{
			Name:   name,
			Fields: outputFields(entity),
		},
	)
}

func outputFields(entity *meta.Entity) graphql.Fields {
	fields := graphql.Fields{}
	for _, column := range meta.Metas.EntityAllColumns(entity) {
		fields[column.Name] = &graphql.Field{
			Type: ColumnType(&column),
			// Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// 	fmt.Println(p.Context.Value("data"))
			// 	return "world", nil
			// },
		}
	}
	return fields
}
