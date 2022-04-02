package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/model/graph"
)

func (c *TypeCache) makeOutputInterfaces(interfaces []*graph.Entity) {
	for i := range interfaces {
		entity := interfaces[i]
		c.InterfaceTypeMap[entity.Name()] = c.InterfaceType(entity)
	}
}

func (c *TypeCache) InterfaceType(entity *graph.Entity) *graphql.Interface {
	name := entity.Name()

	return graphql.NewInterface(
		graphql.InterfaceConfig{
			Name:        name,
			Fields:      outputFields(entity.Attributes),
			Description: entity.Description(),
		},
	)

}
