package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/meta"
)

func (c *TypeCache) makeEnums(enums []*meta.EntityMeta) {
	for i := range enums {
		entity := enums[i]
		c.EnumTypeMap[entity.Name] = EnumType(entity)
	}
}

func EnumType(entity *meta.EntityMeta) *graphql.Enum {
	enumValues := entity.EnumValues
	enumValueConfigMap := graphql.EnumValueConfigMap{}
	for enumName, enumValue := range enumValues {
		var value, ok = enumValue.(string)
		if !ok {
			value = enumValue.(map[string]string)["value"]
		}
		enumValueConfigMap[enumName] = &graphql.EnumValueConfig{
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
