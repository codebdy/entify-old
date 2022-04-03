package graph

import "rxdrag.com/entity-engine/model/table"

func NewEntityTable(entity *Entity) *table.Table {
	var table table.Table

	entity.Table = &table
	return &table
}

func NewRelationTables(relation *Relation) []*table.Table {
	var tables []*table.Table

	return tables
}
