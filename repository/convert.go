package repository

import (
	"database/sql"
	"encoding/json"

	"rxdrag.com/entity-engine/db"
	"rxdrag.com/entity-engine/model/data"
	"rxdrag.com/entity-engine/model/graph"
	"rxdrag.com/entity-engine/model/meta"
	"rxdrag.com/entity-engine/utils"
)

func makeSaveValues(fields []*data.Field) []interface{} {
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

func makeQueryValues(node graph.Noder) []interface{} {
	names := node.AllAttributeNames()
	values := make([]interface{}, len(names))
	for i, attrName := range names {
		attr := node.GetAttributeByName(attrName)
		switch attr.Type {
		case meta.ID:
			var value db.NullUint64
			values[i] = &value
			break
		case meta.INT:
			var value sql.NullInt64
			values[i] = &value
			break
		case meta.FLOAT:
			var value sql.NullFloat64
			values[i] = &value
			break
		case meta.BOOLEAN:
			var value sql.NullBool
			values[i] = &value
			break
		case meta.DATE:
			var value sql.NullTime
			values[i] = &value
			break
		case meta.CLASS_VALUE_OBJECT,
			meta.ID_ARRAY,
			meta.INT_ARRAY,
			meta.FLOAT_ARRAY,
			meta.STRING_ARRAY,
			meta.DATE_ARRAY,
			meta.ENUM_ARRAY,
			meta.VALUE_OBJECT_ARRAY,
			meta.ENTITY_ARRAY:
			var value utils.JSON
			values[i] = &value
			break
			// COLUMN_SIMPLE_ARRAY string = "simpleArray" ##待添加代码
			// COLUMN_JSON_ARRAY   string = "JsonArray"
		default:
			var value sql.NullString
			values[i] = &value
		}
	}

	return values
}

func convertValuesToObject(values []interface{}, node graph.Noder) map[string]interface{} {
	object := make(map[string]interface{})
	names := node.AllAttributeNames()
	for i := range names {
		value := values[i]
		attrName := names[i]
		column := node.GetAttributeByName(attrName)
		switch column.Type {
		case meta.ID:
			nullValue := value.(*db.NullUint64)
			if nullValue.Valid {
				object[attrName] = nullValue.Uint64
			}
			break
		case meta.INT:
			nullValue := value.(*sql.NullInt64)
			if nullValue.Valid {
				object[attrName] = nullValue.Int64
			}
			break
		case meta.FLOAT:
			nullValue := value.(*sql.NullFloat64)
			if nullValue.Valid {
				object[attrName] = nullValue.Float64
			}
			break
		case meta.BOOLEAN:
			nullValue := value.(*sql.NullBool)
			if nullValue.Valid {
				object[attrName] = nullValue.Bool
			}
			break
		case meta.DATE:
			nullValue := value.(*sql.NullTime)
			if nullValue.Valid {
				object[attrName] = nullValue.Time
			}
			break
		case meta.VALUE_OBJECT,
			meta.ID_ARRAY,
			meta.INT_ARRAY,
			meta.FLOAT_ARRAY,
			meta.STRING_ARRAY,
			meta.DATE_ARRAY,
			meta.ENUM_ARRAY,
			meta.VALUE_OBJECT_ARRAY,
			meta.ENTITY_ARRAY:
			object[attrName] = value
			break
		default:
			nullValue := value.(*sql.NullString)
			if nullValue.Valid {
				object[attrName] = nullValue.String
			}
		}

	}
	return object
}
