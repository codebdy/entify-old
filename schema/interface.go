package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/model"
)

func (c *TypeCache) makeOutputInterfaces(interfaces []*model.Interface) {
	for i := range interfaces {
		entity := interfaces[i]
		c.InterfaceTypeMap[entity.Name] = c.InterfaceType(entity)
	}
}

func (c *TypeCache) InterfaceType(entity *model.Interface) *graphql.Interface {
	name := entity.Name

	return graphql.NewInterface(
		graphql.InterfaceConfig{
			Name:        name,
			Fields:      outputFields(entity.Columns),
			Description: entity.Description,
		},
	)

}
