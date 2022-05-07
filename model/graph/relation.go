package graph

import (
	"rxdrag.com/entify/model/domain"
	"rxdrag.com/entify/model/meta"
	"rxdrag.com/entify/model/table"
)

type Relation struct {
	Uuid                   string
	InnerId                uint64
	RelationType           string
	Source                 Noder
	Target                 Noder
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

func NewRelation(r *domain.Relation, s Noder, t Noder) *Relation {
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
	if r.Source.IsInterface() || r.Target.IsInterface() {
		return false
	}

	return true
}
