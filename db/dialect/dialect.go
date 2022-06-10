package dialect

import (
	"rxdrag.com/entify/db/dialect/mysql"
	"rxdrag.com/entify/model"
	"rxdrag.com/entify/model/data"
	"rxdrag.com/entify/model/graph"
	"rxdrag.com/entify/model/table"
)

const (
	MySQL = "mysql"
)

type SQLBuilder interface {
	BuildMeSQL() string
	BuildRolesSQL() string
	BuildLoginSQL() string
	BuildCreateMetaSQL() string
	BuildCreateAbilitySQL() string
	BuildCreateEntityAuthSettingsSQL() string
	BuildBoolExp(argEntity *graph.ArgEntity, where map[string]interface{}) (string, []interface{})
	BuildFieldExp(fieldName string, fieldArgs map[string]interface{}) (string, []interface{})

	BuildCreateTableSQL(table *table.Table) string
	BuildDeleteTableSQL(table *table.Table) string
	BuildColumnSQL(column *table.Column) string
	BuildModifyTableAtoms(diff *model.TableDiff) []model.ModifyAtom
	ColumnTypeSQL(column *table.Column) string

	BuildQuerySQLBody(argEntity *graph.ArgEntity, fields []*graph.Attribute) string
	BuildWhereSQL(argEntity *graph.ArgEntity, fields []*graph.Attribute, where map[string]interface{}) (string, []interface{})
	BuildOrderBySQL(argEntity *graph.ArgEntity, orderBy interface{}) string
	//BuildQuerySQL(tableName string, fields []*graph.Attribute, args map[string]interface{}) (string, []interface{})

	BuildInsertSQL(fields []*data.Field, table *table.Table) string
	BuildUpdateSQL(id uint64, fields []*data.Field, table *table.Table) string

	BuildQueryByIdsSQL(entity *graph.Entity, idCounts int) string
	BuildClearAssociationSQL(ownerId uint64, tableName string, ownerFieldName string) string
	BuildQueryAssociatedInstancesSQL(entity *graph.Entity,
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

	BuildTableCheckSQL(name string, database string) string
}

func GetSQLBuilder() SQLBuilder {
	var builder mysql.MySQLBuilder
	return &builder
}
