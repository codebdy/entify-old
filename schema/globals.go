package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entify/consts"
	"rxdrag.com/entify/entity"
	"rxdrag.com/entify/utils"
)

var TokenCache = map[string]*entity.User{}
var Cache TypeCache

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
var _ServiceType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: consts.SERVICE_TYPE,
		Fields: graphql.Fields{
			// consts.ID: &graphql.Field{ //扩展一个id字段
			// 	Type: graphql.Int,
			// },
			consts.SDL: &graphql.Field{
				Type:        graphql.String,
				Description: "Service SDL",
			},
		},
		Description: "_Service type of federation schema specification",
	},
)

var EntityType *graphql.Union

var installInputType = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "InstallInput",
		Fields: graphql.InputObjectConfigFieldMap{
			consts.ADMIN: &graphql.InputObjectFieldConfig{
				Type: &graphql.NonNull{
					OfType: graphql.String,
				},
			},
			consts.ADMINPASSWORD: &graphql.InputObjectFieldConfig{
				Type: &graphql.NonNull{
					OfType: graphql.String,
				},
			},
			consts.WITHDEMO: &graphql.InputObjectFieldConfig{
				Type: graphql.Boolean,
			},
		},
	},
)

var baseRoleTye = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "BaseRole",
		Fields: graphql.Fields{
			consts.ID: &graphql.Field{
				Type: graphql.Int,
			},
			consts.NAME: &graphql.Field{
				Type: graphql.String,
			},
		},
		Description: "Base role for auth",
	},
)

var baseUserType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "BaseUser",
		Fields: graphql.Fields{
			consts.ID: &graphql.Field{
				Type: graphql.Int,
			},
			consts.NAME: &graphql.Field{
				Type: graphql.String,
			},
			consts.LOGIN_NAME: &graphql.Field{
				Type: graphql.String,
			},
			"roles": &graphql.Field{
				Type: &graphql.List{
					OfType: baseRoleTye,
				},
			},
		},
		Description: "Base user for auth",
	},
)
