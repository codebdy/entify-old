package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/model/graph"
)

func (c *TypeCache) makeArgs() {
	for i := range Model.graph.Interfaces {
		c.makeOneEntityArgs(Model.graph.Interfaces[i])
	}
	for i := range Model.graph.Entities {
		c.makeOneEntityArgs(Model.graph.Entities[i])
	}
	c.makeRelaionWhereExp()
}

func (c *TypeCache) makeOneEntityArgs(entity *graph.Entity) {
	whereExp := makeWhereExp(entity)
	c.WhereExpMap[entity.Name()] = whereExp
	orderByExp := makeOrderBy(entity)
	c.OrderByMap[entity.Name()] = orderByExp
	distinctOnEnum := makeDistinctOnEnum(entity)
	c.DistinctOnEnumMap[entity.Name()] = distinctOnEnum
}

func (c *TypeCache) makeRelaionWhereExp() {
	for entityName := range c.WhereExpMap {
		exp := c.WhereExpMap[entityName]
		entity := Model.graph.GetEntityOrInterfaceByName(entityName)
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

func makeWhereExp(entity *graph.Entity) *graphql.InputObject {
	expName := entity.Name() + consts.BOOLEXP
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

	columns := entity.Attributes

	for i := range columns {
		column := columns[i]
		columnExp := ColumnExp(column)

		if columnExp != nil {
			fields[column.Name] = columnExp
		}
	}
	return boolExp
}

func makeOrderBy(entity *graph.Entity) *graphql.InputObject {
	fields := graphql.InputObjectConfigFieldMap{}

	orderByExp := graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name:   entity.Name() + consts.ORDERBY,
			Fields: fields,
		},
	)

	columns := entity.Attributes
	for i := range columns {
		column := columns[i]
		columnOrderBy := ColumnOrderBy(column)
		if columnOrderBy != nil {
			fields[column.Name] = &graphql.InputObjectFieldConfig{Type: columnOrderBy}
		}
	}
	return orderByExp
}

func makeDistinctOnEnum(entity *graph.Entity) *graphql.Enum {
	enumValueConfigMap := graphql.EnumValueConfigMap{}
	columns := entity.Attributes
	for i := range columns {
		column := columns[i]
		enumValueConfigMap[column.Name] = &graphql.EnumValueConfig{
			Value: column.Name,
		}
	}

	entEnum := graphql.NewEnum(
		graphql.EnumConfig{
			Name:   entity.Name() + consts.DISTINCTEXP,
			Values: enumValueConfigMap,
		},
	)
	return entEnum
}
