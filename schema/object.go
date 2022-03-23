package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/meta"
	"rxdrag.com/entity-engine/model"
)

func (c *TypeCache) makeOutputObjects(normals []*model.Entity) {
	for i := range normals {
		entity := normals[i]
		c.ObjectTypeMap[entity.Name] = c.ObjectType(entity)
	}
}

func (c *TypeCache) ObjectType(entity *model.Entity) *graphql.Object {
	name := entity.Name
	interfaces := c.mapInterfaces(entity.Interfaces)
	if len(interfaces) > 0 {
		return graphql.NewObject(
			graphql.ObjectConfig{
				Name:        name,
				Fields:      outputFields(entity.AllColumns),
				Description: entity.Description,
				Interfaces:  interfaces,
			},
		)
	} else {
		return graphql.NewObject(
			graphql.ObjectConfig{
				Name:        name,
				Fields:      outputFields(entity.AllColumns),
				Description: entity.Description,
			},
		)
	}

}

func outputFields(columns []meta.ColumnMeta) graphql.Fields {
	fields := graphql.Fields{}
	for _, column := range columns {
		fields[column.Name] = &graphql.Field{
			Type:        ColumnType(&column),
			Description: column.Description,
			// Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// 	fmt.Println(p.Context.Value("data"))
			// 	return "world", nil
			// },
		}
	}
	return fields
}
