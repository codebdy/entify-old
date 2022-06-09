package mysql

import (
	"fmt"
	"time"

	"rxdrag.com/entify/consts"
	"rxdrag.com/entify/model/data"
)

func (b *MySQLBuilder) BuildDeleteSQL(id uint64, tableName string) string {
	sql := fmt.Sprintf(
		"DELETE FROM `%s` WHERE (`%s` = '%d')",
		tableName,
		"id",
		id,
	)
	return sql
}

func (b *MySQLBuilder) BuildSoftDeleteSQL(id uint64, tableName string) string {
	sql := fmt.Sprintf(
		"UPDATE `%s` SET `%s` = '%s' WHERE (`%s` = %d)",
		tableName,
		consts.DELETED_AT,
		time.Now(),
		"id",
		id,
	)
	return sql
}

func (b *MySQLBuilder) BuildDeletePovitSQL(povit *data.AssociationPovit) string {
	return fmt.Sprintf(
		"DELETE FROM `%s` WHERE (`%s` = %d AND `%s` = %d)",
		povit.Table().Name,
		povit.Source.Column.Name,
		povit.Source.Value,
		povit.Target.Column.Name,
		povit.Target.Value,
	)
}

func (b *MySQLBuilder) BuildClearAssociationSQL(ownerId uint64, tableName string, ownerFieldName string) string {
	sql := fmt.Sprintf(
		"DELETE FROM `%s` WHERE (`%s` = '%d')",
		tableName,
		ownerFieldName,
		ownerId,
	)
	return sql
}
