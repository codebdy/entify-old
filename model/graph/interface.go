package graph

import "rxdrag.com/entify/model/domain"

type Interface struct {
	Class
	Children []*Entity
	Parents  []*Interface
}

func NewInterface(c *domain.Class) *Interface {
	return &Interface{
		Class: *NewClass(c),
	}
}

func (f *Interface) IsInterface() bool {
	return true
}
func (f *Interface) Interface() *Interface {
	return f
}
func (f *Interface) Entity() *Entity {
	return nil
}

func (f *Interface) AllAttributes() []*Attribute {
	attrs := []*Attribute{}
	attrs = append(attrs, f.attributes...)
	for i := range f.Parents {
		for j := range f.Parents[i].attributes {
			attr := f.Parents[i].attributes[j]
			if findAttribute(attr.Name, attrs) == nil {
				attrs = append(attrs, attr)
			}
		}
	}
	return attrs
}

func (f *Interface) AllMethods() []*Method {
	methods := []*Method{}
	methods = append(methods, f.methods...)
	for i := range f.Parents {
		for j := range f.Parents[i].methods {
			method := f.Parents[i].methods[j]
			if findMethod(method.Name(), methods) == nil {
				methods = append(methods, method)
			}
		}
	}
	return methods
}

func (f *Interface) AllAssociations() []*Association {
	associas := []*Association{}
	associas = append(associas, f.associations...)
	for i := range f.Parents {
		for j := range f.Parents[i].associations {
			asso := f.Parents[i].associations[j]
			if findAssociation(asso.Name(), associas) == nil {
				associas = append(associas, asso)
			}
		}
	}
	return associas
}

func (f *Interface) GetAssociationByName(name string) *Association {
	associations := f.AllAssociations()
	for i := range associations {
		if associations[i].Name() == name {
			return associations[i]
		}
	}

	return nil
}

func (f *Interface) IsEmperty() bool {
	return len(f.AllAttributes()) < 1 && len(f.AllAssociations()) < 1
}

func (f *Interface) AllAttributeNames() []string {
	names := make([]string, len(f.AllAttributes()))

	for i, attr := range f.AllAttributes() {
		names[i] = attr.Name
	}

	return names
}

func (f *Interface) GetAttributeByName(name string) *Attribute {
	for _, attr := range f.AllAttributes() {
		if attr.Name == name {
			return attr
		}
	}

	return nil
}
