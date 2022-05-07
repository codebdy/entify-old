package domain

import "rxdrag.com/entify/model/meta"

type Enum struct {
	Uuid   string
	Name   string
	Values []string
}

func NewEnum(c *meta.ClassMeta) *Enum {
	enum := Enum{
		Uuid:   c.Uuid,
		Name:   c.Name,
		Values: make([]string, len(c.Attributes)),
	}

	for i := range c.Attributes {
		enum.Values[i] = c.Attributes[i].Name
	}

	return &enum
}
