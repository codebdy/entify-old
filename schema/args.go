package schema

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/model"
	"rxdrag.com/entity-engine/model/graph"
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

func (c *TypeCache) makeOneEntityArgs(node graph.Node) {
	whereExp := makeWhereExp(node)
	c.WhereExpMap[node.Name()] = whereExp

	orderByExp := makeOrderBy(node)
	if len(orderByExp.Fields()) > 0 {
		c.OrderByMap[node.Name()] = orderByExp
	}

	distinctOnEnum := makeDistinctOnEnum(node)
	c.DistinctOnEnumMap[node.Name()] = distinctOnEnum
}

func (c *TypeCache) makeRelaionWhereExp() {
	for entityName := range c.WhereExpMap {
		exp := c.WhereExpMap[entityName]
		node := model.GlobalModel.Graph.GetNodeByName(entityName)
		if node.Entity() == nil && node.Interface() == nil {
			panic("Fatal error, can not find entity by name:" + entityName)
		}
		associations := node.Associations()
		for i := range associations {
			fmt.Println(i)
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

	attrs := node.Attributes()
	for i := range attrs {
		attr := attrs[i]
		attrOrderBy := AttributeOrderBy(attr)
		if attrOrderBy != nil {
			fields[attr.Name] = &graphql.InputObjectFieldConfig{Type: attrOrderBy}
		}
	}
	return orderByExp
}

func makeDistinctOnEnum(node graph.Node) *graphql.Enum {
	enumValueConfigMap := graphql.EnumValueConfigMap{}
	attrs := node.Attributes()
	for i := range attrs {
		attr := attrs[i]
		enumValueConfigMap[attr.Name] = &graphql.EnumValueConfig{
			Value: attr.Name,
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
