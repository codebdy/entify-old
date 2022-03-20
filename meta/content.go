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
	//检查空实体（除ID外没有属性跟关联）
	for _, entity := range c.Entities {
		if len(entity.Columns) < 1 && len(c.EntityRelations(&entity)) < 1 && entity.HasTable() {
			panic(fmt.Sprintf("Entity %s should have one normal field at least", entity.Name))
		}
	}

	//检查继承闭环
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
				column.Name = relation.RelationSourceColumnName()
			} else {
				column.Name = relation.RelationTargetColumnName()
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
				Name:  relation.RelationTargetColumnName(),
				Uuid:  relation.Uuid + consts.SUFFIX_TARGET,
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
				Name:  relation.RelationSourceColumnName(),
				Uuid:  relation.Uuid + consts.SUFFIX_SOURCE,
				Index: true,
			}
			ownerTable.Columns = append(ownerTable.Columns, column)
		} else if relation.RelationType == IMPLEMENTS {
			sourceTable := FindTable(relation.SourceId, tables)
			if sourceTable == nil {
				panic("Can not find parent table, relation:" + relation.Uuid)
			}

			column := Column{
				Type:  COLUMN_ID,
				Name:  consts.PARENT_ID,
				Index: true,
				Uuid:  relation.Uuid,
			}
			sourceTable.Columns = append(sourceTable.Columns, column)
		}
	}
	return tables
}

func (c *MetaContent) RelationTableName(relation *Relation) string {
	return c.RelationSouceTableName(relation) + "_" + c.RelationTargetTableName(relation) + consts.SUFFIX_PIVOT
}

func (c *MetaContent) RelationSouceTableName(relation *Relation) string {
	sourceEntity := c.GetEntityByUuid(relation.SourceId)
	return sourceEntity.GetTableName()
}

func (c *MetaContent) RelationTargetTableName(relation *Relation) string {
	targetEntity := c.GetEntityByUuid(relation.TargetId)
	return targetEntity.GetTableName()
}

func (c *MetaContent) entityTables() []*Table {

	normalEntities := c.filterEntity(func(e *Entity) bool {
		return e.HasTable()
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
				Name:  relation.RelationSourceColumnName(),
				Type:  COLUMN_ID,
				Uuid:  relation.Uuid + consts.SUFFIX_SOURCE,
				Index: true,
			},
			{
				Name:  relation.RelationTargetColumnName(),
				Type:  COLUMN_ID,
				Uuid:  relation.Uuid + consts.SUFFIX_TARGET,
				Index: true,
			},
		},
	}
	table.Columns = append(table.Columns, relation.Columns...)

	return table
}

func (c *MetaContent) Interfaces(entity *Entity) []*Entity {
	interfaces := []*Entity{}
	for i := range c.Relations {
		relation := &c.Relations[i]
		if relation.RelationType == IMPLEMENTS {
			if relation.SourceId == entity.Uuid {
				oneInterface := c.GetEntityByUuid(relation.TargetId)
				if oneInterface == nil {
					panic("Can not find interface:" + relation.TargetId)
				}
				interfaces = append(interfaces, oneInterface)
			}
		}
	}
	return interfaces
}

func (c *MetaContent) Children(entity *Entity) []*Entity {
	children := []*Entity{}
	for i := range c.Relations {
		relation := &c.Relations[i]
		if relation.RelationType == IMPLEMENTS {
			if relation.TargetId == entity.Uuid {
				child := c.GetEntityByUuid(relation.SourceId)
				if child == nil {
					panic("Cant find child:" + relation.SourceId)
				}
				children = append(children, child)
			}
		}
	}
	return children
}

func (c *MetaContent) HasChildren(entity *Entity) bool {
	children := c.Children(entity)
	return len(children) > 0
}

func (c *MetaContent) EntityRelations(entity *Entity) []EntityRelation {
	relations := []EntityRelation{}
	for i := range c.Relations {
		relation := &c.Relations[i]
		if relation.RelationType == IMPLEMENTS {
			continue
		}
		if relation.SourceId == entity.Uuid {
			relations = append(relations, EntityRelation{
				Name:       relation.RoleOnSource,
				Relation:   relation,
				OfEntity:   entity,
				TypeEntity: c.GetEntityByUuid(relation.TargetId),
			})
		} else if relation.TargetId == entity.Uuid {
			relations = append(relations, EntityRelation{
				Name:       relation.RoleOnTarget,
				Relation:   relation,
				OfEntity:   entity,
				TypeEntity: c.GetEntityByUuid(relation.SourceId),
			})
		}
	}
	return relations
}

func (c *MetaContent) EntityInheritedRelations(entity *Entity) []EntityRelation {
	relations := []EntityRelation{}
	parents := c.Interfaces(entity)
	for _, parent := range parents {
		relations = append(relations, c.EntityRelations(parent)...)
	}
	return relations
}

func (c *MetaContent) EntityAllRelations(entity *Entity) []EntityRelation {
	var inheritedRelations []EntityRelation
	var allInheritedRelations = c.EntityInheritedRelations(entity)
	entityRelations := c.EntityRelations(entity)
	for i := range allInheritedRelations {
		relation := allInheritedRelations[i]
		if FindRelationByName(relation.Name, entityRelations) == nil {
			inheritedRelations = append(inheritedRelations, relation)
		}
	}
	return append(entityRelations, inheritedRelations...)
}

func (c *MetaContent) EntityInheritedColumns(entity *Entity) []Column {
	columns := []Column{}
	parents := c.Interfaces(entity)
	for _, parent := range parents {
		columns = append(columns, parent.Columns...)
	}

	return columns
}

func (c *MetaContent) EntityAllColumns(entity *Entity) []Column {
	var inheritedColumns []Column
	var allInheritedColumns = c.EntityInheritedColumns(entity)
	for i := range allInheritedColumns {
		column := allInheritedColumns[i]
		if FindColumnByName(column.Name, entity.Columns) == nil {
			inheritedColumns = append(inheritedColumns, column)
		}
	}
	return append(entity.Columns, inheritedColumns...)
}

/**
* 把实体类分类
 */
func (c *MetaContent) SplitEntities() ([]*Entity, []*Entity, []*Entity) {
	var enumEntities []*Entity
	var interfaceEntities []*Entity
	var normalEntities []*Entity
	for i := range c.Entities {
		entity := &c.Entities[i]
		if entity.EntityType == ENTITY_ENUM {
			enumEntities = append(enumEntities, entity)
		} else if entity.EntityType == ENTITY_INTERFACE {
			interfaceEntities = append(interfaceEntities, entity)
		} else {
			normalEntities = append(normalEntities, entity)
		}
	}
	return enumEntities, interfaceEntities, normalEntities
}
