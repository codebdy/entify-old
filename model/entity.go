package model

import (
	"rxdrag.com/entity-engine/meta"
	"rxdrag.com/entity-engine/utils"
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
	Associations map[string]*Association
	Columns      []*Column
	Interfaces   []*Interface
	model        *Model
}

func (entity *Entity) ColumnNames() []string {
	names := make([]string, len(entity.Columns))

	for i, column := range entity.Columns {
		names[i] = column.Name
	}

	return names
}

func (entity *Entity) GetColumn(name string) *Column {
	for _, column := range entity.Columns {
		if column.Name == name {
			return column
		}
	}

	return nil
}

func (entity *Entity) GetTableName() string {
	if (*entity).TableName != "" {
		return (*entity).TableName
	}
	return utils.SnakeString((*entity).Name)
}

func (entity *Entity) Table() *Table {
	table := &Table{Name: entity.GetTableName(), MetaUuid: entity.Uuid}
	table.Columns = append(table.Columns, entity.Columns...)
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
	entity.Columns = columns
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
