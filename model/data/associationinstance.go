package data

import "rxdrag.com/entity-engine/model/table"

type AssociationInstance struct {
	Fields    []*Field
	Reference *Reference
}

type DerivedAssociationInstance struct {
	Fields           []*Field
	DerivedReference *DerivedReference
}

func NewAssociationInstance(ref *Reference, sourceId uint64, targetId uint64) *AssociationInstance {
	sourceColumn := ref.SourceColumn()
	targetColumn := ref.TargetColumn()
	instance := AssociationInstance{
		Reference: ref,
		Fields: []*Field{
			{
				Column: sourceColumn,
			},
			{
				Column: targetColumn,
			},
		},
	}

	return &instance
}

func (a *AssociationInstance) Table() *table.Table {
	return a.Reference.Association.Relation.Table
}
