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
		table.Columns = append(table.Columns, entity.Columns...)
		tables[i] = table
	}

	return tables
}

func (c *MetaContent) Tables() []*Table {
	tables := c.entityTables()

	for i := range c.Relations {
		relation := c.Relations[i]
		if relation.RelationType == MANY_TO_MANY {
			relationTable := c.RelationTable(&relation)
			tables = append(tables, relationTable)
		}

	}
	return tables
}

func (c *MetaContent) RelationTable(relation *Relation) *Table {
	table := &Table{
		MetaUuid: relation.Uuid,
		Name:     c.RelationTableName(relation),
		Columns: []Column{
			{
				Name: c.RelationSourceColumnName(relation),
				Type: COLUMN_ID,
			},
			{
				Name: c.RelationTargetColumnName(relation),
				Type: COLUMN_ID,
			},
		},
	}
	table.Columns = append(table.Columns, relation.columns...)

	return table
}

func (c *MetaContent) RelationTableName(relation *Relation) string {
	return c.RelationSouceTableName(relation) + "_" + c.RelationTargetTableName(relation) + "__pivot"
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
	return c.RelationSouceTableName(relation) + "_id"
}

func (c *MetaContent) RelationTargetColumnName(relation *Relation) string {
	return c.RelationTargetTableName(relation) + "_id"
}
