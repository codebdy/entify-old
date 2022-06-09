package mysql

import (
	"fmt"
	"strings"

	"rxdrag.com/entify/model/data"
	"rxdrag.com/entify/model/table"
)

func (b *MySQLBuilder) BuildUpdateSQL(id uint64, fields []*data.Field, table *table.Table) string {
	sql := fmt.Sprintf(
		"UPDATE `%s` SET %s WHERE ID = %d",
		table.Name,
		updateSetFields(fields),
		id,
	)

	return sql
}

func updateSetFields(fields []*data.Field) string {
	if len(fields) == 0 {
		panic("No update fields")
	}
	newKeys := make([]string, len(fields))
	for i, field := range fields {
		newKeys[i] = field.Column.Name + "=?"
	}
	return strings.Join(newKeys, ",")
}
