package graph

import (
	"rxdrag.com/entity-engine/model/domain"
)

type Enum struct {
	domain.Enum
}

func NewEnum(e *domain.Enum) *Enum {
	return &Enum{Enum: *e}
}
