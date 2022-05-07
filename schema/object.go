package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entify/model/graph"
)

func (c *TypeCache) makeOutputObjects(nodes []*graph.Entity) {
	for i := range nodes {
		entity := nodes[i]
		objType := c.ObjectType(entity)
		c.ObjectTypeMap[entity.Name()] = objType
		c.ObjectMapById[entity.InnerId()] = objType
	}
}

func (c *TypeCache) ObjectType(entity *graph.Entity) *graphql.Object {
	name := entity.Name()
	interfaces := c.mapInterfaces(entity.Interfaces)
	interfaces = append(interfaces, NodeInterfaceType)

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

func outputFields(node graph.Noder) graphql.Fields {
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
