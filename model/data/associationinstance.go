package data

import "rxdrag.com/entity-engine/model/table"

type AssociationInstance struct {
	Fields      []*Field
	Association Associationer
}

type DerivedAssociationInstance struct {
	Fields           []*Field
	DerivedReference *DerivedReference
}

func NewAssociationInstance(association Associationer, sourceId uint64, targetId uint64) *AssociationInstance {
	sourceColumn := association.SourceColumn()
	targetColumn := association.TargetColumn()
	instance := AssociationInstance{
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

func (a *AssociationInstance) Table() *table.Table {
	return a.Association.Table()
}
