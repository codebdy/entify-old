package graph

import (
	"fmt"

	"rxdrag.com/entity-engine/model/domain"
)

type Model struct {
	Enums      []*Enum
	Interfaces []*Interface
	Entities   []*Entity
}

func New(m *domain.Model) *Model {
	model := Model{}

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
