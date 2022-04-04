package schema

import (
	"github.com/graphql-go/graphql"
)

func (c *TypeCache) makeRelations() {
	for i := range Model.Graph.Interfaces {
		intf := Model.Graph.Interfaces[i]
		objectType := c.ObjectTypeMap[intf.Name()]
		for _, assocition := range intf.Associations() {
			objectType.AddFieldConfig(assocition.Name(), &graphql.Field{
				Name:        assocition.Name(),
				Type:        c.OutputType(assocition.TypeClass().Name()),
				Description: assocition.Description(),
			})
		}
	}
	for i := range Model.Graph.Entities {
		entity := Model.Graph.Entities[i]
		objectType := c.ObjectTypeMap[entity.Name()]
		for _, assocition := range entity.Associations() {
			objectType.AddFieldConfig(assocition.Name(), &graphql.Field{
				Name:        assocition.Name(),
				Type:        c.OutputType(assocition.TypeClass().Name()),
				Description: assocition.Description(),
			})
		}
	}
}
