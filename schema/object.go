package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/model/graph"
	"rxdrag.com/entity-engine/model/meta"
)

func (c *TypeCache) makeOutputObjects(nodes []*graph.Entity) {
	for i := range nodes {
		entity := nodes[i]
		c.ObjectTypeMap[entity.Name()] = c.ObjectType(entity)
	}
}

func (c *TypeCache) ObjectType(entity *graph.Entity) *graphql.Object {
	name := entity.Name()
	interfaces := c.mapInterfaces(entity.Interfaces)
	if entity.Domain.StereoType != meta.CLASS_SERVICE && entity.Domain.StereoType != meta.CLASS_VALUE_OBJECT {
		interfaces = append(interfaces, NodeInterfaceType)
	}
	if len(interfaces) > 0 {
		return graphql.NewObject(
			graphql.ObjectConfig{
				Name:        name,
				Fields:      outputFields(entity),
				Description: entity.Description(),
				Interfaces:  interfaces,
			},
		)
	} else {
		return graphql.NewObject(
			graphql.ObjectConfig{
				Name:        name,
				Fields:      outputFields(entity),
				Description: entity.Description(),
			},
		)
	}

}

func outputFields(node graph.Node) graphql.Fields {
	fields := graphql.Fields{}
	for _, attr := range node.AllAttributes() {
		fields[attr.Name] = &graphql.Field{
			Type:        AttributeType(attr),
			Description: attr.Description,
			// Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// 	fmt.Println(p.Context.Value("data"))
			// 	return "world", nil
			// },
		}
	}

	for _, method := range node.AllMethods() {
		fields[method.Name()] = &graphql.Field{
			Type:        MethodType(method),
			Description: method.Method.Description,
			// Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// 	fmt.Println(p.Context.Value("data"))
			// 	return "world", nil
			// },
		}
	}
	return fields
}
