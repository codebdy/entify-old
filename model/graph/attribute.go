package graph

import "rxdrag.com/entify/model/domain"

type Attribute struct {
	domain.Attribute
	Class           *Class
	EumnType        *Enum
	EnityType       *Entity
	ValueObjectType *Class
}

func NewAttribute(a *domain.Attribute, c *Class) *Attribute {
	return &Attribute{
		Attribute: *a,
		Class:     c,
	}
}

func (a *Attribute) GetName() string {
	return a.Attribute.Name
}

func (a *Attribute) GetType() string {
	return a.Attribute.Type
}
func (a *Attribute) GetEumnType() *Enum {
	return a.EumnType
}
func (a *Attribute) GetEnityType() *Entity {
	return a.EnityType
}
