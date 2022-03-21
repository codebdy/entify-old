package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/meta"
)

func (c *TypeCache) makeRelations() {
	for i := range meta.Metas.Relations {
		relation := &meta.Metas.Relations[i]
		if relation.RelationType != meta.IMPLEMENTS {
			c.makeRelationShip(relation)
		}
	}
}

func (c *TypeCache) makeRelationShip(relation *meta.Relation) {
	sourceEntity := meta.Metas.GetEntityByUuid(relation.SourceId)
	targetEntity := meta.Metas.GetEntityByUuid(relation.TargetId)
	if sourceEntity == nil {
		panic("Can find entity:" + relation.SourceId)
	}

	soureInterfaceType := c.InterfaceTypeMap[sourceEntity.Name]
	sourceField := &graphql.Field{
		Name:        relation.RoleOnSource,
		Type:        c.OutputType(targetEntity),
		Description: relation.DescriptionOnSource,
	}
	if soureInterfaceType != nil {
		soureInterfaceType.AddFieldConfig(relation.RoleOnSource, sourceField)
		children := meta.Metas.Children(sourceEntity)
		for i := range children {
			childType := c.ObjectTypeMap[children[i].Name]
			childType.AddFieldConfig(relation.RoleOnSource, sourceField)
		}
	} else {
		soureObjectType := c.ObjectTypeMap[sourceEntity.Name]
		if soureObjectType == nil {
			panic("Can find entity Type in map:" + sourceEntity.Name)
		}
		soureObjectType.AddFieldConfig(relation.RoleOnSource, sourceField)
	}

	targetInterfaceType := c.InterfaceTypeMap[targetEntity.Name]
	targetField := &graphql.Field{
		Name:        relation.RoleOnSource,
		Type:        c.OutputType(sourceEntity),
		Description: relation.DescriptionOnTarget,
	}
	if targetInterfaceType != nil {
		targetInterfaceType.AddFieldConfig(relation.RoleOnTarget, targetField)
		children := meta.Metas.Children(targetEntity)
		for i := range children {
			childType := c.ObjectTypeMap[children[i].Name]
			childType.AddFieldConfig(relation.RoleOnTarget, targetField)
		}
	} else {
		targetObjectType := c.ObjectTypeMap[targetEntity.Name]
		if targetObjectType == nil {
			panic("Can find entity Type in map:" + targetEntity.Name)
		}
		targetObjectType.AddFieldConfig(relation.RoleOnTarget, targetField)
	}
}
