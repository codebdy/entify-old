package graph

import (
	"rxdrag.com/entity-engine/model/meta"
)

type Association struct {
	Relation       *Relation
	OwnerClassUuid string
}

type DerivedAssociation struct {
	Relation       *DerivedRelation
	DerivedFrom    *Association
	OwnerClassUuid string
}

func NewAssociation(r *Relation, ownerUuid string) *Association {
	return &Association{
		Relation:       r,
		OwnerClassUuid: ownerUuid,
	}
}

func (a *Association) Name() string {
	if a.IsSource() {
		return a.Relation.RoleOfTarget
	} else {
		return a.Relation.RoleOfSource
	}
}

func (a *Association) Owner() Node {
	if a.IsSource() {
		return a.Relation.Source
	} else {
		return a.Relation.Target
	}
}

func (a *Association) TypeClass() Node {
	if a.IsSource() {
		return a.Relation.Target
	} else {
		return a.Relation.Source
	}
}

func (a *Association) Description() string {
	if a.IsSource() {
		return a.Relation.DescriptionOnTarget
	} else {
		return a.Relation.DescriptionOnSource
	}
}

func (a *Association) IsArray() bool {
	if a.IsSource() {
		return a.Relation.TargetMultiplicity == meta.ZERO_MANY
	} else {
		return a.Relation.SourceMutiplicity == meta.ZERO_MANY
	}
}

func (a *Association) IsSource() bool {
	return a.Relation.Source.Uuid() == a.OwnerClassUuid
}

//对手实体类
func (a *DerivedAssociation) TargetEntities() []*Entity {
	//targetNode := a.Relation.Target
	return []*Entity{}
}

func (a *Association) GetName() string {
	return a.Name()
}
