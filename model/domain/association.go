package domain

import "rxdrag.com/entity-engine/model/meta"

type Association struct {
	Name        string
	Relation    *meta.RelationMeta
	OfEntity    *Class
	TypeEntity  *Class
	Description string
}
