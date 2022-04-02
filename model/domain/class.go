package domain

import "rxdrag.com/entity-engine/model/meta"

type Class struct {
	Uuid         string
	Type         string
	Name         string
	Description  string
	Associations map[string]*Association
	Attributes   []*Attribute
	Methods      []*Method
	Parents      []*Class
	Children     []*Class
}

func NewClass(c *meta.ClassMeta) *Class {
	cls := Class{
		Name:         c.Name,
		Description:  c.Description,
		Associations: map[string]*Association{},
		Attributes:   make([]*Attribute, len(c.Attributes)),
		Methods:      make([]*Method, len(c.Methods)),
		Parents:      []*Class{},
		Children:     []*Class{},
	}

	for i := range c.Attributes {
		cls.Attributes[i] = NewAttribute(&c.Attributes[i], &cls)
	}

	for i := range c.Methods {
		cls.Methods[i] = NewMethod(&c.Methods[i], &cls)
	}

	return &cls
}
