package domain

import "rxdrag.com/entity-engine/oldmeta"

type Attribute struct {
	oldmeta.ColumnMeta
	Entity *Entity
}
