package graph

import (
	"fmt"

	"rxdrag.com/entify/consts"
	"rxdrag.com/entify/model/meta"
	"rxdrag.com/entify/model/table"
)

func NewEntityTable(entity *Entity, partial bool) *table.Table {
	table := &table.Table{
		Uuid:          entity.Uuid(),
		Name:          entity.TableName(),
		EntityInnerId: entity.Domain.InnerId,
		Partial:       false,
	}

	allAttrs := entity.AllAttributes()
	for i := range allAttrs {
		attr := allAttrs[i]
		table.Columns = append(table.Columns, NewAttributeColumn(attr, partial))
	}

	entity.Table = table
	return table
}

func NewAttributeColumn(attr *Attribute, partial bool) *table.Column {
	return &table.Column{
		AttributeMeta: attr.AttributeMeta,
		PartialId:     partial && attr.Name == consts.ID,
	}
}

func NewRelationTables(relation *Relation) []*table.Table {
	var tables []*table.Table
	name := fmt.Sprintf(
		"%s_%d_%d_%d",
		consts.PIVOT,
		relation.SourceClass().InnerId(),
		relation.InnerId,
		relation.TargetClass().InnerId(),
	)
	if relation.IsRealRelation() {
		tab := &table.Table{
			Uuid: relation.SourceClass().Uuid() + relation.Uuid + relation.TargetClass().Uuid(),
			Name: name,
			Columns: []*table.Column{
				{
					AttributeMeta: meta.AttributeMeta{
						Type:  meta.ID,
						Uuid:  relation.SourceClass().Uuid() + relation.Uuid,
						Name:  relation.SourceClass().TableName(),
						Index: true,
					},
				},
				{
					AttributeMeta: meta.AttributeMeta{
						Type:  meta.ID,
						Uuid:  relation.TargetClass().Uuid() + relation.Uuid,
						Name:  relation.TargetClass().TableName(),
						Index: true,
					},
				},
			},
			PKString: fmt.Sprintf("%s,%s", relation.SourceClass().TableName(), relation.TargetClass().TableName()),
		}
		if relation.EnableAssociaitonClass {
			for i := range relation.AssociationClass.Attributes {
				tab.Columns = append(tab.Columns, &table.Column{
					AttributeMeta: relation.AssociationClass.Attributes[i],
				})
			}
		}
		relation.Table = tab
		tables = append(tables, tab)
	} else {
		for i := range relation.Children {
			derivied := relation.Children[i]
			tables = append(tables, NewDerivedRelationTable(derivied))
		}
	}

	return tables
}

func NewDerivedRelationTable(derived *DerivedRelation) *table.Table {
	name := fmt.Sprintf(
		"%s_%d_%d_%d",
		consts.PIVOT,
		derived.SourceClass().InnerId(),
		derived.Parent.InnerId,
		derived.TargetClass().InnerId(),
	)
	tab := &table.Table{
		Uuid: derived.SourceClass().Uuid() + derived.Parent.Uuid + derived.TargetClass().Uuid(),
		Name: name,
		Columns: []*table.Column{
			{
				AttributeMeta: meta.AttributeMeta{
					Type:  meta.ID,
					Uuid:  derived.SourceClass().Uuid() + derived.Parent.Uuid,
					Name:  derived.SourceClass().TableName(),
					Index: true,
				},
			},
			{
				AttributeMeta: meta.AttributeMeta{
					Type:  meta.ID,
					Uuid:  derived.TargetClass().Uuid() + derived.Parent.Uuid,
					Name:  derived.TargetClass().TableName(),
					Index: true,
				},
			},
		},
		PKString: fmt.Sprintf("%s,%s", derived.SourceClass().TableName(), derived.TargetClass().TableName()),
	}
	if derived.Parent.EnableAssociaitonClass {
		for i := range derived.Parent.AssociationClass.Attributes {
			tab.Columns = append(tab.Columns, &table.Column{
				AttributeMeta: derived.Parent.AssociationClass.Attributes[i],
			})
		}
	}
	derived.Table = tab
	return tab
}
