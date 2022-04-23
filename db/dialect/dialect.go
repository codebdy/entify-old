package dialect

import (
	"rxdrag.com/entity-engine/model"
	"rxdrag.com/entity-engine/model/data"
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
	BuildModifyTableAtoms(diff *model.TableDiff) []model.ModifyAtom
	ColumnTypeSQL(column *table.Column) string

	BuildQuerySQL(entity graph.Noder, args map[string]interface{}) (string, []interface{})

	BuildInsertSQL(fields []*data.Field, table *table.Table) string
	BuildUpdateSQL(id uint64, fields []*data.Field, table *table.Table) string

	BuildClearAssociationSQL(ownerId uint64, tableName string, ownerFieldName string) string
	BuildQueryAssociatedInstancesSQL(node graph.Noder,
		ownerId uint64,
		povitTableName string,
		ownerFieldName string,
		typeFieldName string,
	) string
	BuildBatchAssociationSQL(
		tableName string,
		fields []*graph.Attribute,
		ids []uint64,
		povitTableName string,
		ownerFieldName string,
		typeFieldName string,
	) string
	BuildDeleteSQL(id uint64, tableName string) string
	BuildSoftDeleteSQL(id uint64, tableName string) string

	BuildQueryPovitSQL(povit *data.AssociationPovit) string
	BuildInsertPovitSQL(povit *data.AssociationPovit) string
	BuildDeletePovitSQL(povit *data.AssociationPovit) string
}

func GetSQLBuilder() SQLBuilder {
	var builder MySQLBuilder
	return &builder
}
