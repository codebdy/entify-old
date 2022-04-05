package graph

import (
	"rxdrag.com/entity-engine/model/domain"
)

type Method struct {
	Method *domain.Method
	Class  *Class
}

func NewMethod(m *domain.Method, c *Class) *Method {
	return &Method{
		Method: m,
		Class:  c,
	}
}

func (c *Method) Uuid() string {
	return c.Method.Uuid
}

func (e *Method) Name() string {
	return e.Method.Name
}
