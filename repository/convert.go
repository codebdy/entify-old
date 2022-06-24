package repository

import (
	"database/sql"
	"encoding/json"

	"rxdrag.com/entify/db"
	"rxdrag.com/entify/model/data"
	"rxdrag.com/entify/model/graph"
	"rxdrag.com/entify/model/meta"
	"rxdrag.com/entify/utils"
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

func makeInterfaceQueryValues(intf *graph.Interface) []interface{} {
	names := intf.AllAttributeNames()
	values := make([]interface{}, len(names))
	for i, attrName := range names {
		attr := intf.GetAttributeByName(attrName)
		values[i] = makeAttributeValue(attr)
	}

	return values
}

func makeEntityQueryValues(ent *graph.Entity) []interface{} {
	names := ent.AllAttributeNames()
	values := make([]interface{}, len(names))
	for i, attrName := range names {
		attr := ent.GetAttributeByName(attrName)
		values[i] = makeAttributeValue(attr)
	}

	return values
}

func makeAttributeValue(attr *graph.Attribute) interface{} {
	switch attr.Type {
	case meta.ID:
		var value db.NullUint64
		return &value
	case meta.INT:
		var value sql.NullInt64
		return &value
	case meta.FLOAT:
		var value sql.NullFloat64
		return &value
	case meta.BOOLEAN:
		var value sql.NullBool
		return &value
	case meta.DATE:
		var value sql.NullTime
		return &value
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
		return &value
		// COLUMN_SIMPLE_ARRAY string = "simpleArray" ##待添加代码
		// COLUMN_JSON_ARRAY   string = "JsonArray"
	default:
		var value sql.NullString
		return &value
	}
}

func convertValuesToInterface(values []interface{}, intf *graph.Interface) map[string]interface{} {
	object := make(map[string]interface{})
	names := intf.AllAttributeNames()
	for i := range names {
		value := values[i]
		attrName := names[i]
		column := intf.GetAttributeByName(attrName)
		object[column.Name] = convertOneColumnValue(column, value)

	}
	return object
}

func convertValuesToEntity(values []interface{}, ent *graph.Entity) map[string]interface{} {
	object := make(map[string]interface{})
	names := ent.AllAttributeNames()
	for i := range names {
		value := values[i]
		attrName := names[i]
		column := ent.GetAttributeByName(attrName)
		object[column.Name] = convertOneColumnValue(column, value)

	}
	return object
}

func convertOneColumnValue(column *graph.Attribute, value interface{}) interface{} {
	switch column.Type {
	case meta.ID:
		nullValue := value.(*db.NullUint64)
		if nullValue.Valid {
			return nullValue.Uint64
		}
	case meta.INT:
		nullValue := value.(*sql.NullInt64)
		if nullValue.Valid {
			return nullValue.Int64
		}
		break
	case meta.FLOAT:
		nullValue := value.(*sql.NullFloat64)
		if nullValue.Valid {
			return nullValue.Float64
		}
		break
	case meta.BOOLEAN:
		nullValue := value.(*sql.NullBool)
		if nullValue.Valid {
			return nullValue.Bool
		}
		break
	case meta.DATE:
		nullValue := value.(*sql.NullTime)
		if nullValue.Valid {
			return nullValue.Time
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
		return value
	default:
		nullValue := value.(*sql.NullString)
		if nullValue.Valid {
			return nullValue.String
		}
	}
	return value
}
