package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entify/model/graph"
)

func (c *TypeCache) makeEntityOutputObjects(entities []*graph.Entity) {
	for i := range entities {
		c.makeEntityObject(entities[i])

	}
}

func (c *TypeCache) makePartialOutputObjects(partials []*graph.Partial) {
	for i := range partials {
		patial := partials[i]
		c.makeEntityObject(&patial.Entity)

		// partialName := patial.NameWithPartial()

		// objType := graphql.NewObject(
		// 	graphql.ObjectConfig{
		// 		Name:        partialName,
		// 		Fields:      outputFields(patial.AllAttributes(), patial.AllMethods()),
		// 		Description: patial.Description(),
		// 	},
		// )
		// c.ObjectTypeMap[partialName] = objType
	}
}

func (c *TypeCache) makeEntityObject(entity *graph.Entity) {
	objType := c.ObjectType(entity)
	c.ObjectTypeMap[entity.Name()] = objType
	c.ObjectMapById[entity.InnerId()] = objType
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
			Type:        PropertyType(attr),
			Description: attr.Description,
			// Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// 	fmt.Println(p.Context.Value("data"))
			// 	return "world", nil
			// },
		}
	}

	for _, method := range methods {
		fields[method.GetName()] = &graphql.Field{
			Type:        PropertyType(method),
			Description: method.Method.Description,
			// Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// 	fmt.Println(p.Context.Value("data"))
			// 	return "world", nil
			// },
		}
	}
	return fields
}
