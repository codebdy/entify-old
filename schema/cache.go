package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entify/model"
	"rxdrag.com/entify/model/graph"
)

type TypeCache struct {
	ObjectTypeMap        map[string]*graphql.Object
	ObjectMapById        map[uint64]*graphql.Object
	EnumTypeMap          map[string]*graphql.Enum
	InterfaceTypeMap     map[string]*graphql.Interface
	UpdateInputMap       map[string]*graphql.InputObject
	SaveInputMap         map[string]*graphql.InputObject
	HasManyInputMap      map[string]*graphql.InputObject
	HasOneInputMap       map[string]*graphql.InputObject
	WhereExpMap          map[string]*graphql.InputObject
	DistinctOnEnumMap    map[string]*graphql.Enum
	OrderByMap           map[string]*graphql.InputObject
	EnumComparisonExpMap map[string]*graphql.InputObjectFieldConfig
	MutationResponseMap  map[string]*graphql.Object
	AggregateMap         map[string]*graphql.Object
}

func (c *TypeCache) MakeCache() {
	c.clearCache()
	c.makeEnums(model.GlobalModel.Graph.Enums)
	c.makeOutputInterfaces(model.GlobalModel.Graph.Interfaces)
	c.makeOutputObjects(model.GlobalModel.Graph.Entities)
	c.makeArgs()
	c.makeRelations()
	c.makeInputs()
}

// func (c *TypeCache) OutputInterfaceType(entity *model.Entity) graphql.Type {
// 	return c.InterfaceTypeMap[entity.Name]
// }

func (c *TypeCache) InterfaceOutputType(name string) *graphql.Interface {
	intf := c.InterfaceTypeMap[name]
	if intf != nil {
		return intf
	}
	panic("Can not find interface output type of " + name)
}

func (c *TypeCache) EntityeOutputType(name string) *graphql.Object {
	obj := c.ObjectTypeMap[name]
	if obj == nil {
		panic("Can not find output type of " + name)
	}
	return obj
}

func (c *TypeCache) OutputType(name string) graphql.Type {
	intf := c.InterfaceTypeMap[name]
	if intf != nil {
		return intf
	}
	obj := c.ObjectTypeMap[name]
	if obj == nil {
		panic("Can not find output type of " + name)
	}
	return obj
}

func (c *TypeCache) GetEntityTypeByInnerId(id uint64) *graphql.Object {
	return c.ObjectMapById[id]
}

func (c *TypeCache) EntityTypes() []graphql.Type {
	objs := []graphql.Type{}
	for key := range c.ObjectTypeMap {
		objs = append(objs, c.ObjectTypeMap[key])
	}

	return objs
}

func (c *TypeCache) EntityObjects() []*graphql.Object {
	objs := []*graphql.Object{}
	for key := range c.ObjectTypeMap {
		objs = append(objs, c.ObjectTypeMap[key])
	}

	return objs
}

func (c *TypeCache) EnumType(name string) *graphql.Enum {
	return c.EnumTypeMap[name]
}

func (c *TypeCache) WhereExp(name string) *graphql.InputObject {
	return c.WhereExpMap[name]
}

func (c *TypeCache) OrderByExp(name string) *graphql.InputObject {
	return c.OrderByMap[name]
}

func (c *TypeCache) DistinctOnEnum(name string) *graphql.Enum {
	return c.DistinctOnEnumMap[name]
}

func (c *TypeCache) DistinctOnEnums() map[string]*graphql.Enum {
	return c.DistinctOnEnumMap
}

func (c *TypeCache) SaveInput(name string) *graphql.InputObject {
	return c.SaveInputMap[name]
}

func (c *TypeCache) UpdateInput(name string) *graphql.InputObject {
	return c.UpdateInputMap[name]
}
func (c *TypeCache) HasManyInput(name string) *graphql.InputObject {
	return c.HasManyInputMap[name]
}
func (c *TypeCache) HasOneInput(name string) *graphql.InputObject {
	return c.HasOneInputMap[name]
}

func (c *TypeCache) MutationResponse(name string) *graphql.Object {
	return c.MutationResponseMap[name]
}

func (c *TypeCache) mapInterfaces(entities []*graph.Interface) []*graphql.Interface {
	interfaces := []*graphql.Interface{}
	for i := range entities {
		interfaces = append(interfaces, c.InterfaceTypeMap[entities[i].Name()])
	}

	return interfaces
}

func (c *TypeCache) clearCache() {
	c.ObjectTypeMap = make(map[string]*graphql.Object)
	c.ObjectMapById = make(map[uint64]*graphql.Object)
	c.EnumTypeMap = make(map[string]*graphql.Enum)
	c.InterfaceTypeMap = make(map[string]*graphql.Interface)
	c.UpdateInputMap = make(map[string]*graphql.InputObject)
	c.SaveInputMap = make(map[string]*graphql.InputObject)
	c.HasManyInputMap = make(map[string]*graphql.InputObject)
	c.HasOneInputMap = make(map[string]*graphql.InputObject)
	c.WhereExpMap = make(map[string]*graphql.InputObject)
	c.DistinctOnEnumMap = make(map[string]*graphql.Enum)
	c.OrderByMap = make(map[string]*graphql.InputObject)
	c.EnumComparisonExpMap = make(map[string]*graphql.InputObjectFieldConfig)
	c.MutationResponseMap = make(map[string]*graphql.Object)
	c.AggregateMap = make(map[string]*graphql.Object)
}
