package data

import "rxdrag.com/entity-engine/model/table"

type AssociationPovit struct {
	Fields      []*Field
	Association Associationer
}

type DerivedAssociationPovit struct {
	Fields           []*Field
	DerivedReference *DerivedReference
}

func NewAssociationPovit(association Associationer, sourceId uint64, targetId uint64) *AssociationPovit {
	sourceColumn := association.SourceColumn()
	targetColumn := association.TargetColumn()
	instance := AssociationPovit{
		Association: association,
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

func (a *AssociationPovit) Table() *table.Table {
	return a.Association.Table()
}
