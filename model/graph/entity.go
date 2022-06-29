package graph

import (
	"rxdrag.com/entify/consts"
	"rxdrag.com/entify/model/domain"
	"rxdrag.com/entify/model/table"
	"rxdrag.com/entify/utils"
)

type Entity struct {
	Class
	Table      *table.Table
	Interfaces []*Interface
}

func NewEntity(c *domain.Class) *Entity {
	return &Entity{
		Class: *NewClass(c),
	}
}

func (e *Entity) GetHasManyName() string {
	return utils.FirstUpper(consts.SET) + e.Name() + consts.HAS_MANY
}

func (e *Entity) GetHasOneName() string {
	return utils.FirstUpper(consts.SET) + e.Name() + consts.HAS_ONE
}

//有同名接口
func (e *Entity) hasInterfaceWithSameName() bool {
	return e.Domain.HasChildren()
}

//包含继承来的
func (e *Entity) AllAttributes() []*Attribute {
	attrs := []*Attribute{}
	attrs = append(attrs, e.attributes...)
	for i := range e.Interfaces {
		for j := range e.Interfaces[i].attributes {
			attr := e.Interfaces[i].attributes[j]
			if findAttribute(attr.Name, attrs) == nil {
				attrs = append(attrs, attr)
			}
		}
	}
	return attrs
}

func (e *Entity) AllMethods() []*Method {
	methods := []*Method{}
	methods = append(methods, e.methods...)
	for i := range e.Interfaces {
		for j := range e.Interfaces[i].methods {
			method := e.Interfaces[i].methods[j]
			if findMethod(method.GetName(), methods) == nil {
				methods = append(methods, method)
			}
		}
	}
	return methods
}

//包含继承来的
func (e *Entity) AllAssociations() []*Association {
	associas := []*Association{}
	associas = append(associas, e.associations...)
	for i := range e.Interfaces {
		for j := range e.Interfaces[i].associations {
			asso := e.Interfaces[i].associations[j]
			if findAssociation(asso.Name(), associas) == nil {
				associas = append(associas, asso)
			}
		}
	}
	return associas
}

func (e *Entity) GetAssociationByName(name string) *Association {
	associations := e.AllAssociations()
	for i := range associations {
		if associations[i].Name() == name {
			return associations[i]
		}
	}

	return nil
}

func (e *Entity) IsEmperty() bool {
	return len(e.AllAttributes()) < 1 && len(e.AllAssociations()) < 1
}

func (e *Entity) AllAttributeNames() []string {
	names := make([]string, len(e.AllAttributes()))

	for i, attr := range e.AllAttributes() {
		names[i] = attr.Name
	}

	return names
}

func (e *Entity) GetAttributeByName(name string) *Attribute {
	for _, attr := range e.AllAttributes() {
		if attr.Name == name {
			return attr
		}
	}

	return nil
}

func findAttribute(name string, attrs []*Attribute) *Attribute {
	for i := range attrs {
		if attrs[i].Name == name {
			return attrs[i]
		}
	}
	return nil
}

func findMethod(name string, methods []*Method) *Method {
	for i := range methods {
		if methods[i].GetName() == name {
			return methods[i]
		}
	}
	return nil
}

func findAssociation(name string, assos []*Association) *Association {
	for i := range assos {
		if assos[i].Name() == name {
			return assos[i]
		}
	}
	return nil
}
