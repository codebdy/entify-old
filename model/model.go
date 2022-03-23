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

type Enum struct {
	meta.EntityMeta
	model *Model
}

type Model struct {
	Enums      []*Enum
	Interfaces []*Interface
	Entities   []*Entity
	Relations  []*Relation
	Tables     []*Table
}

func NewModel(c *meta.MetaContent) *Model {
	enums, interfaces, entities := c.SplitEntities()
	inherits, relations := c.SplitRelations()

	model := Model{
		Enums:      make([]*Enum, len(enums)),
		Interfaces: make([]*Interface, len(interfaces)),
		Entities:   make([]*Entity, len(entities)),
		Relations:  []*Relation{},
		Tables:     entityTables(c),
	}
	model.buildEnums(enums)
	model.buildInterfaces(interfaces)
	model.buildEntities(entities)
	model.buildInherits(inherits)
	model.buildRelations(relations)
	model.buildTables()
	return &model
}

func (model *Model) buildEnums(metas []*meta.EntityMeta) {
	for i := range metas {
		model.Enums[i] = &Enum{
			EntityMeta: *metas[i],
			model:      model,
		}
	}
}

func (model *Model) buildInterfaces(metas []*meta.EntityMeta) {
	for i := range metas {
		model.Interfaces[i] = &Interface{
			EntityMeta:   *metas[i],
			Associations: []*Association{},
			Children:     []*Entity{},
			model:        model,
		}
	}
}

func (model *Model) buildEntities(metas []*meta.EntityMeta) {
	for i := range metas {
		model.Entities[i] = &Entity{
			EntityMeta:   *metas[i],
			Interfaces:   []*Interface{},
			Associations: []*Association{},
			model:        model,
		}
	}
}

func (model *Model) buildInherits(relations []*meta.RelationMeta) {
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
func (model *Model) buildRelations(relations []*meta.RelationMeta) {
	// for i := range relations {
	// 	relation := relations[i]
	// 	model.Relations = append(model.Relations, &Relation{
	// 		RelationMeta: *relation,
	// 		model:        model,
	// 	})
	// }
}

func (model *Model) buildTables() {
	for i := range model.Relations {
		relation := model.Relations[i]
		if relation.RelationType != meta.IMPLEMENTS {
			relationTable := relation.Table()
			model.Tables = append(model.Tables, relationTable)
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

func (m *Model) GetInterfaceByUuid(uuid string) *Interface {
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
