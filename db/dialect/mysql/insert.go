package mysql

import (
	"fmt"
	"strings"

	"rxdrag.com/entify/model/data"
	"rxdrag.com/entify/model/table"
)

func (b *MySQLBuilder) BuildInsertSQL(fields []*data.Field, table *table.Table) string {
	sql := fmt.Sprintf("INSERT INTO `%s`(%s) VALUES(%s)", table.Name, insertFields(fields), insertValueSymbols(fields))

	return sql
}

func (b *MySQLBuilder) BuildInsertPovitSQL(povit *data.AssociationPovit) string {
	return fmt.Sprintf(
		"INSERT INTO `%s`(%s,%s) VALUES(%d, %d)",
		povit.Table().Name,
		povit.Source.Column.Name,
		povit.Target.Column.Name,
		povit.Source.Value,
		povit.Target.Value,
	)
}

func insertFields(fields []*data.Field) string {
	names := make([]string, len(fields))
	for i := range fields {
		names[i] = fields[i].Column.Name
	}
	return strings.Join(names, ",")
}

func insertValueSymbols(fields []*data.Field) string {
	array := make([]string, len(fields))
	for i := range array {
		array[i] = "?"
	}
	return strings.Join(array, ",")
}
