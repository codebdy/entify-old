package graph

import "rxdrag.com/entify/model/domain"

type External struct {
	Class
}

func NewExternal(c *domain.Class) *External {
	return &External{
		Class: *NewClass(c),
	}
}
