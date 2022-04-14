package data

import (
	"encoding/json"

	"rxdrag.com/entity-engine/model/meta"
)

func MakeValues(fields []*Field) []interface{} {
	objValues := make([]interface{}, 0, len(fields))
	for _, field := range fields {
		value := field.Value
		column := field.Column

		if column.Type == meta.VALUE_OBJECT ||
			column.Type == meta.ID_ARRAY ||
			column.Type == meta.INT_ARRAY ||
			column.Type == meta.FLOAT_ARRAY ||
			column.Type == meta.STRING_ARRAY ||
			column.Type == meta.DATE_ARRAY ||
			column.Type == meta.ENUM_ARRAY ||
			column.Type == meta.VALUE_OBJECT_ARRAY ||
			column.Type == meta.ENTITY_ARRAY {
			value, _ = json.Marshal(value)
		}
		objValues = append(objValues, value)
	}
	return objValues
}
