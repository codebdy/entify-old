package graph

import "rxdrag.com/entity-engine/model/domain"

type Interface struct {
	Class
	Children []*Entity
}

func NewInterface(c *domain.Class) *Interface {
	return &Interface{
		Class: *NewClass(c),
	}
}
