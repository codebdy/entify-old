package model

import "rxdrag.com/entity-engine/meta"

type Association struct {
	Name        string
	Relation    *Relation
	OfEntity    *Entity
	TypeEntity  *Entity
	Description string
}

type Enum struct {
	meta.EntityMeta
	model *Model
}

type Entity struct {
	meta.EntityMeta
	Associations map[string]*Association
	Interfaces   []*Interface
	model        *Model
}

func entityTables(c *meta.MetaContent) []*Table {

	normalEntities := c.FilterEntity(func(e *meta.EntityMeta) bool {
		return e.HasTable()
	})

	tables := make([]*Table, len(normalEntities))

	for i := range normalEntities {
		entity := normalEntities[i]
		table := &Table{Name: entity.GetTableName(), MetaUuid: entity.Uuid}
		table.Columns = append(table.Columns, entity.Columns...)
		tables[i] = table
	}

	return tables
}
