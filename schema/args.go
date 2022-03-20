package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/meta"
)

func (c *TypeCache) makeArgs() {
	for i := range meta.Metas.Entities {
		entity := &meta.Metas.Entities[i]
		if entity.EntityType != meta.ENTITY_ENUM {
			whereExp := makeWhereExp(entity)
			c.WhereExpMap[entity.Name] = whereExp
			orderByExp := makeOrderBy(entity)
			c.OrderByMap[entity.Name] = orderByExp
			distinctOnEnum := makeDistinctOnEnum(entity)
			c.DistinctOnEnumMap[entity.Name] = distinctOnEnum
		}
	}
	c.makeRelaionWhereExp()
}

func (c *TypeCache) makeRelaionWhereExp() {
	for entityName := range c.WhereExpMap {
		exp := c.WhereExpMap[entityName]
		entity := meta.Metas.GetEntityByName(entityName)
		if entity == nil {
			panic("Fatal error, can not find entity by name:" + entityName)
		}
		relations := meta.Metas.EntityAllRelations(entity)
		for i := range relations {
			relation := relations[i]
			exp.AddFieldConfig(relation.Name, &graphql.InputObjectFieldConfig{
				Type: c.WhereExp(relation.TypeEntity),
			})
		}
	}
}

func makeWhereExp(entity *meta.Entity) *graphql.InputObject {
	expName := entity.Name + consts.BOOLEXP
	andExp := graphql.InputObjectFieldConfig{}
	notExp := graphql.InputObjectFieldConfig{}
	orExp := graphql.InputObjectFieldConfig{}

	fields := graphql.InputObjectConfigFieldMap{
		consts.ARG_AND: &andExp,
		consts.ARG_NOT: &notExp,
		consts.ARG_OR:  &orExp,
	}

	boolExp := graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name:   expName,
			Fields: fields,
		},
	)
	andExp.Type = &graphql.List{
		OfType: &graphql.NonNull{
			OfType: boolExp,
		},
	}
	notExp.Type = boolExp
	orExp.Type = &graphql.List{
		OfType: &graphql.NonNull{
			OfType: boolExp,
		},
	}

	columns := meta.Metas.EntityAllColumns(entity)

	for i := range columns {
		column := &columns[i]
		columnExp := ColumnExp(column)

		if columnExp != nil {
			fields[column.Name] = columnExp
		}
	}
	return boolExp
}

func makeOrderBy(entity *meta.Entity) *graphql.InputObject {
	fields := graphql.InputObjectConfigFieldMap{}

	orderByExp := graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name:   entity.Name + consts.ORDERBY,
			Fields: fields,
		},
	)

	columns := meta.Metas.EntityAllColumns(entity)
	for i := range columns {
		column := &columns[i]
		columnOrderBy := ColumnOrderBy(column)
		if columnOrderBy != nil {
			fields[column.Name] = &graphql.InputObjectFieldConfig{Type: columnOrderBy}
		}
	}
	return orderByExp
}

func makeDistinctOnEnum(entity *meta.Entity) *graphql.Enum {
	enumValueConfigMap := graphql.EnumValueConfigMap{}
	columns := meta.Metas.EntityAllColumns(entity)
	for i := range columns {
		column := &columns[i]
		enumValueConfigMap[column.Name] = &graphql.EnumValueConfig{
			Value: column.Name,
		}
	}

	entEnum := graphql.NewEnum(
		graphql.EnumConfig{
			Name:   entity.Name + consts.DISTINCTEXP,
			Values: enumValueConfigMap,
		},
	)
	return entEnum
}
