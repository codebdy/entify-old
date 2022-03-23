package dialect

import (
	"rxdrag.com/entity-engine/meta"
	"rxdrag.com/entity-engine/model"
)

const (
	MySQL = "mysql"
)

type SQLBuilder interface {
	BuildBoolExp(where map[string]interface{}) (string, []interface{})
	BuildFieldExp(fieldName string, fieldArgs map[string]interface{}) (string, []interface{})

	BuildCreateTableSQL(table *model.Table) string
	BuildDeleteTableSQL(table *model.Table) string
	BuildColumnSQL(column *meta.ColumnMeta) string
	BuildModifyTableAtoms(diff *model.TableDiff) []model.ModifyAtom
	ColumnTypeSQL(column *meta.ColumnMeta) string

	BuildInsertSQL(object map[string]interface{}, entity *model.Entity) (string, []interface{})
	BuildUpdateSQL(object map[string]interface{}, entity *model.Entity) (string, []interface{})
}

func GetSQLBuilder() SQLBuilder {
	var builder MySQLBuilder
	return &builder
}
