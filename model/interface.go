package model

import "rxdrag.com/entity-engine/meta"

type Interface struct {
	meta.EntityMeta
	Associations []*Association
	Children     []*Entity
	model        *Model
}
