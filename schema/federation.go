package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/consts"
)

//union _Entity

var _ServiceType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: consts.SERVICE_TYPE,
		Fields: graphql.Fields{
			consts.ID: &graphql.Field{ //扩展一个id字段
				Type: graphql.Int,
			},
			consts.SDL: &graphql.Field{
				Type:        graphql.String,
				Description: "Service SDL",
			},
		},
		Description: "_Service type of federation schema specification",
	},
)
