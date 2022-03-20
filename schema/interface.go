package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/meta"
)

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
