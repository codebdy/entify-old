package model

import (
	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/meta"
	"rxdrag.com/entity-engine/utils"
)

type Association struct {
	Name        string
	Relation    *Relation
	OfEntity    *Entity
	TypeEntity  *Entity
	Description string
}

type Entity struct {
	meta.EntityMeta
	Parent       *Entity
	Children     []*Entity
	Associations []*Association
	model        *Model
}

type Relation struct {
	meta.RelationMeta
	model *Model
}

type Model struct {
	Entities  []*Entity
	Relations []*Relation
	Tables    []*Table
}

func NewModel(c *meta.MetaContent) *Model {
	model := Model{
		Entities:  make([]*Entity, len(c.Entities)),
		Relations: []*Relation{},
		Tables:    entityTables(c),
	}

	model.buildEntities(c)
	model.buildRelations(c)
	model.buildInheritEntities()
	model.buildInheritRelations()
	model.buildTables()
	return &model
}

func (model *Model) buildTables() {
	for i := range model.Relations {
		relation := model.Relations[i]
		if relation.RelationType != meta.IMPLEMENTS {
			relationTable := relation.Table()
			model.Tables = append(model.Tables, relationTable)
		}
	}
}

//展开继承的关系
func (model *Model) buildInheritRelations() {
}

func (model *Model) buildInheritEntities() {
	for i := range model.Relations {
		relation := model.Relations[i]
		if relation.RelationType == meta.IMPLEMENTS {
			sourceEntity := model.GetEntityByUuid(relation.SourceId)
			if sourceEntity == nil {
				panic("Can not find entity by relation:" + relation.SourceId)
			}
			targetEntity := model.GetEntityByUuid(relation.TargetId)
			if targetEntity == nil {
				panic("Can not find entity by relation:" + relation.TargetId)
			}

			sourceEntity.Parent = targetEntity
			targetEntity.Children = append(targetEntity.Children, sourceEntity)
		}
	}
}

func (model *Model) buildRelations(c *meta.MetaContent) {
	for i := range c.Relations {
		relation := c.Relations[i]
		if relation.RelationType != meta.IMPLEMENTS {
			model.Relations = append(model.Relations, &Relation{
				RelationMeta: relation,
				model:        model,
			})
		}
	}
}

func (model *Model) buildEntities(c *meta.MetaContent) {
	for i := range c.Entities {
		model.Entities[i] = &Entity{
			EntityMeta:   c.Entities[i],
			Parent:       nil,
			Children:     []*Entity{},
			Associations: []*Association{},
			model:        model,
		}
	}
}

func (m *Model) GetEntityByUuid(uuid string) *Entity {
	for i := range m.Entities {
		entity := m.Entities[i]
		if entity.Uuid == uuid {
			return entity
		}
	}
	return nil
}

func FindTable(metaUuid string, tables []*Table) *Table {
	for i := range tables {
		if tables[i].MetaUuid == metaUuid {
			return tables[i]
		}
	}
	return nil
}

func entityTables(c *meta.MetaContent) []*Table {

	normalEntities := c.FilterEntity(func(e *meta.EntityMeta) bool {
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

func (relation *Relation) Table() *Table {
	table := &Table{
		MetaUuid: relation.Uuid,
		Name:     relation.TableName(),
		Columns: []meta.ColumnMeta{
			{
				Name:  relation.RelationSourceColumnName(),
				Type:  meta.COLUMN_ID,
				Uuid:  relation.Uuid + consts.SUFFIX_SOURCE,
				Index: true,
			},
			{
				Name:  relation.RelationTargetColumnName(),
				Type:  meta.COLUMN_ID,
				Uuid:  relation.Uuid + consts.SUFFIX_TARGET,
				Index: true,
			},
		},
	}
	table.Columns = append(table.Columns, relation.Columns...)

	return table
}

func (relation *Relation) TableName() string {
	return relation.SouceTableName() +
		"_" + utils.SnakeString(relation.RoleOnSource) +
		"_" + relation.TargetTableName() +
		"_" + utils.SnakeString(relation.RoleOnTarget) +
		consts.SUFFIX_PIVOT
}

func (relation *Relation) SouceTableName() string {
	sourceEntity := relation.model.GetEntityByUuid(relation.SourceId)
	return sourceEntity.GetTableName()
}

func (relation *Relation) TargetTableName() string {
	targetEntity := relation.model.GetEntityByUuid(relation.TargetId)
	return targetEntity.GetTableName()
}
