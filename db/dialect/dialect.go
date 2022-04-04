package dialect

import (
	"rxdrag.com/entity-engine/model/graph"
	"rxdrag.com/entity-engine/model/table"
)

const (
	MySQL = "mysql"
)

type SQLBuilder interface {
	BuildBoolExp(where map[string]interface{}) (string, []interface{})
	BuildFieldExp(fieldName string, fieldArgs map[string]interface{}) (string, []interface{})

	BuildCreateTableSQL(table *table.Table) string
	BuildDeleteTableSQL(table *table.Table) string
	BuildColumnSQL(column *table.Column) string
	BuildModifyTableAtoms(diff *table.TableDiff) []table.ModifyAtom
	AttributeTypeSQL(column *table.Column) string

	BuildQuerySQL(entity graph.Node, args map[string]interface{}) (string, []interface{})

	BuildInsertSQL(object map[string]interface{}, entity *graph.Entity) (string, []interface{})
	BuildUpdateSQL(object map[string]interface{}, entity *graph.Entity) (string, []interface{})
}

func GetSQLBuilder() SQLBuilder {
	var builder MySQLBuilder
	return &builder
}
