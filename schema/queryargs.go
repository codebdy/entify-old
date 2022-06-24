package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entify/consts"
	"rxdrag.com/entify/model"
	"rxdrag.com/entify/model/graph"
)

func (c *TypeCache) makeArgs() {
	for i := range model.GlobalModel.Graph.Interfaces {
		c.makeOneInterfaceArgs(model.GlobalModel.Graph.Interfaces[i])
	}
	for i := range model.GlobalModel.Graph.Entities {
		c.makeOneEntityArgs(model.GlobalModel.Graph.Entities[i])
	}
	for i := range model.GlobalModel.Graph.Partials {
		c.makeOnePartailArgs(model.GlobalModel.Graph.Partials[i])
	}

	for i := range model.GlobalModel.Graph.Externals {
		c.makeOneExternalArgs(model.GlobalModel.Graph.Externals[i])
	}
	c.makeRelaionWhereExp()
}

func (c *TypeCache) makeOneEntityArgs(entity *graph.Entity) {
	c.makeOneArgs(entity.Name(), entity.AllAttributes())
}

func (c *TypeCache) makeOneInterfaceArgs(intf *graph.Interface) {
	c.makeOneArgs(intf.Name(), intf.AllAttributes())
}

func (c *TypeCache) makeOnePartailArgs(partial *graph.Partial) {
	c.makeOneArgs(partial.Name(), partial.AllAttributes())
}

func (c *TypeCache) makeOneExternalArgs(external *graph.External) {
	c.makeOneArgs(external.Name(), external.AllAttributes())
}

func (c *TypeCache) makeOneArgs(name string, attrs []*graph.Attribute) {
	whereExp := makeWhereExp(name, attrs)
	c.WhereExpMap[name] = whereExp

	orderByExp := makeOrderBy(name, attrs)
	if len(orderByExp.Fields()) > 0 {
		c.OrderByMap[name] = orderByExp
	}

	distinctOnEnum := makeDistinctOnEnum(name, attrs)
	c.DistinctOnEnumMap[name] = distinctOnEnum
}

func (c *TypeCache) makeRelaionWhereExp() {
	for className := range c.WhereExpMap {
		exp := c.WhereExpMap[className]
		intf := model.GlobalModel.Graph.GetInterfaceByName(className)
		entity := model.GlobalModel.Graph.GetEntityByName(className)
		partial := model.GlobalModel.Graph.GetPartialByName(className)
		external := model.GlobalModel.Graph.GetExternalByName(className)
		if intf == nil && entity == nil && partial == nil && external == nil {
			panic("Fatal error, can not find class by name:" + className)
		}
		var associations []*graph.Association
		if intf != nil {
			associations = intf.AllAssociations()
		} else if entity != nil {
			associations = entity.AllAssociations()
		} else if partial != nil {
			associations = partial.AllAssociations()
		} else if external != nil {
			associations = external.AllAssociations()
		}
		for i := range associations {
			assoc := associations[i]
			exp.AddFieldConfig(assoc.Name(), &graphql.InputObjectFieldConfig{
				Type: c.WhereExp(assoc.TypeClass().Name()),
			})
		}
	}
}

func makeWhereExp(name string, attrs []*graph.Attribute) *graphql.InputObject {
	expName := name + consts.BOOLEXP
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

	for i := range attrs {
		attr := attrs[i]
		columnExp := AttributeExp(attr)

		if columnExp != nil {
			fields[attr.Name] = columnExp
		}
	}
	return boolExp
}

func makeOrderBy(name string, attrs []*graph.Attribute) *graphql.InputObject {
	fields := graphql.InputObjectConfigFieldMap{}

	orderByExp := graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name:   name + consts.ORDERBY,
			Fields: fields,
		},
	)

	for i := range attrs {
		attr := attrs[i]
		attrOrderBy := AttributeOrderBy(attr)
		if attrOrderBy != nil {
			fields[attr.Name] = &graphql.InputObjectFieldConfig{Type: attrOrderBy}
		}
	}
	return orderByExp
}

func makeDistinctOnEnum(name string, attrs []*graph.Attribute) *graphql.Enum {
	enumValueConfigMap := graphql.EnumValueConfigMap{}
	for i := range attrs {
		attr := attrs[i]
		enumValueConfigMap[attr.Name] = &graphql.EnumValueConfig{
			Value: attr.Name,
		}
	}

	entEnum := graphql.NewEnum(
		graphql.EnumConfig{
			Name:   name + consts.DISTINCTEXP,
			Values: enumValueConfigMap,
		},
	)
	return entEnum
}
