package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/model"
)

func (c *TypeCache) makeEnums(enums []*model.Enum) {
	for i := range enums {
		enum := enums[i]
		c.EnumTypeMap[enum.Name] = EnumType(enum)
	}
}

func EnumType(entity *model.Enum) *graphql.Enum {
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
