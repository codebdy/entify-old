package graph

import (
	"rxdrag.com/entity-engine/model/table"
	"rxdrag.com/entity-engine/utils"
)

func NewEntityTable(entity *Entity) *table.Table {
	table := &table.Table{
		Uuid: entity.Uuid(),
		Name: utils.SnakeString(entity.Name()),
	}

	for i := range entity.Attributes {
		attr := entity.Attributes[i]
		table.Columns = append(table.Columns, NewAttributeColumn(attr))
	}

	entity.Table = table
	return table
}

func NewAttributeColumn(attr *Attribute) *table.Column {
	return &table.Column{
		AttributeMeta: attr.AttributeMeta,
	}
}

func NewRelationTables(relation *Relation) []*table.Table {
	var tables []*table.Table

	return tables
}
