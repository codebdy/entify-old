package domain

import "rxdrag.com/entity-engine/model/meta"

type Class struct {
	Uuid        string
	StereoType  string
	Name        string
	Description string
	Attributes  []*Attribute
	Methods     []*Method
	Parents     []*Class
	Children    []*Class
}

func NewClass(c *meta.ClassMeta) *Class {
	cls := Class{
		Uuid:        c.Uuid,
		StereoType:  c.StereoType,
		Name:        c.Name,
		Description: c.Description,
		Attributes:  make([]*Attribute, len(c.Attributes)),
		Methods:     make([]*Method, len(c.Methods)),
		Parents:     []*Class{},
		Children:    []*Class{},
	}

	for i := range c.Attributes {
		cls.Attributes[i] = NewAttribute(&c.Attributes[i], &cls)
	}

	for i := range c.Methods {
		cls.Methods[i] = NewMethod(&c.Methods[i], &cls)
	}

	return &cls
}

func (c *Class) HasChildren() bool {
	return len(c.Children) > 0
}
