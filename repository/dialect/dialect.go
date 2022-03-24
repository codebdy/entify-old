package dialect

import (
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
	BuildColumnSQL(column *model.Column) string
	BuildModifyTableAtoms(diff *model.TableDiff) []model.ModifyAtom
	ColumnTypeSQL(column *model.Column) string

	BuildInsertSQL(object map[string]interface{}, entity *model.Entity) (string, []interface{})
	BuildUpdateSQL(object map[string]interface{}, entity *model.Entity) (string, []interface{})
}

func GetSQLBuilder() SQLBuilder {
	var builder MySQLBuilder
	return &builder
}
