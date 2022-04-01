package domain

import (
	"rxdrag.com/entity-engine/oldmeta"
)

type Association struct {
	Name        string
	Relation    *Relation
	OfEntity    *Entity
	TypeEntity  *Entity
	Description string
}

type Entity struct {
	oldmeta.EntityMeta
	Associations map[string]*Association
	Columns      []*Attribute
	Parents      []*Entity
	Children     []*Entity
}
