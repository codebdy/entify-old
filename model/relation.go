package model

import (
	"fmt"

	"rxdrag.com/entity-engine/config"
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
	tableName := relation.SouceTableName() +
		"_" + utils.SnakeString(relation.RoleOnSource) +
		"_" + relation.TargetTableName() +
		"_" + utils.SnakeString(relation.RoleOnTarget) +
		consts.SUFFIX_PIVOT

	if len([]rune(tableName)) >= config.TABLE_NAME_MAX_LENGTH {
		tableName = string([]byte(tableName)[:config.TABLE_NAME_MAX_LENGTH-20])
		tableName = fmt.Sprintf("%s%s_%d_%d", tableName, consts.SUFFIX_PIVOT, relation.SouceInnerId(), relation.TargetInnerId())
	}

	return tableName
}

func (relation *Relation) SouceTableName() string {
	sourceEntity := relation.model.GetEntityByUuid(relation.SourceId)
	return sourceEntity.GetTableName()
}

func (relation *Relation) TargetTableName() string {
	targetEntity := relation.model.GetEntityByUuid(relation.TargetId)
	return targetEntity.GetTableName()
}

func (relation *Relation) SouceInnerId() uint {
	sourceEntity := relation.model.GetEntityByUuid(relation.SourceId)
	return sourceEntity.InnerId
}

func (relation *Relation) TargetInnerId() uint {
	targetEntity := relation.model.GetEntityByUuid(relation.TargetId)
	return targetEntity.InnerId
}
