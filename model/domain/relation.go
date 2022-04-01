package domain

import (
	"rxdrag.com/entity-engine/oldmeta"
)

type Relation struct {
	oldmeta.RelationMeta
}

type InheritedRelation struct {
	Relation
	InheritFrom *Relation
}
