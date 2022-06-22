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
		relation.Source.InnerId(),
		relation.InnerId,
		relation.Target.InnerId(),
	)
	if relation.IsRealRelation() {
		tab := &table.Table{
			Uuid: relation.Source.Uuid() + relation.Uuid + relation.Target.Uuid(),
			Name: name,
			Columns: []*table.Column{
				{
					AttributeMeta: meta.AttributeMeta{
						Type:  meta.ID,
						Uuid:  relation.Source.Uuid() + relation.Uuid,
						Name:  relation.Source.TableName(),
						Index: true,
					},
				},
				{
					AttributeMeta: meta.AttributeMeta{
						Type:  meta.ID,
						Uuid:  relation.Target.Uuid() + relation.Uuid,
						Name:  relation.Target.TableName(),
						Index: true,
					},
				},
			},
			PKString: fmt.Sprintf("%s,%s", relation.Source.TableName(), relation.Target.TableName()),
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
		derived.Source.InnerId(),
		derived.Parent.InnerId,
		derived.Target.InnerId(),
	)
	tab := &table.Table{
		Uuid: derived.Source.Uuid() + derived.Parent.Uuid + derived.Target.Uuid(),
		Name: name,
		Columns: []*table.Column{
			{
				AttributeMeta: meta.AttributeMeta{
					Type:  meta.ID,
					Uuid:  derived.Source.Uuid() + derived.Parent.Uuid,
					Name:  derived.Source.TableName(),
					Index: true,
				},
			},
			{
				AttributeMeta: meta.AttributeMeta{
					Type:  meta.ID,
					Uuid:  derived.Target.Uuid() + derived.Parent.Uuid,
					Name:  derived.Target.TableName(),
					Index: true,
				},
			},
		},
		PKString: fmt.Sprintf("%s,%s", derived.Source.TableName(), derived.Target.TableName()),
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
