package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/meta"
)

var Cache TypeCache

type TypeCache struct {
	ObjectTypeMap_old    map[string]*graphql.Output
	ObjectTypeMap        map[string]*graphql.Object
	EnumTypeMap          map[string]*graphql.Enum
	InterfaceTypeMap     map[string]*graphql.Interface
	UpdateInputMap       map[string]*graphql.Input
	PostInputMap         map[string]*graphql.Input
	WhereExpMap          map[string]*graphql.InputObject
	DistinctOnEnumMap    map[string]*graphql.Enum
	OrderByMap           map[string]*graphql.InputObject
	EnumComparisonExpMap map[string]*graphql.InputObjectFieldConfig
	MutationResponseMap  map[string]*graphql.Output
	AggregateMap         map[string]*graphql.Output
}

//where表达式缓存，query跟mutation都用

func (c *TypeCache) MakeCache() {
	c.clearCache()
	for i := range meta.Metas.Entities {
		entity := &meta.Metas.Entities[i]
		if entity.EntityType == meta.ENTITY_ENUM {
			c.EnumTypeMap[entity.Name] = EnumType(entity)
		} else {
			if meta.Metas.HasChildren(entity) {
				c.InterfaceTypeMap[entity.Name] = InterfaceType(entity)
			} else {
				c.ObjectTypeMap[entity.Name] = ObjectType(entity)
			}
		}
	}

	for i := range meta.Metas.Relations {
		relation := meta.Metas.Relations[i]
		sourceEntity := meta.Metas.GetEntityByUuid(relation.SourceId)
		targetEntity := meta.Metas.GetEntityByUuid(relation.TargetId)
		if sourceEntity == nil {
			panic("Can find entity:" + relation.SourceId)
		}

		soureInterfaceType := c.InterfaceTypeMap[sourceEntity.Name]
		sourceField := &graphql.Field{
			Name: relation.RoleOnSource,
			Type: c.OutputType(targetEntity),
		}
		if soureInterfaceType != nil {
			soureInterfaceType.AddFieldConfig(relation.RoleOnSource, sourceField)
		} else {
			soureObjectType := c.ObjectTypeMap[sourceEntity.Name]
			if soureObjectType == nil {
				panic("Can find entity Type in map:" + sourceEntity.Name)
			}
			soureObjectType.AddFieldConfig(relation.RoleOnSource, sourceField)
		}

		targetInterfaceType := c.InterfaceTypeMap[targetEntity.Name]
		targetField := &graphql.Field{
			Name: relation.RoleOnSource,
			Type: c.OutputType(sourceEntity),
		}
		if targetInterfaceType != nil {
			targetInterfaceType.AddFieldConfig(relation.RoleOnTarget, targetField)
		} else {
			targetObjectType := c.ObjectTypeMap[targetEntity.Name]
			if targetObjectType == nil {
				panic("Can find entity Type in map:" + targetEntity.Name)
			}
			targetObjectType.AddFieldConfig(relation.RoleOnTarget, targetField)
		}
	}
}

func (c *TypeCache) OutputType(entity *meta.Entity) graphql.Type {
	if entity.EntityType == meta.ENTITY_ENUM {
		return c.EnumTypeMap[entity.Name]
	} else {
		if meta.Metas.HasChildren(entity) {
			return c.InterfaceTypeMap[entity.Name]
		} else {
			return c.ObjectTypeMap[entity.Name]
		}
	}
}

func (c *TypeCache) clearCache() {
	c.ObjectTypeMap_old = make(map[string]*graphql.Output)
	c.ObjectTypeMap = make(map[string]*graphql.Object)
	c.EnumTypeMap = make(map[string]*graphql.Enum)
	c.InterfaceTypeMap = make(map[string]*graphql.Interface)
	c.UpdateInputMap = make(map[string]*graphql.Input)
	c.PostInputMap = make(map[string]*graphql.Input)
	c.WhereExpMap = make(map[string]*graphql.InputObject)
	c.DistinctOnEnumMap = make(map[string]*graphql.Enum)
	c.OrderByMap = make(map[string]*graphql.InputObject)
	c.EnumComparisonExpMap = make(map[string]*graphql.InputObjectFieldConfig)
	c.MutationResponseMap = make(map[string]*graphql.Output)
	c.AggregateMap = make(map[string]*graphql.Output)
}
