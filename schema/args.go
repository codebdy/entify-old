package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entify/consts"
	"rxdrag.com/entify/model"
	"rxdrag.com/entify/model/graph"
)

func (c *TypeCache) makeArgs() {
	for i := range model.GlobalModel.Graph.Interfaces {
		c.makeOneEntityArgs(model.GlobalModel.Graph.Interfaces[i])
	}
	for i := range model.GlobalModel.Graph.Entities {
		c.makeOneEntityArgs(model.GlobalModel.Graph.Entities[i])
	}
	c.makeRelaionWhereExp()
}

func (c *TypeCache) makeOneEntityArgs(node graph.Noder) {
	whereExp := makeWhereExp(node)
	c.WhereExpMap[extractName(node)] = whereExp

	orderByExp := makeOrderBy(node)
	if len(orderByExp.Fields()) > 0 {
		c.OrderByMap[extractName(node)] = orderByExp
	}

	distinctOnEnum := makeDistinctOnEnum(node)
	c.DistinctOnEnumMap[extractName(node)] = distinctOnEnum
}

func (c *TypeCache) makeRelaionWhereExp() {
	for entityName := range c.WhereExpMap {
		exp := c.WhereExpMap[entityName]
		node := model.GlobalModel.Graph.GetNodeByName(entityName)
		if node.Entity() == nil && node.Interface() == nil {
			panic("Fatal error, can not find entity by name:" + entityName)
		}
		associations := node.AllAssociations()
		for i := range associations {
			assoc := associations[i]
			exp.AddFieldConfig(assoc.Name(), &graphql.InputObjectFieldConfig{
				Type: c.WhereExp(extractName(assoc.TypeClass())),
			})
		}
	}
}

func makeWhereExp(node graph.Noder) *graphql.InputObject {
	expName := extractName(node) + consts.BOOLEXP
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

	attrs := node.AllAttributes()

	for i := range attrs {
		attr := attrs[i]
		columnExp := AttributeExp(attr)

		if columnExp != nil {
			fields[attr.Name] = columnExp
		}
	}
	return boolExp
}

func makeOrderBy(node graph.Noder) *graphql.InputObject {
	fields := graphql.InputObjectConfigFieldMap{}

	orderByExp := graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name:   extractName(node) + consts.ORDERBY,
			Fields: fields,
		},
	)

	attrs := node.AllAttributes()
	for i := range attrs {
		attr := attrs[i]
		attrOrderBy := AttributeOrderBy(attr)
		if attrOrderBy != nil {
			fields[attr.Name] = &graphql.InputObjectFieldConfig{Type: attrOrderBy}
		}
	}
	return orderByExp
}

func makeDistinctOnEnum(node graph.Noder) *graphql.Enum {
	enumValueConfigMap := graphql.EnumValueConfigMap{}
	attrs := node.AllAttributes()
	for i := range attrs {
		attr := attrs[i]
		enumValueConfigMap[attr.Name] = &graphql.EnumValueConfig{
			Value: attr.Name,
		}
	}

	entEnum := graphql.NewEnum(
		graphql.EnumConfig{
			Name:   extractName(node) + consts.DISTINCTEXP,
			Values: enumValueConfigMap,
		},
	)
	return entEnum
}
