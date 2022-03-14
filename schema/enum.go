package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/meta"
)

func EnumType(entity *meta.Entity) *graphql.Enum {
	if Cache.EnumMap[entity.Name] != nil {
		return Cache.EnumMap[entity.Name]
	}
	enumValues := entity.EnumValues //make(map[string]interface{})
	//json.Unmarshal(entity.EnumValues, &enumValues)
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
	Cache.EnumMap[entity.Name] = enum
	return enum
}
