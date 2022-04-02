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
			model.Interfaces = append(model.Interfaces, NewInterface(cls))
		}
	}

	//构建所有实体
	for i := range m.Classes {
		cls := m.Classes[i]
		if cls.StereoType == meta.CLASSS_ENTITY ||
			cls.StereoType == meta.CLASS_VALUE_OBJECT ||
			cls.StereoType == meta.CLASS_SERVICE {
			model.Entities = append(model.Entities, NewEntity(cls))
		}
	}

	//处理所有管理
	for i := range m.Relations {
		relation := m.Relations[i]
		s := model.GetNodeByUuid(relation.Source.Uuid)
		t := model.GetNodeByUuid(relation.Target.Uuid)
		model.Relations = append(model.Relations, NewRelation(relation, s, t))
	}

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
