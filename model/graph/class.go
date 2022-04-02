package graph

import (
	"rxdrag.com/entity-engine/model/domain"
)

type Class struct {
	Attributes   []*Attribute
	Associations []*Association
	Methods      []*Method
	Domain       *domain.Class
}

func (c *Class) Uuid() string {
	return c.Domain.Uuid
}

func (e *Class) Name() string {
	return e.Domain.Name
}

func (e *Class) Description() string {
	return e.Domain.Description
}
