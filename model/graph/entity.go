package graph

import (
	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/model/domain"
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
