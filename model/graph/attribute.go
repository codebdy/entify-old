package graph

import "rxdrag.com/entity-engine/model/domain"

type Attribute struct {
	domain.Attribute
	Entity    *Entity
	EumnType  *Enum
	EnityType *Entity
}
