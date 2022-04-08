package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/model"
	"rxdrag.com/entity-engine/model/graph"
)

func (c *TypeCache) makeRelations() {
	for i := range model.GlobalModel.Graph.Interfaces {
		intf := model.GlobalModel.Graph.Interfaces[i]
		interfaceType := c.InterfaceTypeMap[intf.Name()]
		if interfaceType == nil {
			panic("Can find object type:" + intf.Name())
		}
		for _, association := range intf.AllAssociations() {
			interfaceType.AddFieldConfig(association.Name(), &graphql.Field{
				Name:        association.Name(),
				Type:        c.AssociationType(association),
				Description: association.Description(),
			})
		}
	}
	for i := range model.GlobalModel.Graph.Entities {
		entity := model.GlobalModel.Graph.Entities[i]
		objectType := c.ObjectTypeMap[entity.Name()]
		for _, association := range entity.AllAssociations() {
			objectType.AddFieldConfig(association.Name(), &graphql.Field{
				Name:        association.Name(),
				Type:        c.AssociationType(association),
				Description: association.Description(),
			})
		}
	}
}

func (c *TypeCache) AssociationType(association *graph.Association) graphql.Output {
	if association.IsArray() {
		return &graphql.NonNull{
			OfType: &graphql.List{
				OfType: c.OutputType(association.TypeClass().Name()),
			},
		}
	} else {
		return c.OutputType(association.TypeClass().Name())
	}
}
