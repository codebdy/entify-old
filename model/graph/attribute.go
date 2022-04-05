package graph

import "rxdrag.com/entity-engine/model/domain"

type Attribute struct {
	domain.Attribute
	Class     *Class
	EumnType  *Enum
	EnityType *Entity
}

func NewAttribute(a *domain.Attribute, c *Class) *Attribute {
	return &Attribute{
		Attribute: *a,
		Class:     c,
	}
}
