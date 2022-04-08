package graph

import "rxdrag.com/entity-engine/model/domain"

type Interface struct {
	Class
	Children []*Entity
}

func NewInterface(c *domain.Class) *Interface {
	return &Interface{
		Class: *NewClass(c),
	}
}

func (i *Interface) isInterface() bool {
	return true
}
func (i *Interface) Interface() *Interface {
	return i
}
func (i *Interface) Entity() *Entity {
	return nil
}

func (i *Interface) AllAttributes() []*Attribute {
	return i.attributes
}

func (i *Interface) AllMethods() []*Method {
	return i.methods
}

func (i *Interface) QueryAssociations() []*Association {
	return i.associations
}

func (c *Interface) IsEmperty() bool {
	return len(c.AllAttributes()) < 1 && len(c.QueryAssociations()) < 1
}

func (c *Interface) AllAttributeNames() []string {
	names := make([]string, len(c.AllAttributes()))

	for i, attr := range c.AllAttributes() {
		names[i] = attr.Name
	}

	return names
}

func (i *Interface) GetAttributeByName(name string) *Attribute {
	for _, attr := range i.AllAttributes() {
		if attr.Name == name {
			return attr
		}
	}

	return nil
}
