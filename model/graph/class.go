package graph

import (
	"rxdrag.com/entity-engine/model/domain"
	"rxdrag.com/entity-engine/utils"
)

type Class struct {
	attributes   []*Attribute
	associations []*Association
	Methods      []*Method
	Domain       *domain.Class
}

func NewClass(c *domain.Class) *Class {
	return &Class{
		Domain: c,
	}
}

func (c *Class) Uuid() string {
	return c.Domain.Uuid
}

func (c *Class) Name() string {
	return c.Domain.Name
}

func (e *Class) Description() string {
	return e.Domain.Description
}

func (e *Class) AddAssociation(a *Association) {
	e.associations = append(e.associations, a)
}

func (c *Class) TableName() string {
	return utils.SnakeString(c.Domain.Name)
}

func (c *Class) Attributes() []*Attribute {
	return c.attributes
}

func (c *Class) Associations() []*Association {
	return c.associations
}
