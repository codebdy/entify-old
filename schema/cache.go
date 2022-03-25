package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/model"
	"rxdrag.com/entity-engine/utils"
)

var Cache TypeCache

type TypeCache struct {
	ObjectTypeMap        map[string]*graphql.Object
	EnumTypeMap          map[string]*graphql.Enum
	InterfaceTypeMap     map[string]*graphql.Interface
	UpdateInputMap       map[string]*graphql.InputObject
	SaveInputMap         map[string]*graphql.InputObject
	WhereExpMap          map[string]*graphql.InputObject
	DistinctOnEnumMap    map[string]*graphql.Enum
	OrderByMap           map[string]*graphql.InputObject
	EnumComparisonExpMap map[string]*graphql.InputObjectFieldConfig
	MutationResponseMap  map[string]*graphql.Output
	AggregateMap         map[string]*graphql.Output
}

var NodeInterfaceType = graphql.NewInterface(
	graphql.InterfaceConfig{
		Name: utils.FirstUpper(consts.NODE),
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
			},
		},
		Description: "Node interface",
	},
)

func (c *TypeCache) MakeCache() {
	c.clearCache()
	c.makeEnums(model.TheModel.Enums)
	c.makeOutputInterfaces(model.TheModel.Interfaces)
	c.makeOutputObjects(model.TheModel.Entities)
	c.makeRelations()
	c.makeArgs()
	c.makeInputs()
}

func (c *TypeCache) OutputInterfaceType(entity *model.Entity) graphql.Type {
	return c.InterfaceTypeMap[entity.Name]
}

func (c *TypeCache) OutputObjectType(entity *model.Entity) graphql.Type {
	return c.ObjectTypeMap[entity.Name]
}

func (c *TypeCache) EnumType(entity *model.Enum) graphql.Type {
	return c.EnumTypeMap[entity.Name]
}

func (c *TypeCache) WhereExp(entity *model.Entity) *graphql.InputObject {
	return c.WhereExpMap[entity.Name]
}

func (c *TypeCache) OrderByExp(entity *model.Entity) *graphql.InputObject {
	return c.OrderByMap[entity.Name]
}

func (c *TypeCache) DistinctOnEnum(entity *model.Entity) *graphql.Enum {
	return c.DistinctOnEnumMap[entity.Name]
}

func (c *TypeCache) SaveInput(entity *model.Entity) *graphql.InputObject {
	return c.SaveInputMap[entity.Name]
}

func (c *TypeCache) UpdateInput(entity *model.Entity) *graphql.InputObject {
	return c.UpdateInputMap[entity.Name]
}

func (c *TypeCache) MutationResponse(entity *model.Entity) *graphql.Output {
	return c.MutationResponseMap[entity.Name]
}

func (c *TypeCache) mapInterfaces(entities []*model.Interface) []*graphql.Interface {
	interfaces := []*graphql.Interface{NodeInterfaceType}
	for i := range entities {
		interfaces = append(interfaces, c.InterfaceTypeMap[entities[i].Name])
	}

	return interfaces
}

func (c *TypeCache) clearCache() {
	c.ObjectTypeMap = make(map[string]*graphql.Object)
	c.EnumTypeMap = make(map[string]*graphql.Enum)
	c.InterfaceTypeMap = make(map[string]*graphql.Interface)
	c.UpdateInputMap = make(map[string]*graphql.InputObject)
	c.SaveInputMap = make(map[string]*graphql.InputObject)
	c.WhereExpMap = make(map[string]*graphql.InputObject)
	c.DistinctOnEnumMap = make(map[string]*graphql.Enum)
	c.OrderByMap = make(map[string]*graphql.InputObject)
	c.EnumComparisonExpMap = make(map[string]*graphql.InputObjectFieldConfig)
	c.MutationResponseMap = make(map[string]*graphql.Output)
	c.AggregateMap = make(map[string]*graphql.Output)
}
