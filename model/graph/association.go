package graph

import (
	"rxdrag.com/entify/model/meta"
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

func (a *Association) Owner() *Class {
	if a.IsSource() {
		return a.Relation.SourceClass()
	} else {
		return a.Relation.TargetClass()
	}
}

func (a *Association) TypeClass() *Class {
	if !a.IsSource() {
		return a.Relation.SourceClass()
	} else {
		return a.Relation.TargetClass()
	}
}

func (a *Association) TypeInterface() *Interface {
	if !a.IsSource() {
		return a.Relation.SourceInterface
	} else {
		return a.Relation.TargetInterface
	}
}

func (a *Association) TypeEntity() *Entity {
	if !a.IsSource() {
		return a.Relation.SourceEntity
	} else {
		return a.Relation.TargetEntity
	}
}

func (a *Association) TypePartial() *Partial {
	if !a.IsSource() {
		return a.Relation.SourcePartial
	} else {
		return a.Relation.TargetPartial
	}
}

func (a *Association) TypeExternal() *External {
	if !a.IsSource() {
		return a.Relation.SourceExternal
	} else {
		return a.Relation.TargetExternal
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
	return a.Relation.SourceClass().Uuid() == a.OwnerClassUuid
}

func (a *Association) IsAbstract() bool {
	return len(a.Relation.Children) > 0
}

func (a *Association) DerivedAssociations() []*DerivedAssociation {
	associations := []*DerivedAssociation{}
	for i := range a.Relation.Children {
		derivedRelation := a.Relation.Children[i]
		ownerUuid := derivedRelation.SourceClass().Uuid()
		if a.Relation.TargetClass().Uuid() == a.OwnerClassUuid {
			ownerUuid = derivedRelation.TargetClass().Uuid()
		}
		associations = append(associations, &DerivedAssociation{
			Relation:       derivedRelation,
			DerivedFrom:    a,
			OwnerClassUuid: ownerUuid,
		})
	}
	return associations
}

func (a *Association) DerivedAssociationsByOwnerUuid(ownerUuid string) []*DerivedAssociation {
	associations := []*DerivedAssociation{}
	allDerived := a.DerivedAssociations()
	for i := range allDerived {
		if allDerived[i].OwnerClassUuid == ownerUuid {
			associations = append(associations, allDerived[i])
		}
	}
	return associations
}

func (a *Association) GetName() string {
	return a.Name()
}

func (a *Association) Path() string {
	return a.Owner().Domain.Name + "." + a.Name()
}

//对手实体类
func (d *DerivedAssociation) TypeClass() *Class {
	if d.Relation.SourceClass().Uuid() == d.OwnerClassUuid {
		return d.Relation.TargetClass()
	} else {
		return d.Relation.SourceClass()
	}

}

func (d *DerivedAssociation) Owner() *Class {
	if d.Relation.SourceClass().Uuid() == d.OwnerClassUuid {
		return d.Relation.SourceClass()
	} else {
		return d.Relation.TargetClass()
	}

}

func (d *DerivedAssociation) Name() string {
	if d.TypeClass().Uuid() == d.DerivedFrom.TypeClass().Uuid() {
		return d.DerivedFrom.Name()
	} else {
		return d.DerivedFrom.Name() + "For" + d.TypeClass().Name()
	}
}

func (a *DerivedAssociation) TypeEntity() *Entity {
	if !a.DerivedFrom.IsSource() {
		return a.Relation.SourceEntity
	} else {
		return a.Relation.TargetEntity
	}
}

func (a *DerivedAssociation) TypePartial() *Partial {
	if !a.DerivedFrom.IsSource() {
		return a.Relation.SourcePartial
	} else {
		return a.Relation.TargetPartial
	}
}

func (a *DerivedAssociation) TypeExternal() *External {
	if !a.DerivedFrom.IsSource() {
		return a.Relation.SourceExternal
	} else {
		return a.Relation.TargetExternal
	}
}
