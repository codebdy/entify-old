package model

import (
	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/meta"
	"rxdrag.com/entity-engine/utils"
)

type Relation struct {
	meta.RelationMeta
	model *Model
}

func (relation *Relation) Table() *Table {
	table := &Table{
		MetaUuid: relation.Uuid,
		Name:     relation.TableName(),
		Columns: []*Column{

			{
				ColumnMeta: meta.ColumnMeta{
					Name:  relation.RelationSourceColumnName(),
					Type:  meta.COLUMN_ID,
					Uuid:  relation.Uuid + consts.SUFFIX_SOURCE,
					Index: true,
				},
				model: relation.model,
			},
			{
				ColumnMeta: meta.ColumnMeta{
					Name:  relation.RelationTargetColumnName(),
					Type:  meta.COLUMN_ID,
					Uuid:  relation.Uuid + consts.SUFFIX_TARGET,
					Index: true,
				},
				model: relation.model,
			},
		},
	}
	table.Columns = append(table.Columns, mapColumns(relation.Columns, relation.model)...)

	return table
}

func (relation *Relation) TableName() string {
	return relation.SouceTableName() +
		"_" + utils.SnakeString(relation.RoleOnSource) +
		"_" + relation.TargetTableName() +
		"_" + utils.SnakeString(relation.RoleOnTarget) +
		consts.SUFFIX_PIVOT
}

func (relation *Relation) SouceTableName() string {
	sourceEntity := relation.model.GetEntityByUuid(relation.SourceId)
	return sourceEntity.GetTableName()
}

func (relation *Relation) TargetTableName() string {
	targetEntity := relation.model.GetEntityByUuid(relation.TargetId)
	return targetEntity.GetTableName()
}
