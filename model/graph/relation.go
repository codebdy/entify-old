package graph

import "rxdrag.com/entity-engine/model/meta"

type Relation struct {
	Uuid                string
	InnerId             uint64
	RelationType        string
	Source              *Entity
	Target              *Entity
	RoleOfTarget        string
	RoleOfSource        string
	DescriptionOnSource string
	DescriptionOnTarget string
	SourceMutiplicity   string
	TargetMultiplicity  string
	AssociationClass    *meta.AssociationClass
}
