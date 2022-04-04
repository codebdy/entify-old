package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/model/graph"
)

func (c *TypeCache) makeOutputObjects(normals []*graph.Entity) {
	for i := range normals {
		entity := normals[i]
		c.ObjectTypeMap[entity.Name()] = c.ObjectType(entity)
	}
}

func (c *TypeCache) ObjectType(entity *graph.Entity) *graphql.Object {
	name := entity.Name()
	interfaces := c.mapInterfaces(entity.Interfaces)
	if len(interfaces) > 0 {
		return graphql.NewObject(
			graphql.ObjectConfig{
				Name:        name,
				Fields:      outputFields(entity.AllAttributes()),
				Description: entity.Description(),
				Interfaces:  interfaces,
			},
		)
	} else {
		return graphql.NewObject(
			graphql.ObjectConfig{
				Name:        name,
				Fields:      outputFields(entity.AllAttributes()),
				Description: entity.Description(),
			},
		)
	}

}

func outputFields(attrs []*graph.Attribute) graphql.Fields {
	fields := graphql.Fields{}
	for _, column := range attrs {
		fields[column.Name] = &graphql.Field{
			Type:        AttributeType(column),
			Description: column.Description,
			// Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// 	fmt.Println(p.Context.Value("data"))
			// 	return "world", nil
			// },
		}
	}
	return fields
}
