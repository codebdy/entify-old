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
	switch column.Type {
	case meta.COLUMN_ID:
		return "int"
	case meta.COLUMN_INT:
		return "int"
	case meta.COLUMN_FLOAT:
		return "int"
	case meta.COLUMN_BOOLEAN:
		return "tinyint(1)"
	case meta.COLUMN_STRING:
		return "text"
	case meta.COLUMN_DATE:
		return "datetime"
	case meta.COLUMN_SIMPLE_JSON:
		return "json"
	case meta.COLUMN_SIMPLE_ARRAY:
		return "json"
	case meta.COLUMN_JSON_ARRAY:
		return "json"
	case meta.COLUMN_ENUM:
		return "varchar(255)"
	}
	return "varchar(255)"
}

func (b *MySQLBuilder) BuildColumnSQL(column *meta.Column) string {
	return column.Name + " " + b.BuildColumnSQL(column)
}

func (b *MySQLBuilder) BuildCreateEntitySQL(entity *meta.Entity) string {
	sql := `CREATE TABLE %s (%s)`
	fieldSqls := make([]string, len(entity.Columns))
	for i := range entity.Columns {
		fieldSqls[i] = b.BuildColumnSQL(&entity.Columns[i])
	}
	return fmt.Sprintf(sql, entity.GetTableName(), strings.Join(fieldSqls, ","))
}
