package model

import "rxdrag.com/entity-engine/meta"

type Column struct {
	meta.ColumnMeta
	model *Model
}

func (c *Column) GetEnum() *Enum {
	return c.model.GetEnumByUuid(c.EnumUuid)
}

func mapColumns(metas []meta.ColumnMeta, model *Model) []*Column {

	columns := make([]*Column, len(metas))

	for i := range metas {
		columns[i] = &Column{ColumnMeta: metas[i], model: model}
	}

	return columns
}
