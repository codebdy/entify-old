package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entify/model"
	"rxdrag.com/entify/model/graph"
	"rxdrag.com/entify/resolve"
)

func (c *TypeCache) makeRelations() {
	for i := range model.GlobalModel.Graph.Interfaces {
		intf := model.GlobalModel.Graph.Interfaces[i]
		interfaceType := c.InterfaceTypeMap[intf.Name()]
		if interfaceType == nil {
			panic("Can find object type:" + intf.Name())
		}
		for _, association := range intf.AllAssociations() {
			if interfaceType.Fields()[association.Name()] != nil {
				panic("Duplicate interface field: " + intf.Name() + "." + association.Name())
			}
			interfaceType.AddFieldConfig(association.Name(), &graphql.Field{
				Name:        association.Name(),
				Type:        c.AssociationType(association),
				Description: association.Description(),
				Resolve:     resolve.QueryAssociationFn(association),
				Args:        quryeArgs(association.TypeClass().Name()),
			})
		}
	}
	for i := range model.GlobalModel.Graph.Entities {
		entity := model.GlobalModel.Graph.Entities[i]
		objectType := c.ObjectTypeMap[entity.Name()]
		for _, association := range entity.AllAssociations() {
			if objectType.Fields()[association.Name()] != nil {
				panic("Duplicate entity field: " + entity.Name() + "." + association.Name())
			}
			objectType.AddFieldConfig(association.Name(), &graphql.Field{
				Name:        association.Name(),
				Type:        c.AssociationType(association),
				Description: association.Description(),
				Resolve:     resolve.QueryAssociationFn(association),
				Args:        quryeArgs(association.TypeClass().Name()),
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
