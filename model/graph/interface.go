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

func (i *Interface) Uuid() string {
	return i.Domain.Uuid
}

func (i *Interface) Name() string {
	return i.Domain.Name
}

func (i *Interface) isInterface() bool {
	return true
}
func (i *Interface) Interface() *Interface {
	return i
}
func (i *Interface) Entity() *Entity {
	return nil
}
