package model

import "rxdrag.com/entity-engine/meta"

type Interface struct {
	meta.EntityMeta
	Columns      []*Column
	Associations map[string]*Association
	Children     []*Entity
	model        *Model
}
