package model

import (
	"fmt"

	"rxdrag.com/entity-engine/model/domain"
	"rxdrag.com/entity-engine/model/graph"
	"rxdrag.com/entity-engine/model/meta"
	"rxdrag.com/entity-engine/model/table"
	"rxdrag.com/entity-engine/oldmeta"
)

type Model struct {
	Meta   *meta.Model
	Domain *domain.Model
	Grahp  *graph.Model
	Tables []*table.Table

	Enums              []*Enum
	Interfaces         []*Entity
	Entities           []*Entity
	Relations          []*Relation
	InheritedRelations []*InheritedRelation
	OldTables          []*Table
}

func (m *Model) Validate() {
	//检查空实体（除ID外没有属性跟关联）
	for _, entity := range m.Entities {
		if len(entity.Columns) < 1 && len(entity.Associations) < 1 {
			panic(fmt.Sprintf("Entity %s should have one normal field at least", entity.Name))
		}
	}
}

func NewModel(c *oldmeta.MetaContent) *Model {
	enums, interfaces, entities := c.SplitEntities()
	inherits, relations := c.SplitRelations()

	model := Model{
		Enums:      make([]*Enum, len(enums)),
		Interfaces: make([]*Entity, len(interfaces)),
		Entities:   make([]*Entity, len(entities)),
		Relations:  []*Relation{},
		OldTables:  []*Table{},
	}
	model.buildEnums(enums)
	model.buildInterfaces(interfaces)
	model.buildEntities(entities)
	model.buildInherits(inherits)
	model.buildRelations(relations)
	model.buildColumns()
	model.buildTables()
	return &model
}

func (model *Model) buildEnums(metas []*oldmeta.EntityMeta) {
	for i := range metas {
		model.Enums[i] = &Enum{
			EntityMeta: *metas[i],
			model:      model,
		}
	}
}

func (model *Model) buildInterfaces(metas []*oldmeta.EntityMeta) {
	for i := range metas {
		entity := &Entity{
			EntityMeta:   *metas[i],
			Associations: map[string]*Association{},
			Children:     []*Entity{},
			model:        model,
		}

		entity.Columns = mapColumns(metas[i].Columns, entity, model)
		model.Interfaces[i] = entity
	}
}

func (model *Model) buildEntities(metas []*oldmeta.EntityMeta) {
	for i := range metas {
		entity := &Entity{
			EntityMeta:   *metas[i],
			Interfaces:   []*Entity{},
			Associations: map[string]*Association{},
			model:        model,
		}
		entity.Columns = mapColumns(metas[i].Columns, entity, model)
		model.Entities[i] = entity
	}
}

func (model *Model) buildInherits(relations []*oldmeta.RelationMeta) {
	for i := range relations {
		relation := relations[i]

		sourceEntity := model.GetEntityByUuid(relation.SourceId)
		if sourceEntity == nil {
			panic("Can not find entity by relation:" + relation.SourceId)
		}
		interfaceEntity := model.GetInterfaceByUuid(relation.TargetId)
		if interfaceEntity == nil {
			panic("Can not find interface by relation:" + relation.TargetId)
		}

		sourceEntity.Interfaces = append(sourceEntity.Interfaces, interfaceEntity)
		interfaceEntity.Children = append(interfaceEntity.Children, sourceEntity)
	}
}
func (model *Model) buildRelations(relations []*oldmeta.RelationMeta) {
	for i := range relations {
		relation := relations[i]

		sourceEntities := []*Entity{}
		targetEntities := []*Entity{}

		sourceInterface := model.GetInterfaceByUuid(relation.SourceId)
		if sourceInterface != nil {
			sourceEntities = append(sourceEntities, sourceInterface.Children...)
		} else {
			sourceEntity := model.GetEntityByUuid(relation.SourceId)
			if sourceEntity == nil {
				panic("Can not find entity by relation source:" + relation.SourceId)
			}
			sourceEntities = append(sourceEntities, sourceEntity)
		}

		targetInterface := model.GetInterfaceByUuid(relation.TargetId)

		if targetInterface != nil {
			targetEntities = append(targetEntities, targetInterface.Children...)
		} else {
			targetEntity := model.GetEntityByUuid(relation.TargetId)
			if targetEntity == nil {
				panic("Can not find entity by relation source:" + relation.TargetId)
			}
			targetEntities = append(targetEntities, targetEntity)
		}

		if len(sourceEntities) == 1 && len(targetEntities) == 1 {
			src := sourceEntities[0]
			tar := targetEntities[0]

			newRelationMeta := *relation
			newRelationMeta.SourceId = src.Uuid
			newRelationMeta.TargetId = tar.Uuid
			model.Relations = append(model.Relations, &Relation{
				RelationMeta: newRelationMeta,
				model:        model,
			})
			model.decomposeRelation(src, tar, &newRelationMeta)
		} else {
			//根据继承关系，创建新的关联
			for i := range sourceEntities {
				src := sourceEntities[i]
				for j := range targetEntities {
					tar := targetEntities[j]
					newRelationMeta := *relation
					newRelationMeta.Uuid = fmt.Sprintf("%s-%s-%s", relation.Uuid, src.Uuid, tar.Uuid)
					newRelationMeta.SourceId = src.Uuid
					newRelationMeta.TargetId = tar.Uuid
					model.decomposeRelation(src, tar, &newRelationMeta)
				}
			}
		}
	}
}

func (model *Model) decomposeRelation(src *Entity, tar *Entity, relation *oldmeta.RelationMeta) {
	src.Associations[relation.RoleOfTarget] = &Association{
		Name: relation.RoleOfTarget,
		Relation: &Relation{
			RelationMeta: *relation,
			model:        model,
		},
		TypeEntity:  model.GetEntityByUuid(relation.TargetId),
		OfEntity:    model.GetEntityByUuid(relation.SourceId),
		Description: relation.DescriptionOnSource,
	}
	tar.Associations[relation.RoleOfSource] = &Association{
		Name: relation.RoleOfSource,
		Relation: &Relation{
			RelationMeta: *relation,
			model:        model,
		},
		TypeEntity:  model.GetEntityByUuid(relation.SourceId),
		OfEntity:    model.GetEntityByUuid(relation.TargetId),
		Description: relation.DescriptionOnTarget,
	}
}

func (model *Model) buildColumns() {
	for i := range model.Entities {
		model.Entities[i].makeColumns()
	}
}

func (model *Model) buildTables() {
	for i := range model.Entities {
		model.OldTables = append(model.OldTables, model.Entities[i].Table())
	}

	for i := range model.Relations {
		relation := model.Relations[i]
		if relation.RelationType != oldmeta.IMPLEMENTS {
			relationTable := relation.Table()
			model.OldTables = append(model.OldTables, relationTable)
		}
	}
}

func (m *Model) GetEnumByUuid(uuid string) *Enum {
	for i := range m.Enums {
		enum := m.Enums[i]
		if enum.Uuid == uuid {
			return enum
		}
	}
	return nil
}

func (m *Model) GetInterfaceByUuid(uuid string) *Entity {
	for i := range m.Interfaces {
		intf := m.Interfaces[i]
		if intf.Uuid == uuid {
			return intf
		}
	}
	return nil
}

func (m *Model) GetEntityByUuid(uuid string) *Entity {
	for i := range m.Entities {
		entity := m.Entities[i]
		if entity.Uuid == uuid {
			return entity
		}
	}
	return nil
}

func (m *Model) GetEntityOrInterfaceByName(name string) *Entity {
	for i := range m.Interfaces {
		intf := m.Interfaces[i]
		if intf.Name == name {
			return intf
		}
	}
	for i := range m.Entities {
		entity := m.Entities[i]
		if entity.Name == name {
			return entity
		}
	}
	return nil
}

func (m *Model) GetMetaEntity() *Entity {
	return m.GetEntityByUuid(MetaEntity.Uuid)
}

var TheModel *Model
