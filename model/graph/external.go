package graph

import "rxdrag.com/entify/model/domain"

type External struct {
	Entity
}

func NewExternal(c *domain.Class) *External {
	return &External{
		Entity: Entity{
			Class: *NewClass(c),
		},
	}
}
