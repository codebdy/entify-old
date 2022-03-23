package model

import (
	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/meta"
)

type Association struct {
	Name        string
	Relation    *Relation
	OfEntity    *Entity
	TypeEntity  *Entity
	Description string
}

type Entity struct {
	meta.EntityMeta
	Parent       *Entity
	Children     []*Entity
	Associations []*Association
	model        *Model
}

type Relation struct {
	meta.RelationMeta
	model *Model
}

type Model struct {
	Entities  []*Entity
	Relations []*Relation
	Tables    []*Table
}

func NewModel(c *meta.MetaContent) *Model {
	model := Model{
		Entities:  make([]*Entity, len(c.Entities)),
		Relations: []*Relation{},
		Tables:    entityTables(c),
	}

	for i := range c.Entities {
		model.Entities[i] = &Entity{
			EntityMeta:   c.Entities[i],
			Parent:       nil,
			Children:     []*Entity{},
			Associations: []*Association{},
			model:        &model,
		}
	}

	for i := range c.Relations {
		relation := c.Relations[i]
		if relation.RelationType != meta.IMPLEMENTS {
			model.Relations = append(model.Relations, &Relation{
				RelationMeta: relation,
				model:        &model,
			})
		}
	}

	for i := range c.Relations {
		relation := c.Relations[i]
		if relation.RelationType == meta.IMPLEMENTS {
		}
	}

	for i := range c.Relations {
		relation := c.Relations[i]
		if relation.RelationType != meta.IMPLEMENTS {
			relationTable := relationTable(c, &relation)
			model.Tables = append(model.Tables, relationTable)
		}
	}
	return &model
}

func (m *Model) FindEntityByUuid(uuid string) *Entity {
	for i := range m.Entities {
		entity := m.Entities[i]
		if entity.Uuid == uuid {
			return entity
		}
	}
	return nil
}

func FindTable(metaUuid string, tables []*Table) *Table {
	for i := range tables {
		if tables[i].MetaUuid == metaUuid {
			return tables[i]
		}
	}
	return nil
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

func relationTable(c *meta.MetaContent, relation *meta.RelationMeta) *Table {
	table := &Table{
		MetaUuid: relation.Uuid,
		Name:     c.RelationTableName(relation),
		Columns: []meta.ColumnMeta{
			{
				Name:  relation.RelationSourceColumnName(),
				Type:  meta.COLUMN_ID,
				Uuid:  relation.Uuid + consts.SUFFIX_SOURCE,
				Index: true,
			},
			{
				Name:  relation.RelationTargetColumnName(),
				Type:  meta.COLUMN_ID,
				Uuid:  relation.Uuid + consts.SUFFIX_TARGET,
				Index: true,
			},
		},
	}
	table.Columns = append(table.Columns, relation.Columns...)

	return table
}
