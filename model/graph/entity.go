package graph

import (
	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/model/table"
	"rxdrag.com/entity-engine/utils"
)

type Entity struct {
	Class
	Table      *table.Table
	Interfaces []*Interface
}

func (e *Entity) Name() string {
	return e.Domain.Name
}

func (e *Entity) Uuid() string {
	return e.Domain.Uuid
}

func (entity *Entity) GetHasManyName() string {
	return utils.FirstUpper(consts.UPDATE) + entity.Name() + consts.HAS_MANY
}

func (entity *Entity) GetHasOneName() string {
	return utils.FirstUpper(consts.UPDATE) + entity.Name() + consts.HAS_ONE
}
