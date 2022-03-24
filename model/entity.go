package model

import (
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
	Columns      []*Column
	Associations map[string]*Association
	AllColumns   []*Column
	Interfaces   []*Interface
	model        *Model
}

func (entity *Entity) Table() *Table {
	table := &Table{Name: entity.GetTableName(), MetaUuid: entity.Uuid}
	table.Columns = append(table.Columns, entity.AllColumns...)
	return table
}

func (entity *Entity) makeColumns() {
	entity.Columns = mapColumns(entity.EntityMeta.Columns, entity.model)
	columns := entity.Columns
	for i := range entity.Interfaces {
		intf := entity.Interfaces[i]
		for _, column := range intf.Columns {
			if !findColumnByName(column.Name, columns) {
				columns = append(columns, column)
			}
		}
	}
	entity.AllColumns = columns
}

func findColumnByName(name string, columns []*Column) bool {
	for _, column := range columns {
		if column.Name == name {
			return true
		}
	}
	return false
}

func (a *Association) IsArray() bool {
	if a.Relation.RelationType == meta.ONE_TO_MANY {
		if a.OfEntity.Uuid == a.Relation.SourceId {
			return true
		}
	} else if a.Relation.RelationType == meta.MANY_TO_ONE {
		if a.OfEntity.Uuid == a.Relation.TargetId {
			return true
		}
	} else if a.Relation.RelationType == meta.MANY_TO_MANY {
		return true
	}
	return false
}
