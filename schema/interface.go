package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/meta"
)

func InterfaceType(entity *meta.Entity) *graphql.Interface {
	name := entity.Name

	return graphql.NewInterface(
		graphql.InterfaceConfig{
			Name:   name,
			Fields: outputFields(entity),
		},
	)
}
