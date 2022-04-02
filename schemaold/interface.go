package schemaold

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/model"
)

func (c *TypeCache) makeOutputInterfaces(interfaces []*model.Entity) {
	for i := range interfaces {
		entity := interfaces[i]
		c.InterfaceTypeMap[entity.Name] = c.InterfaceType(entity)
	}
}

func (c *TypeCache) InterfaceType(entity *model.Entity) *graphql.Interface {
	name := entity.Name

	return graphql.NewInterface(
		graphql.InterfaceConfig{
			Name:        name,
			Fields:      outputFields(entity.Columns),
			Description: entity.Description,
		},
	)

}
