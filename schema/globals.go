package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entify/config"
	"rxdrag.com/entify/consts"
	"rxdrag.com/entify/utils"
)

var Cache TypeCache

var _ServiceType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: consts.SERVICE_TYPE,
		Fields: graphql.Fields{
			consts.ID: &graphql.Field{ //扩展一个id字段
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					defer utils.PrintErrorStack()
					return config.ServiceId(), nil
				},
			},

			//扩展字段
			consts.INSTALLED: &graphql.Field{
				Type: graphql.Boolean,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					defer utils.PrintErrorStack()
					return true, nil
				},
			},

			consts.CAN_UPLOAD: &graphql.Field{
				Type: graphql.Boolean,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					defer utils.PrintErrorStack()
					return config.Storage() != "", nil
				},
			},

			consts.SDL: &graphql.Field{
				Type:        graphql.String,
				Description: "Service SDL",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					defer utils.PrintErrorStack()
					return makeFederationSDL(), nil
				},
			},
		},
		Description: "_Service type of federation schema specification, and extends other fields",
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

var fileOutputType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: consts.FILE,
		Fields: graphql.Fields{
			consts.FILE_NAME: &graphql.Field{
				Type: graphql.String,
			},
			consts.FILE_SIZE: &graphql.Field{
				Type: graphql.Int,
			},
			consts.FILE_MIMETYPE: &graphql.Field{
				Type: graphql.String,
			},
			consts.FILE_URL: &graphql.Field{
				Type: graphql.String,
			},
			consts.File_EXTNAME: &graphql.Field{
				Type: graphql.String,
			},
			consts.FILE_THMUBNAIL: &graphql.Field{
				Type: graphql.String,
			},
			consts.FILE_RESIZE: &graphql.Field{
				Type: graphql.String,
				Args: graphql.FieldConfigArgument{
					consts.FILE_WIDTH: &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					consts.FILE_HEIGHT: &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
			},
		},
		Description: "File type",
	},
)

var baseRoleTye = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "BaseRole",
		Fields: graphql.Fields{
			consts.ID: &graphql.Field{
				Type: graphql.ID,
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
				Type: graphql.ID,
			},
			consts.NAME: &graphql.Field{
				Type: graphql.String,
			},
			consts.LOGIN_NAME: &graphql.Field{
				Type: graphql.String,
			},
			consts.IS_SUPPER: &graphql.Field{
				Type: graphql.Boolean,
			},
			consts.IS_DEMO: &graphql.Field{
				Type: graphql.Boolean,
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
