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
		table := &Table{Name: entity.GetTableName()}
		for j := range entity.Columns {
			table.Columns = append(table.Columns, &entity.Columns[j])
		}

		tables[i] = table
	}

	return tables
}

func (c *MetaContent) Tables() []*Table {
	tables := c.entityTables()

	return tables
}
