package grahp

import "rxdrag.com/entity-engine/model/domain"

type Interface struct {
	Class    *domain.Class
	Children []*Object
}
