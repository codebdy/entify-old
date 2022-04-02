package graph

import (
	"rxdrag.com/entity-engine/model/table"
)

type Entity struct {
	Class
	Table      *table.Table
	Interfaces []*Interface
}
