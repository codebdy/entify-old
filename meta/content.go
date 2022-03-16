package meta

type MetaContent struct {
	Entities  []Entity      `json:"entities"`
	Relations []Relation    `json:"relations"`
	Diagrams  []interface{} `json:"diagrams"`
	X6Nodes   []interface{} `json:"x6Nodes"`
	X6Edges   []interface{} `json:"x6Edges"`
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

func (c *MetaContent) EntityRelations(entityUuid string) []EntityRelation {
	return []EntityRelation{}
}

func (c *MetaContent) entityTables() []*Table {

	normalEntities := c.filterEntity(func(e *Entity) bool {
		return e.EntityType == Entity_NORMAL
	})

	tables := make([]*Table, len(normalEntities))

	for i := range normalEntities {
		entity := normalEntities[i]
		table := &Table{Name: entity.GetTableName(), MetaUuid: entity.Uuid}
		for j := range entity.Columns {
			table.Columns = append(table.Columns, &entity.Columns[j])
		}

		tables[i] = table
	}

	return tables
}

func (c *MetaContent) Tables() []*Table {
	tables := c.entityTables()

	for i := range c.Relations {
		relation := c.Relations[i]
		if relation.RelationType == MANY_TO_MANY {
			sourceEntity := c.GetEntityByUuid(relation.SourceId)
			targetEntity := c.GetEntityByUuid(relation.TargetId)
			tables = append(tables, &Table{
				MetaUuid: relation.Uuid,
				Name:     sourceEntity.GetTableName() + "_" + targetEntity.GetTableName() + "__middle",
			})
		}

	}
	return tables
}
