package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entify/consts"
	"rxdrag.com/entify/model/graph"
	"rxdrag.com/entify/utils"
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

func resolveTypeFn(p graphql.ResolveTypeParams) *graphql.Object {
	if value, ok := p.Value.(map[string]interface{}); ok {
		if id, ok := value[consts.ID].(uint64); ok {
			entityInnerId := utils.DecodeEntityInnerId(id)
			return Cache.GetEntityTypeByInnerId(entityInnerId)
		}
	}
	return nil
}
