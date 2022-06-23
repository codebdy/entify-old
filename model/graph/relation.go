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
	SourceInterface        *Interface
	TargetInterface        *Interface
	SourceEntity           *Entity
	TargetEntity           *Entity
	SourcePartial          *Partial
	TargetPartial          *Partial
	SourceExternal         *External
	TargetExternal         *External
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
	Parent         *Relation
	SourceEntity   *Entity
	TargetEntity   *Entity
	SourcePartial  *Partial
	TargetPartial  *Partial
	SourceExternal *External
	TargetExternal *External
	Table          *table.Table
}

func NewRelation(
	r *domain.Relation,
	sourceInterface *Interface,
	targetInterface *Interface,
	sourceEntity *Entity,
	targetEntity *Entity,
	sourcePartial *Partial,
	targetPartial *Partial,
	sourceExternal *External,
	targetExternal *External,
) *Relation {
	relation := &Relation{
		Uuid:                   r.Uuid,
		InnerId:                r.InnerId,
		RelationType:           r.RelationType,
		SourceInterface:        sourceInterface,
		TargetInterface:        targetInterface,
		SourceEntity:           sourceEntity,
		TargetEntity:           targetEntity,
		SourcePartial:          sourcePartial,
		TargetPartial:          targetPartial,
		SourceExternal:         sourceExternal,
		TargetExternal:         targetExternal,
		RoleOfTarget:           r.RoleOfTarget,
		RoleOfSource:           r.RoleOfSource,
		DescriptionOnSource:    r.DescriptionOnSource,
		DescriptionOnTarget:    r.DescriptionOnTarget,
		SourceMutiplicity:      r.SourceMutiplicity,
		TargetMultiplicity:     r.TargetMultiplicity,
		EnableAssociaitonClass: r.EnableAssociaitonClass,
		AssociationClass:       r.AssociationClass,
	}

	return relation
}

func (r *Relation) SourceClass() *Class {
	if r.SourceInterface != nil {
		return &r.SourceInterface.Class
	}

	if r.SourceEntity != nil {
		return &r.SourceEntity.Class
	}

	if r.SourcePartial != nil {
		return &r.SourcePartial.Class
	}
	if r.SourceExternal != nil {
		return &r.SourceExternal.Class
	}
	return nil
}

func (r *Relation) TargetClass() *Class {
	if r.TargetInterface != nil {
		return &r.TargetInterface.Class
	}

	if r.TargetEntity != nil {
		return &r.TargetEntity.Class
	}

	if r.TargetPartial != nil {
		return &r.TargetPartial.Class
	}
	if r.TargetExternal != nil {
		return &r.TargetExternal.Class
	}
	return nil
}

func (r *Relation) IsRealRelation() bool {
	if r.SourceInterface != nil || r.TargetInterface != nil {
		return false
	}

	return true
}
