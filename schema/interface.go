package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/meta"
)

func (c *TypeCache) makeOutputInterfaces(interfaces []*meta.EntityMeta) {
	for i := range interfaces {
		entity := interfaces[i]
		c.InterfaceTypeMap[entity.Name] = c.InterfaceType(entity)
	}
}

func (c *TypeCache) InterfaceType(entity *meta.EntityMeta) *graphql.Interface {
	name := entity.Name

	return graphql.NewInterface(
		graphql.InterfaceConfig{
			Name:        name,
			Fields:      outputFields(entity),
			Description: entity.Description,
		},
	)

}
