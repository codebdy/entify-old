package graph

import (
	"rxdrag.com/entity-engine/model/domain"
)

type Method struct {
	Method *domain.Method
	Entity *Entity
}

func (c *Method) Uuid() string {
	return c.Method.Uuid
}

func (e *Method) Name() string {
	return e.Method.Name
}
