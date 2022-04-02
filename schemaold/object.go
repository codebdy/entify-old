package schemaold

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/modleold"
)

func (c *TypeCache) makeOutputObjects(normals []*modleold.Entity) {
	for i := range normals {
		entity := normals[i]
		c.ObjectTypeMap[entity.Name] = c.ObjectType(entity)
	}
}

func (c *TypeCache) ObjectType(entity *modleold.Entity) *graphql.Object {
	name := entity.Name
	interfaces := c.mapInterfaces(entity.Interfaces)
	if len(interfaces) > 0 {
		return graphql.NewObject(
			graphql.ObjectConfig{
				Name:        name,
				Fields:      outputFields(entity.Columns),
				Description: entity.Description,
				Interfaces:  interfaces,
			},
		)
	} else {
		return graphql.NewObject(
			graphql.ObjectConfig{
				Name:        name,
				Fields:      outputFields(entity.Columns),
				Description: entity.Description,
			},
		)
	}

}

func outputFields(columns []*modleold.Column) graphql.Fields {
	fields := graphql.Fields{}
	for _, column := range columns {
		fields[column.Name] = &graphql.Field{
			Type:        ColumnType(column),
			Description: column.Description,
			// Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// 	fmt.Println(p.Context.Value("data"))
			// 	return "world", nil
			// },
		}
	}
	return fields
}
