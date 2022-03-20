package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/meta"
)

func (c *TypeCache) makeOutputInterfaces(interfaces []*meta.Entity) {
	for i := range interfaces {
		entity := interfaces[i]
		c.InterfaceTypeMap[entity.Name] = c.InterfaceType(entity)
	}
}

func (c *TypeCache) InterfaceType(entity *meta.Entity) *graphql.Interface {
	name := entity.Name

	parent := meta.Metas.Interfaces(entity)
	if parent != nil {
		return graphql.NewInterface(
			graphql.InterfaceConfig{
				Name:   name,
				Fields: outputFields(entity),
			},
		)
	} else {
		return graphql.NewInterface(
			graphql.InterfaceConfig{
				Name:   name,
				Fields: outputFields(entity),
			},
		)
	}

}
