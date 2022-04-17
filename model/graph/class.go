package graph

import (
	"rxdrag.com/entity-engine/model/domain"
	"rxdrag.com/entity-engine/utils"
)

type Class struct {
	attributes   []*Attribute
	associations []*Association
	methods      []*Method
	Domain       *domain.Class
}

func NewClass(c *domain.Class) *Class {
	cls := Class{
		Domain:     c,
		attributes: make([]*Attribute, len(c.Attributes)),
		methods:    make([]*Method, len(c.Methods)),
	}

	for i := range c.Attributes {
		cls.attributes[i] = NewAttribute(c.Attributes[i], &cls)
	}

	for i := range c.Methods {
		cls.methods[i] = NewMethod(c.Methods[i], &cls)
	}

	return &cls
}

func (c *Class) Uuid() string {
	return c.Domain.Uuid
}

func (c *Class) InnerId() uint64 {
	return c.Domain.InnerId
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

// func (c *Class) Attributes() []*Attribute {
// 	return c.attributes
// }

// func (c *Class) Associations() []*Association {
// 	return c.associations
// }

func (c *Class) MethodsByType(operateType string) []*Method {
	methods := []*Method{}
	for i := range c.methods {
		method := c.methods[i]
		if method.Method.OperateType == operateType {
			methods = append(methods, method)
		}
	}

	return methods
}

func (c *Class) IsSoftDelete() bool {
	return c.Domain.SoftDelete
}
