package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entify/model/graph"
	"rxdrag.com/entify/model/meta"
)

func (c *TypeCache) makeOutputObjects(nodes []*graph.Entity) {
	for i := range nodes {
		entity := nodes[i]
		objType := c.ObjectType(entity)
		c.ObjectTypeMap[entity.Name()] = objType
		c.ObjectMapById[entity.InnerId()] = objType
		if entity.Domain.StereoType == meta.CLASS_PARTIAL {
			partialName := entity.Domain.Name
			c.ObjectTypeMap[partialName] = graphql.NewObject(
				graphql.ObjectConfig{
					Name:        partialName,
					Fields:      outputFields(entity.AllAttributes(), entity.AllMethods()),
					Description: entity.Description(),
				},
			)
		}
	}
}

func (c *TypeCache) ObjectType(entity *graph.Entity) *graphql.Object {
	name := entity.Name()
	interfaces := c.mapInterfaces(entity.Interfaces)

	if len(interfaces) > 0 {
		return graphql.NewObject(
			graphql.ObjectConfig{
				Name:        name,
				Fields:      outputFields(entity.AllAttributes(), entity.AllMethods()),
				Description: entity.Description(),
				Interfaces:  interfaces,
			},
		)
	} else {
		return graphql.NewObject(
			graphql.ObjectConfig{
				Name:        name,
				Fields:      outputFields(entity.AllAttributes(), entity.AllMethods()),
				Description: entity.Description(),
			},
		)
	}

}

func outputFields(attrs []*graph.Attribute, methods []*graph.Method) graphql.Fields {
	fields := graphql.Fields{}
	for _, attr := range attrs {
		fields[attr.Name] = &graphql.Field{
			Type:        AttributeType(attr),
			Description: attr.Description,
			// Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// 	fmt.Println(p.Context.Value("data"))
			// 	return "world", nil
			// },
		}
	}

	for _, method := range methods {
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
