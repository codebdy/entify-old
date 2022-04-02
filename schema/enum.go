package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/model/graph"
)

func (c *TypeCache) makeEnums(enums []*graph.Enum) {
	for i := range enums {
		enum := enums[i]
		c.EnumTypeMap[enum.Name] = EnumType(enum)
	}
}

func EnumType(entity *graph.Enum) *graphql.Enum {
	enumValueConfigMap := graphql.EnumValueConfigMap{}
	for _, value := range entity.Values {
		enumValueConfigMap[value] = &graphql.EnumValueConfig{
			Value: value,
		}
	}
	enum := graphql.NewEnum(
		graphql.EnumConfig{
			Name:   entity.Name,
			Values: enumValueConfigMap,
		},
	)
	return enum
}
