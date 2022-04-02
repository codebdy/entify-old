package domain

import "rxdrag.com/entity-engine/model/meta"

type Enum struct {
	Name   string
	Values []string
}

func NewEnum(c *meta.ClassMeta) *Enum {
	enum := Enum{
		Name:   c.Name,
		Values: make([]string, len(c.Attributes)),
	}

	for i := range c.Attributes {
		enum.Values[i] = c.Attributes[i].Name
	}

	return &enum
}
