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

	return &model
}

func (m *Model) Validate() {
	//检查空实体（除ID外没有属性跟关联）
	for _, entity := range m.Entities {
		if len(entity.Attributes) < 1 && len(entity.Associations) < 1 {
			panic(fmt.Sprintf("Entity %s should have one normal field at least", entity.Name()))
		}
	}
}

func (m *Model) RootEnities() {

}

func (m *Model) RootInterfaces() {

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
