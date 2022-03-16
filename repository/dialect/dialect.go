package dialect

import "rxdrag.com/entity-engine/meta"

const (
	MySQL = "mysql"
)

type SQLBuilder interface {
	BuildBoolExp(where map[string]interface{}) (string, []interface{})
	BuildFieldExp(fieldName string, fieldArgs map[string]interface{}) (string, []interface{})

	BuildCreateTableSQL(table *meta.Table) (string, string)
	BuildColumnSQL(column *meta.Column) string
	ColumnTypeSQL(column *meta.Column) string
}

func GetSQLBuilder() SQLBuilder {
	var builder MySQLBuilder
	return &builder
}
