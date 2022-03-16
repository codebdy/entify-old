package meta

import (
	"fmt"

	"rxdrag.com/entity-engine/consts"
)

type MetaContent struct {
	Entities  []Entity      `json:"entities"`
	Relations []Relation    `json:"relations"`
	Diagrams  []interface{} `json:"diagrams"`
	X6Nodes   []interface{} `json:"x6Nodes"`
	X6Edges   []interface{} `json:"x6Edges"`
}

func (c *MetaContent) Validate() {
	for _, entity := range c.Entities {
		if len(entity.Columns) <= 1 && entity.IsNormal() {
			panic(fmt.Sprintf("Entity %s should have one normal field at least", entity.Name))
		}
	}
}

func FindTable(metaUuid string, tables []*Table) *Table {
	for i := range tables {
		if tables[i].MetaUuid == metaUuid {
			return tables[i]
		}
	}
	return nil
}

func (c *MetaContent) filterEntity(equal func(entity *Entity) bool) []*Entity {
	entities := []*Entity{}
	for i := range c.Entities {
		entity := &c.Entities[i]
		if equal(entity) {
			entities = append(entities, entity)
		}
	}
	return entities
}

func (c *MetaContent) GetEntityByUuid(uuid string) *Entity {
	for i := range c.Entities {
		entity := &c.Entities[i]
		if entity.Uuid == uuid {
			return entity
		}
	}
	return nil
}

func (c *MetaContent) Tables() []*Table {
	tables := c.entityTables()
	for i := range c.Relations {
		relation := c.Relations[i]
		if relation.RelationType == MANY_TO_MANY {
			relationTable := c.relationTable(&relation)
			tables = append(tables, relationTable)
		} else if relation.RelationType == ONE_TO_ONE {
			ownerId := relation.OwnerId
			if ownerId == "" {
				ownerId = relation.SourceId
			}

			ownerTable := FindTable(ownerId, tables)
			if ownerTable == nil {
				panic("Can not find relation owner table, relation:" + relation.RoleOnSource + "-" + relation.RoleOnTarget)
			}

			column := Column{
				Type:  COLUMN_ID,
				Index: true,
			}
			if ownerId == relation.SourceId {
				column.Name = c.RelationTargetColumnName(&relation)
			} else {
				column.Name = c.RelationSourceColumnName(&relation)
			}
			ownerTable.Columns = append(ownerTable.Columns, column)

		} else if relation.RelationType == ONE_TO_MANY {
			ownerId := relation.TargetId
			ownerTable := FindTable(ownerId, tables)
			if ownerTable == nil {
				panic("Can not find relation owner table, relation:" + relation.RoleOnSource + "-" + relation.RoleOnTarget)
			}

			column := Column{
				Type:  COLUMN_ID,
				Name:  c.RelationTargetColumnName(&relation),
				Index: true,
			}
			ownerTable.Columns = append(ownerTable.Columns, column)

		} else if relation.RelationType == MANY_TO_ONE {
			ownerId := relation.SourceId
			ownerTable := FindTable(ownerId, tables)
			if ownerTable == nil {
				panic("Can not find relation owner table, relation:" + relation.RoleOnSource + "-" + relation.RoleOnTarget)
			}

			column := Column{
				Type:  COLUMN_ID,
				Name:  c.RelationTargetColumnName(&relation),
				Index: true,
			}
			ownerTable.Columns = append(ownerTable.Columns, column)
		} else if relation.RelationType == INHERIT {
			sourceTable := FindTable(relation.SourceId, tables)
			if sourceTable == nil {
				panic("Can not find parent table, relation:" + relation.Uuid)
			}

			column := Column{
				Type:  COLUMN_ID,
				Name:  consts.PARENT_ID,
				Index: true,
			}
			sourceTable.Columns = append(sourceTable.Columns, column)
		}
	}
	return tables
}

func (c *MetaContent) RelationTableName(relation *Relation) string {
	return c.RelationSouceTableName(relation) + "_" + c.RelationTargetTableName(relation) + consts.PIVOT_SUFFIX
}

func (c *MetaContent) RelationSouceTableName(relation *Relation) string {
	sourceEntity := c.GetEntityByUuid(relation.SourceId)
	return sourceEntity.GetTableName()
}

func (c *MetaContent) RelationTargetTableName(relation *Relation) string {
	targetEntity := c.GetEntityByUuid(relation.TargetId)
	return targetEntity.GetTableName()
}

func (c *MetaContent) RelationSourceColumnName(relation *Relation) string {
	return c.RelationSouceTableName(relation) + consts.ID_SUFFIX
}

func (c *MetaContent) RelationTargetColumnName(relation *Relation) string {
	return c.RelationTargetTableName(relation) + consts.ID_SUFFIX
}

func (c *MetaContent) entityTables() []*Table {

	normalEntities := c.filterEntity(func(e *Entity) bool {
		return e.IsNormal()
	})

	tables := make([]*Table, len(normalEntities))

	for i := range normalEntities {
		entity := normalEntities[i]
		table := &Table{Name: entity.GetTableName(), MetaUuid: entity.Uuid}
		table.Columns = append(table.Columns, entity.Columns...)
		tables[i] = table
	}

	return tables
}

func (c *MetaContent) relationTable(relation *Relation) *Table {
	table := &Table{
		MetaUuid: relation.Uuid,
		Name:     c.RelationTableName(relation),
		Columns: []Column{
			{
				Name:  c.RelationSourceColumnName(relation),
				Type:  COLUMN_ID,
				Index: true,
			},
			{
				Name:  c.RelationTargetColumnName(relation),
				Type:  COLUMN_ID,
				Index: true,
			},
		},
	}
	table.Columns = append(table.Columns, relation.Columns...)

	return table
}
