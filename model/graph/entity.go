package graph

import (
	"rxdrag.com/entity-engine/model/table"
)

type Entity struct {
	Class
	Table      *table.Table
	Interfaces []*Interface
}

func (e *Entity) Name() string {
	return e.Domain.Name
}
