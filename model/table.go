package model

import "rxdrag.com/entity-engine/meta"

type Table struct {
	MetaUuid string
	Name     string
	Columns  []meta.ColumnMeta
}

func FindTable(metaUuid string, tables []*Table) *Table {
	for i := range tables {
		if tables[i].MetaUuid == metaUuid {
			return tables[i]
		}
	}
	return nil
}
