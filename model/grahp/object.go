package grahp

import (
	"rxdrag.com/entity-engine/model/domain"
	"rxdrag.com/entity-engine/model/table"
)

type Object struct {
	Class      *domain.Class
	Table      *table.Table
	Interfaces []*Interface
}
