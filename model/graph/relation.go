package graph

import (
	"rxdrag.com/entity-engine/model/domain"
	"rxdrag.com/entity-engine/model/meta"
	"rxdrag.com/entity-engine/model/table"
)

type Relation struct {
	Uuid                   string
	InnerId                uint64
	RelationType           string
	Source                 Node
	Target                 Node
	RoleOfTarget           string
	RoleOfSource           string
	DescriptionOnSource    string
	DescriptionOnTarget    string
	SourceMutiplicity      string
	TargetMultiplicity     string
	EnableAssociaitonClass bool
	AssociationClass       meta.AssociationClass
	Children               []*DerivedRelation
	Table                  *table.Table
}

type DerivedRelation struct {
	Parent *Relation
	Source *Entity
	Target *Entity
	Table  *table.Table
}

func NewRelation(r *domain.Relation, s Node, t Node) *Relation {
	return &Relation{
		Uuid:                   r.Uuid,
		InnerId:                r.InnerId,
		RelationType:           r.RelationType,
		Source:                 s,
		Target:                 t,
		RoleOfTarget:           r.RoleOfTarget,
		RoleOfSource:           r.RoleOfSource,
		DescriptionOnSource:    r.DescriptionOnSource,
		DescriptionOnTarget:    r.DescriptionOnTarget,
		SourceMutiplicity:      r.SourceMutiplicity,
		TargetMultiplicity:     r.TargetMultiplicity,
		EnableAssociaitonClass: r.EnableAssociaitonClass,
		AssociationClass:       r.AssociationClass,
	}
}

func (r *Relation) IsRealRelation() bool {
	if r.Source.isInterface() || r.Target.isInterface() {
		return false
	}

	return true
}
