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
	typeStr := "text"
	switch column.Type {
	case meta.COLUMN_ID:
		typeStr = "int UNSIGNED"
		break
	case meta.COLUMN_INT:
		typeStr = "int"
		if column.Length == 1 {
			typeStr = "tinyint"
		} else if column.Length == 2 {
			typeStr = "smallint"
		} else if column.Length == 3 {
			typeStr = "mediumint"
		} else if column.Length == 4 {
			typeStr = "int"
		} else if column.Length > 4 {
			typeStr = "bigint"
		}
		if column.Unsigned {
			typeStr = typeStr + " UNSIGNED"
		}
		break
	case meta.COLUMN_FLOAT:
		if column.Length > 4 {
			typeStr = "double"
		} else {
			typeStr = "float"
		}
		if column.FloatM > 0 && column.FloatD > 0 && column.FloatM >= column.FloatD {
			typeStr = fmt.Sprint(typeStr+"(%d,%d)", column.FloatM, column.FloatD)
		}
		if column.Unsigned {
			typeStr = typeStr + " UNSIGNED"
		}
		break
	case meta.COLUMN_BOOLEAN:
		typeStr = "tinyint(1)"
		break
	case meta.COLUMN_STRING:
		typeStr = "text"
		if column.Length > 0 {
			if column.Length <= 255 {
				typeStr = fmt.Sprintf("varchar(%d)", column.Length)
			} else if column.Length <= 65535 {
				typeStr = "text"
			} else if column.Length <= 16777215 {
				typeStr = "mediumtext"
			} else {
				typeStr = "longtext"
			}
		}
		break
	case meta.COLUMN_DATE:
		typeStr = "datetime"
		break
	case meta.COLUMN_SIMPLE_JSON:
		typeStr = "json"
		break
	case meta.COLUMN_SIMPLE_ARRAY:
		typeStr = "json"
		break
	case meta.COLUMN_JSON_ARRAY:
		typeStr = "json"
		break
	case meta.COLUMN_ENUM:
		typeStr = "tinytext"
		break
	}
	return typeStr
}

func (b *MySQLBuilder) BuildColumnSQL(column *meta.Column) string {
	sql := "`" + column.Name + "` " + b.ColumnTypeSQL(column)
	if column.Generated {
		sql = sql + " AUTO_INCREMENT"
	}
	return sql
}

func (b *MySQLBuilder) BuildCreateEntitySQL(entity *meta.Entity) string {
	sql := "CREATE TABLE `%s` (%s)"
	fieldSqls := make([]string, len(entity.Columns))
	for i := range entity.Columns {
		columnSql := b.BuildColumnSQL(&entity.Columns[i])
		if entity.Columns[i].Nullable {
			columnSql = columnSql + " NULL"
		} else {
			columnSql = columnSql + " NOT NULL"
		}
		fieldSqls[i] = columnSql
	}
	for _, column := range entity.Columns {
		if column.Primary {
			fieldSqls = append(fieldSqls, fmt.Sprintf("PRIMARY KEY (`%s`)", column.Name))
		}
	}
	return fmt.Sprintf(sql, entity.GetTableName(), strings.Join(fieldSqls, ","))
}
