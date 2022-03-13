package dialect

import (
	"fmt"
	"strings"

	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/meta"
)

type MySQLBuilder struct {
}

func (*MySQLBuilder) BuildFieldExp(fieldName string, fieldArgs map[string]interface{}) (string, []interface{}) {
	var params []interface{}
	queryStr := "true "
	for key, value := range fieldArgs {
		switch key {
		case consts.ARG_EQ:
			queryStr = queryStr + " AND " + fieldName + "=?"
			params = append(params, value)
			break
		case consts.ARG_ISNULL:
			if value == true {
				queryStr = queryStr + " AND ISNULL(" + fieldName + ")"
			}
			break
		default:
			panic("Can not find token:" + key)
		}
	}
	return "(" + queryStr + ")", params
}

func (b *MySQLBuilder) BuildBoolExp(where map[string]interface{}) (string, []interface{}) {
	var params []interface{}
	queryStr := ""
	for key, value := range where {
		switch key {
		case consts.ARG_AND:
			break
		case consts.ARG_NOT:
			break
		case consts.ARG_OR:
			break
		default:
			fiedleStr, fieldParam := b.BuildFieldExp(key, value.(map[string]interface{}))
			queryStr = queryStr + " AND " + fiedleStr
			params = append(params, fieldParam...)
		}
	}
	return queryStr, params
}

func (b *MySQLBuilder) ColumnTypeSQL(column *meta.Column) string {
	typeStr := "TEXT"
	fmt.Println("哈哈：", column.Type)
	switch column.Type {
	case meta.COLUMN_ID:
		typeStr = "INT UNSIGNED"
		break
	case meta.COLUMN_INT:
		typeStr = "int"
		if column.Length == 1 {
			typeStr = "TINYINT"
		}
		if column.Length == 2 {
			typeStr = "SMALLINT"
		}
		if column.Length == 3 {
			typeStr = "MEDIUMINT"
		}
		if column.Length == 4 {
			typeStr = "INT"
		}
		if column.Length > 4 {
			typeStr = "BIGINT"
		}
		if column.Unsigned {
			typeStr = typeStr + " UNSIGNED"
		}
		break
	case meta.COLUMN_FLOAT:
		typeStr = "FLOAT"
		if column.Length > 4 {
			typeStr = "DOUBLE"
		}
		if column.FloatM > 0 && column.FloatD > 0 && column.FloatM >= column.FloatD {
			typeStr = fmt.Sprint(typeStr+"(%d,%d)", column.FloatM, column.FloatD)
		}
		if column.Unsigned {
			typeStr = typeStr + " UNSIGNED"
		}
		break
	case meta.COLUMN_BOOLEAN:
		typeStr = "TINYINT(1)"
		break
	case meta.COLUMN_STRING:
		typeStr = "TEXT"
		if column.Length > 0 {
			if column.Length <= 255 {
				typeStr = "TINYTEXT"
			} else if column.Length <= 65535 {
				typeStr = "TEXT"
			} else if column.Length <= 16777215 {
				typeStr = "MEDIUMTEXT"
			} else {
				typeStr = "LONGTEXT"
			}
		}
		break
	case meta.COLUMN_DATE:
		typeStr = "DATETIME"
		break
	case meta.COLUMN_SIMPLE_JSON:
		typeStr = "JSON"
		break
	case meta.COLUMN_SIMPLE_ARRAY:
		typeStr = "JSON"
		break
	case meta.COLUMN_JSON_ARRAY:
		typeStr = "JSON"
		break
	case meta.COLUMN_ENUM:
		typeStr = "TINYTEXT"
		break
	}
	return typeStr
}

func (b *MySQLBuilder) BuildColumnSQL(column *meta.Column) string {
	return column.Name + " " + b.ColumnTypeSQL(column)
}

func (b *MySQLBuilder) BuildCreateEntitySQL(entity *meta.Entity) string {
	sql := `CREATE TABLE %s (%s)`
	fieldSqls := make([]string, len(entity.Columns))
	for i := range entity.Columns {
		fieldSqls[i] = b.BuildColumnSQL(&entity.Columns[i])
	}
	return fmt.Sprintf(sql, entity.GetTableName(), strings.Join(fieldSqls, ","))
}
