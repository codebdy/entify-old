package graph

import (
	"rxdrag.com/entify/consts"
	"rxdrag.com/entify/model/domain"
	"rxdrag.com/entify/utils"
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

func (c *Class) Description() string {
	return c.Domain.Description
}

func (c *Class) AddAssociation(a *Association) {
	c.associations = append(c.associations, a)
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

func (c *Class) QueryName() string {
	return utils.FirstLower(c.Name())
}

func (c *Class) QueryOneName() string {
	return consts.ONE + c.Name()
}

func (c *Class) QueryAggregateName() string {
	return utils.FirstLower(c.Name()) + utils.FirstUpper(consts.AGGREGATE)
}

func (c *Class) DeleteName() string {
	return consts.DELETE + c.Name()
}

func (c *Class) DeleteByIdName() string {
	return consts.DELETE + c.Name() + consts.BY_ID
}

func (c *Class) UpdateName() string {
	return utils.FirstLower(c.Name())
}

func (c *Class) UpsertName() string {
	return consts.UPSERT + c.Name()
}

func (c *Class) UpsertOneName() string {
	return consts.UPSERT_ONE + c.Name()
}

func (c *Class) AggregateName() string {
	return c.Name() + utils.FirstUpper(consts.AGGREGATE)
}
