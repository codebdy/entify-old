package model

import (
	"fmt"

	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/meta"
)

type Relation struct {
	meta.RelationMeta
	model *Model
}

type InheritedRelation struct {
	Relation
	InheritFrom *Relation
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
	table.Columns = append(table.Columns, mapColumns(relation.Columns, nil, relation.model)...)

	return table
}

func (relation *Relation) TableName() string {
	tableName := fmt.Sprintf(consts.PIVOT+"_%d_%d_%d", relation.SouceInnerId(), relation.TargetInnerId(), relation.InnerId)
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

func (relation *Relation) SouceInnerId() uint64 {
	sourceEntity := relation.model.GetEntityByUuid(relation.SourceId)
	return sourceEntity.InnerId
}

func (relation *Relation) TargetInnerId() uint64 {
	targetEntity := relation.model.GetEntityByUuid(relation.TargetId)
	return targetEntity.InnerId
}

func (r *Relation) RelationSourceColumnName() string {
	return r.SouceTableName() + consts.ID_SUFFIX
}

func (r *Relation) RelationTargetColumnName() string {
	return r.TargetTableName() + consts.ID_SUFFIX
}
