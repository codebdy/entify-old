package model

import "rxdrag.com/entity-engine/meta"

type Table struct {
	MetaUuid string
	Name     string
	Columns  []meta.ColumnMeta
}
