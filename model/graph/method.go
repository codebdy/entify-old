package graph

import (
	"rxdrag.com/entify/model/domain"
)

type Method struct {
	Method          *domain.Method
	EumnType        *Enum
	EnityType       *Entity
	ValueObjectType *Class
	Class           *Class
}

func NewMethod(m *domain.Method, c *Class) *Method {
	return &Method{
		Method: m,
		Class:  c,
	}
}

func (m *Method) Uuid() string {
	return m.Method.Uuid
}

func (m *Method) GetName() string {
	return m.Method.Name
}

func (m *Method) GetType() string {
	return m.Method.Type
}
func (m *Method) GetEumnType() *Enum {
	return m.EumnType
}
func (m *Method) GetEnityType() *Entity {
	return m.EnityType
}
