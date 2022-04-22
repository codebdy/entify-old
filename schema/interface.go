package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/model/graph"
	"rxdrag.com/entity-engine/utils"
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
			ResolveType: func(p graphql.ResolveTypeParams) *graphql.Object {
				if value, ok := p.Value.(map[string]interface{}); ok {
					if id, ok := value[consts.ID].(uint64); ok {
						entityInnerId := utils.DecodeEntityInnerId(id)
						return Cache.GetEntityTypeByInnerId(entityInnerId)
					}
				}
				return nil
			},
		},
	)
}
