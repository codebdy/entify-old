package domain

import "rxdrag.com/entity-engine/model/meta"

type Relation struct {
	Uuid                   string
	InnerId                uint64
	RelationType           string
	Source                 *Class
	Target                 *Class
	RoleOfTarget           string
	RoleOfSource           string
	DescriptionOnSource    string
	DescriptionOnTarget    string
	SourceMutiplicity      string
	TargetMultiplicity     string
	EnableAssociaitonClass bool
	AssociationClass       meta.AssociationClass
}

func NewRelation(r *meta.RelationMeta, s *Class, t *Class) *Relation {
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
