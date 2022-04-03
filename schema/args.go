package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/model/graph"
)

func (c *TypeCache) makeArgs() {
	for i := range Model.Grahp.Interfaces {
		c.makeOneEntityArgs(Model.Grahp.Interfaces[i])
	}
	for i := range Model.Grahp.Entities {
		c.makeOneEntityArgs(Model.Grahp.Entities[i])
	}
	c.makeRelaionWhereExp()
}

func (c *TypeCache) makeOneEntityArgs(node graph.Node) {
	whereExp := makeWhereExp(node)
	c.WhereExpMap[node.Name()] = whereExp
	orderByExp := makeOrderBy(node)
	c.OrderByMap[node.Name()] = orderByExp
	distinctOnEnum := makeDistinctOnEnum(node)
	c.DistinctOnEnumMap[node.Name()] = distinctOnEnum
}

func (c *TypeCache) makeRelaionWhereExp() {
	for entityName := range c.WhereExpMap {
		exp := c.WhereExpMap[entityName]
		node := Model.Grahp.GetNodeByUuid(entityName)
		if node == nil {
			panic("Fatal error, can not find entity by name:" + entityName)
		}
		associations := node.Associations()
		for i := range associations {
			assoc := associations[i]
			exp.AddFieldConfig(assoc.Name(), &graphql.InputObjectFieldConfig{
				Type: c.WhereExp(assoc.TypeClass().Name()),
			})
		}
	}
}

func makeWhereExp(node graph.Node) *graphql.InputObject {
	expName := node.Name() + consts.BOOLEXP
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

	attrs := node.Attributes()

	for i := range attrs {
		attr := attrs[i]
		columnExp := AttributeExp(attr)

		if columnExp != nil {
			fields[attr.Name] = columnExp
		}
	}
	return boolExp
}

func makeOrderBy(node graph.Node) *graphql.InputObject {
	fields := graphql.InputObjectConfigFieldMap{}

	orderByExp := graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name:   node.Name() + consts.ORDERBY,
			Fields: fields,
		},
	)

	columns := node.Attributes
	for i := range columns {
		column := columns[i]
		columnOrderBy := ColumnOrderBy(column)
		if columnOrderBy != nil {
			fields[column.Name] = &graphql.InputObjectFieldConfig{Type: columnOrderBy}
		}
	}
	return orderByExp
}

func makeDistinctOnEnum(node graph.Node) *graphql.Enum {
	enumValueConfigMap := graphql.EnumValueConfigMap{}
	columns := node.Attributes
	for i := range columns {
		column := columns[i]
		enumValueConfigMap[column.Name] = &graphql.EnumValueConfig{
			Value: column.Name,
		}
	}

	entEnum := graphql.NewEnum(
		graphql.EnumConfig{
			Name:   node.Name() + consts.DISTINCTEXP,
			Values: enumValueConfigMap,
		},
	)
	return entEnum
}
