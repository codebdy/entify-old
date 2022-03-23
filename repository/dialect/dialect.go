package dialect

import "rxdrag.com/entity-engine/meta"

const (
	MySQL = "mysql"
)

type SQLBuilder interface {
	BuildBoolExp(where map[string]interface{}) (string, []interface{})
	BuildFieldExp(fieldName string, fieldArgs map[string]interface{}) (string, []interface{})

	BuildCreateTableSQL(table *meta.Table) string
	BuildDeleteTableSQL(table *meta.Table) string
	BuildColumnSQL(column *meta.ColumnMeta) string
	BuildModifyTableAtoms(diff *meta.TableDiff) []meta.ModifyAtom
	ColumnTypeSQL(column *meta.ColumnMeta) string

	BuildInsertSQL(object map[string]interface{}, entity *meta.EntityMeta) (string, []interface{})
	BuildUpdateSQL(object map[string]interface{}, entity *meta.EntityMeta) (string, []interface{})
}

func GetSQLBuilder() SQLBuilder {
	var builder MySQLBuilder
	return &builder
}
