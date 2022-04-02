package schemaold

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/model"
)

func (c *TypeCache) makeArgs() {
	for i := range model.TheModel.Interfaces {
		c.makeOneEntityArgs(model.TheModel.Interfaces[i])
	}
	for i := range model.TheModel.Entities {
		c.makeOneEntityArgs(model.TheModel.Entities[i])
	}
	c.makeRelaionWhereExp()
}

func (c *TypeCache) makeOneEntityArgs(entity *model.Entity) {
	whereExp := makeWhereExp(entity)
	c.WhereExpMap[entity.Name] = whereExp
	orderByExp := makeOrderBy(entity)
	c.OrderByMap[entity.Name] = orderByExp
	distinctOnEnum := makeDistinctOnEnum(entity)
	c.DistinctOnEnumMap[entity.Name] = distinctOnEnum
}

func (c *TypeCache) makeRelaionWhereExp() {
	for entityName := range c.WhereExpMap {
		exp := c.WhereExpMap[entityName]
		entity := model.TheModel.GetEntityOrInterfaceByName(entityName)
		if entity == nil {
			panic("Fatal error, can not find entity by name:" + entityName)
		}
		associations := entity.Associations
		for i := range associations {
			assoc := associations[i]
			exp.AddFieldConfig(assoc.Name, &graphql.InputObjectFieldConfig{
				Type: c.WhereExp(assoc.TypeEntity.Name),
			})
		}
	}
}

func makeWhereExp(entity *model.Entity) *graphql.InputObject {
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

	columns := entity.Columns

	for i := range columns {
		column := columns[i]
		columnExp := ColumnExp(column)

		if columnExp != nil {
			fields[column.Name] = columnExp
		}
	}
	return boolExp
}

func makeOrderBy(entity *model.Entity) *graphql.InputObject {
	fields := graphql.InputObjectConfigFieldMap{}

	orderByExp := graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name:   entity.Name + consts.ORDERBY,
			Fields: fields,
		},
	)

	columns := entity.Columns
	for i := range columns {
		column := columns[i]
		columnOrderBy := ColumnOrderBy(column)
		if columnOrderBy != nil {
			fields[column.Name] = &graphql.InputObjectFieldConfig{Type: columnOrderBy}
		}
	}
	return orderByExp
}

func makeDistinctOnEnum(entity *model.Entity) *graphql.Enum {
	enumValueConfigMap := graphql.EnumValueConfigMap{}
	columns := entity.Columns
	for i := range columns {
		column := columns[i]
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