package meta

import (
	"fmt"

	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/utils"
)

type MetaContent struct {
	Entities  []EntityMeta   `json:"entities"`
	Relations []RelationMeta `json:"relations"`
	Diagrams  []interface{}  `json:"diagrams"`
	X6Nodes   []interface{}  `json:"x6Nodes"`
	X6Edges   []interface{}  `json:"x6Edges"`
}

func (c *MetaContent) Validate() {
	//检查空实体（除ID外没有属性跟关联）
	for _, entity := range c.Entities {
		if len(entity.Columns) < 1 && len(c.EntityRelations(&entity)) < 1 && entity.HasTable() {
			panic(fmt.Sprintf("Entity %s should have one normal field at least", entity.Name))
		}
	}
}

func (c *MetaContent) FilterEntity(equal func(entity *EntityMeta) bool) []*EntityMeta {
	entities := []*EntityMeta{}
	for i := range c.Entities {
		entity := &c.Entities[i]
		if equal(entity) {
			entities = append(entities, entity)
		}
	}
	return entities
}

func (c *MetaContent) GetEntityByUuid(uuid string) *EntityMeta {
	for i := range c.Entities {
		entity := &c.Entities[i]
		if entity.Uuid == uuid {
			return entity
		}
	}
	return nil
}

func (c *MetaContent) GetEntityByName(expName string) *EntityMeta {
	for i := range c.Entities {
		entity := &c.Entities[i]
		if entity.Name == expName {
			return entity
		}
	}
	return nil
}

func (c *MetaContent) RelationTableName(relation *RelationMeta) string {
	return c.RelationSouceTableName(relation) +
		"_" + utils.SnakeString(relation.RoleOnSource) +
		"_" + c.RelationTargetTableName(relation) +
		"_" + utils.SnakeString(relation.RoleOnTarget) +
		consts.SUFFIX_PIVOT
}

func (c *MetaContent) RelationSouceTableName(relation *RelationMeta) string {
	sourceEntity := c.GetEntityByUuid(relation.SourceId)
	return sourceEntity.GetTableName()
}

func (c *MetaContent) RelationTargetTableName(relation *RelationMeta) string {
	targetEntity := c.GetEntityByUuid(relation.TargetId)
	return targetEntity.GetTableName()
}

func (c *MetaContent) Interfaces(entity *EntityMeta) []*EntityMeta {
	interfaces := []*EntityMeta{}
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

func (c *MetaContent) Children(entity *EntityMeta) []*EntityMeta {
	children := []*EntityMeta{}
	for i := range c.Relations {
		relation := &c.Relations[i]
		if relation.RelationType == IMPLEMENTS {
			if relation.TargetId == entity.Uuid {
				child := c.GetEntityByUuid(relation.SourceId)
				if child == nil {
					panic("Can't find child:" + relation.SourceId)
				}
				children = append(children, child)
			}
		}
	}
	return children
}

func (c *MetaContent) HasChildren(entity *EntityMeta) bool {
	children := c.Children(entity)
	return len(children) > 0
}

func (c *MetaContent) EntityRelations(entity *EntityMeta) []EntityRelation {
	relations := []EntityRelation{}
	for i := range c.Relations {
		relation := &c.Relations[i]
		if relation.RelationType == IMPLEMENTS {
			continue
		}
		if relation.SourceId == entity.Uuid {
			relations = append(relations, EntityRelation{
				Name:        relation.RoleOnSource,
				Relation:    relation,
				OfEntity:    entity,
				TypeEntity:  c.GetEntityByUuid(relation.TargetId),
				Description: relation.DescriptionOnSource,
			})
		} else if relation.TargetId == entity.Uuid {
			relations = append(relations, EntityRelation{
				Name:        relation.RoleOnTarget,
				Relation:    relation,
				OfEntity:    entity,
				TypeEntity:  c.GetEntityByUuid(relation.SourceId),
				Description: relation.DescriptionOnTarget,
			})
		}
	}
	return relations
}

func (c *MetaContent) EntityInheritedRelations(entity *EntityMeta) []EntityRelation {
	relations := []EntityRelation{}
	parents := c.Interfaces(entity)
	for _, parent := range parents {
		relations = append(relations, c.EntityRelations(parent)...)
	}
	return relations
}

func (c *MetaContent) EntityAllRelations(entity *EntityMeta) []EntityRelation {
	var inheritedRelations []EntityRelation
	var allInheritedRelations = c.EntityInheritedRelations(entity)
	entityRelations := c.EntityRelations(entity)
	for i := range allInheritedRelations {
		relation := allInheritedRelations[i]
		if findRelationByName(relation.Name, entityRelations) == nil {
			inheritedRelations = append(inheritedRelations, relation)
		}
	}
	return append(entityRelations, inheritedRelations...)
}

func (c *MetaContent) EntityInheritedColumns(entity *EntityMeta) []ColumnMeta {
	columns := []ColumnMeta{}
	parents := c.Interfaces(entity)
	for _, parent := range parents {
		columns = append(columns, parent.Columns...)
	}

	return columns
}

func (c *MetaContent) EntityAllColumns(entity *EntityMeta) []ColumnMeta {
	var inheritedColumns []ColumnMeta
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
func (c *MetaContent) SplitEntities() ([]*EntityMeta, []*EntityMeta, []*EntityMeta) {
	var enumEntities []*EntityMeta
	var interfaceEntities []*EntityMeta
	var normalEntities []*EntityMeta
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

func findRelationByName(name string, relations []EntityRelation) *EntityRelation {
	for i := range relations {
		if relations[i].Name == name {
			return &relations[i]
		}
	}
	return nil
}
