package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entify/model/graph"
)

func (c *TypeCache) makeOutputInterfaces(interfaces []*graph.Interface) {
	for i := range interfaces {
		intf := interfaces[i]
		c.InterfaceTypeMap[intf.Name()] = c.InterfaceType(intf)
	}
}

func (c *TypeCache) InterfaceType(intf *graph.Interface) *graphql.Interface {
	name := intf.Name()

	return graphql.NewInterface(
		graphql.InterfaceConfig{
			Name:        name,
			Fields:      outputFields(intf),
			Description: intf.Description(),
			ResolveType: resolveTypeFn,
		},
	)
}
