package schema

import (
	"encoding/json"

	"github.com/graphql-go/graphql"
)

//类型缓存， query mutaion通用
var InsertInputMap = make(map[string]*graphql.Input)

func (entity *EntityMeta) createInsertFields() graphql.InputObjectConfigFieldMap {
	fields := graphql.InputObjectConfigFieldMap{}
	for _, column := range entity.Columns {
		if column.Name != "id" {
			fields[column.Name] = &graphql.InputObjectFieldConfig{
				Type: column.toInputType(),
			}
		}
	}

	return fields
}

func (entity *EntityMeta) toInsertInput() graphql.Input {
	if InsertInputMap[entity.Name] != nil {
		return *InsertInputMap[entity.Name]
	}
	var returnValue graphql.Input

	if entity.EntityType == Entity_ENUM {
		enumValues := make(map[string]interface{})
		json.Unmarshal(entity.EnumValues, &enumValues)
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
		returnValue = graphql.NewEnum(
			graphql.EnumConfig{
				Name:   entity.Name,
				Values: enumValueConfigMap,
			},
		)
	} else {
		returnValue = graphql.NewInputObject(
			graphql.InputObjectConfig{
				Name:   entity.Name + "InsertInput",
				Fields: entity.createInsertFields(),
			},
		)
	}
	InsertInputMap[entity.Name] = &returnValue
	return returnValue
}
