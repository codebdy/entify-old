package graph

import (
	"fmt"

	"rxdrag.com/entity-engine/model/domain"
	"rxdrag.com/entity-engine/model/meta"
	"rxdrag.com/entity-engine/model/table"
)

type Model struct {
	Enums      []*Enum
	Interfaces []*Interface
	Entities   []*Entity
	Relations  []*Relation
	Tables     []*table.Table
}

func New(m *domain.Model) *Model {
	model := Model{}

	for i := range m.Enums {
		model.Enums = append(model.Enums, NewEnum(m.Enums[i]))
	}

	//构建所有接口
	for i := range m.Classes {
		cls := m.Classes[i]
		if cls.StereoType == meta.CLASSS_ABSTRACT {
			model.Interfaces = append(model.Interfaces, NewInterface(cls))
		} else if cls.StereoType == meta.CLASSS_ENTITY && cls.HasChildren() {
			intf := NewInterface(cls)
			entity := NewEntity(cls)
			intf.Children = append(intf.Children, entity)
			entity.Interfaces = append(entity.Interfaces, intf)
			model.Interfaces = append(model.Interfaces, intf)
			model.Entities = append(model.Entities, entity)
		}
	}

	//构建所有实体
	for i := range m.Classes {
		cls := m.Classes[i]
		if (cls.StereoType == meta.CLASSS_ENTITY && !cls.HasChildren()) ||
			cls.StereoType == meta.CLASS_VALUE_OBJECT ||
			cls.StereoType == meta.CLASS_SERVICE {
			entity := NewEntity(cls)
			model.Entities = append(model.Entities, entity)
		}
	}

	//处理所有关联
	for i := range m.Relations {
		relation := m.Relations[i]
		source := model.GetNodeByUuid(relation.Source.Uuid)
		target := model.GetNodeByUuid(relation.Target.Uuid)
		r := NewRelation(relation, source, target)
		model.Relations = append(model.Relations, r)
		source.AddAssociation(NewAssociation(r, source.Uuid()))
		target.AddAssociation(NewAssociation(r, target.Uuid()))
	}

	//处理Table

	return &model
}

func (m *Model) Validate() {
	//检查空实体（除ID外没有属性跟关联）
	for _, entity := range m.Entities {
		if entity.IsEmperty() {
			panic(fmt.Sprintf("Entity %s should have one normal field at least", entity.Name()))
		}
	}
}

func (m *Model) RootEnities() {

}

func (m *Model) RootInterfaces() {

}

func (m *Model) GetNodeByUuid(uuid string) Node {
	intf := m.GetInterfaceByUuid(uuid)

	if intf != nil {
		return intf
	}

	return m.GetEntityByUuid(uuid)
}

func (m *Model) GetInterfaceByUuid(uuid string) *Interface {
	for i := range m.Interfaces {
		intf := m.Interfaces[i]
		if intf.Uuid() == uuid {
			return intf
		}
	}
	return nil
}

func (m *Model) GetEntityByUuid(uuid string) *Entity {
	for i := range m.Entities {
		ent := m.Entities[i]
		if ent.Uuid() == uuid {
			return ent
		}
	}
	return nil
}

/*
处理枚举
	for i := range model.Classes {
		cls := model.Classes[i]
		for j := range cls.Attributes {
			attr := cls.Attributes[j]
			if attr.Type == meta.ENUM || attr.Type == meta.ENTITY_ARRAY {
				attr.
			}
		}
	}
*/
