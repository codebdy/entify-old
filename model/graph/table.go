package graph

import (
	"fmt"

	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/model/meta"
	"rxdrag.com/entity-engine/model/table"
)

func NewEntityTable(entity *Entity) *table.Table {
	table := &table.Table{
		Uuid:          entity.Uuid(),
		Name:          entity.TableName(),
		EntityInnerId: entity.Domain.InnerId,
	}

	allAttrs := entity.AllAttributes()
	for i := range allAttrs {
		attr := allAttrs[i]
		table.Columns = append(table.Columns, NewAttributeColumn(attr))
	}

	entity.Table = table
	return table
}

func NewAttributeColumn(attr *Attribute) *table.Column {
	return &table.Column{
		AttributeMeta: attr.AttributeMeta,
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
		table := &table.Table{
			Uuid: name,
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
		}
		relation.Table = table
		tables = append(tables, table)
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
	table := &table.Table{
		Uuid: name,
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
	}
	derived.Table = table
	return table
}
