package graph

import (
	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/model/domain"
	"rxdrag.com/entity-engine/model/meta"
	"rxdrag.com/entity-engine/model/table"
	"rxdrag.com/entity-engine/utils"
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
	return utils.FirstUpper(consts.UPDATE) + e.Name() + consts.HAS_MANY
}

func (e *Entity) GetHasOneName() string {
	return utils.FirstUpper(consts.UPDATE) + e.Name() + consts.HAS_ONE
}

//有同名接口
func (e *Entity) hasInterfaceWithSameName() bool {
	return e.Domain.HasChildren()
}

func (e *Entity) isInterface() bool {
	return false
}
func (e *Entity) Interface() *Interface {
	return nil
}
func (e *Entity) Entity() *Entity {
	return e
}

//包含继承来的
func (e *Entity) AllAttributes() []*Attribute {
	attrs := []*Attribute{}
	attrs = append(attrs, e.attributes...)
	for i := range e.Interfaces {
		attrs = append(attrs, e.Interfaces[i].attributes...)
	}
	return attrs
}

//包含继承来的
func (e *Entity) AllAssociations() []*Association {
	associas := []*Association{}
	associas = append(associas, e.associations...)
	for i := range e.Interfaces {
		associas = append(associas, e.Interfaces[i].associations...)
	}
	return associas
}

func (c *Entity) IsEmperty() bool {
	return len(c.AllAttributes()) < 1 && len(c.AllAssociations()) < 1 && c.Domain.StereoType != meta.CLASS_SERVICE
}

func (c *Entity) AllAttributeNames() []string {
	names := make([]string, len(c.AllAttributes()))

	for i, attr := range c.AllAttributes() {
		names[i] = attr.Name
	}

	return names
}

func (c *Entity) GetAttributeByName(name string) *Attribute {
	for _, attr := range c.AllAttributes() {
		if attr.Name == name {
			return attr
		}
	}

	return nil
}
