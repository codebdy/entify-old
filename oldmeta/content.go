package oldmeta

type MetaContent struct {
	Entities  []EntityMeta   `json:"entities"`
	Relations []RelationMeta `json:"relations"`
	Diagrams  []interface{}  `json:"diagrams"`
	X6Nodes   []interface{}  `json:"x6Nodes"`
	X6Edges   []interface{}  `json:"x6Edges"`
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

/**
* 把关系分类
 */
func (c *MetaContent) SplitRelations() ([]*RelationMeta, []*RelationMeta) {
	var inherits []*RelationMeta
	var relations []*RelationMeta

	for i := range c.Relations {
		relation := &c.Relations[i]
		if relation.RelationType == IMPLEMENTS {
			inherits = append(inherits, relation)
		} else {
			relations = append(relations, relation)
		}
	}

	return inherits, relations
}
