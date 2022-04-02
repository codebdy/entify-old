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

func NewClass(c *domain.Class) *Class {
	return &Class{
		Domain: c,
	}
}

func (e *Class) Description() string {
	return e.Domain.Description
}

func (e *Class) AddAssociation(a *Association) {
	e.Associations = append(e.Associations, a)
}

//包含继承来的
func (e *Class) AllAttributes() []*Attribute {
	return []*Attribute{}
}

//包含继承来的
func (e *Class) AllAssociations() []*Association {
	return []*Association{}
}

func (e *Class) IsEmperty() bool {
	return len(e.AllAttributes()) < 1 && len(e.AllAssociations()) < 1
}
