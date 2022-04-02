package graph

import (
	"rxdrag.com/entity-engine/model/domain"
)

type Class struct {
	Attributes   []*Attribute
	Associations []*Association
	Methods      []*Method
	Domain       *domain.Class
}
