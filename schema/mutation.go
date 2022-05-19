package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entify/authentication"
	"rxdrag.com/entify/authentication/jwt"
	"rxdrag.com/entify/consts"
	"rxdrag.com/entify/model"
	"rxdrag.com/entify/model/graph"
	"rxdrag.com/entify/resolve"
	"rxdrag.com/entify/utils"
)

const INPUT = "input"

var installInputType = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "InstallInput",
		Fields: graphql.InputObjectConfigFieldMap{
			consts.DB_DRIVER: &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			consts.DB_HOST: &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			consts.DB_PORT: &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			consts.DB_DATABASE: &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			consts.DB_USER: &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			consts.DB_PASSWORD: &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			consts.SERVICE_ID: &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			consts.ADMIN: &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			consts.ADMINPASSWORD: &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			consts.WITHDEMO: &graphql.InputObjectFieldConfig{
				Type: graphql.Boolean,
			},
			consts.URL: &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
		},
	},
)

func rootMutation() *graphql.Object {
	metaEntity := model.GlobalModel.Graph.GetMetaEntity()
	mutationFields := graphql.Fields{
		consts.LOGIN: &graphql.Field{
			Type: graphql.String,
			Args: graphql.FieldConfigArgument{
				consts.LOGIN_NAME: &graphql.ArgumentConfig{
					Type: &graphql.NonNull{OfType: graphql.String},
				},
				consts.PASSWORD: &graphql.ArgumentConfig{
					Type: &graphql.NonNull{OfType: graphql.String},
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				loginName, err := authentication.Login(p.Args[consts.LOGIN_NAME].(string), p.Args[consts.PASSWORD].(string))
				if err != nil {
					return "", err
				}
				return jwt.GenerateToken(loginName)
			},
		},
		consts.LOGOUT: &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world2", nil
			},
		},
		consts.PUBLISH: &graphql.Field{
			Type:    Cache.OutputType(metaEntity.Name()),
			Resolve: publishResolve,
		},
		consts.ROLLBACK: &graphql.Field{
			Type:    Cache.OutputType(metaEntity.Name()),
			Resolve: resolve.SyncMetaResolve,
		},
		consts.SYNC_META: &graphql.Field{
			Type:    Cache.OutputType(metaEntity.Name()),
			Resolve: resolve.SyncMetaResolve,
		},
	}

	for _, entity := range model.GlobalModel.Graph.RootEnities() {
		if entity.Domain.Root {
			appendToMutationFields(entity, &mutationFields)
		}
	}

	for _, service := range model.GlobalModel.Graph.RootServices() {
		appendServiceMutationFields(service, &mutationFields)
	}

	// if !config.GetBool(consts.INSTALLED) {
	// 	mutationFields["install"] = &graphql.Field{
	// 		Type: graphql.Boolean,
	// 		Args: graphql.FieldConfigArgument{
	// 			INPUT: &graphql.ArgumentConfig{
	// 				Type: &graphql.NonNull{
	// 					OfType: installInputType,
	// 				},
	// 			},
	// 		},
	// 		Resolve: installResolve,
	// 	}
	// }

	rootMutation := graphql.ObjectConfig{
		Name:        consts.ROOT_MUTATION_NAME,
		Fields:      mutationFields,
		Description: "Root mutation of entity engine.",
	}

	return graphql.NewObject(rootMutation)
}

func appendToMutationFields(entity *graph.Entity, feilds *graphql.Fields) {
	name := utils.FirstUpper(entity.Name())

	(*feilds)[consts.DELETE+name] = &graphql.Field{
		Type: *Cache.MutationResponse(entity.Name()),
		Args: graphql.FieldConfigArgument{
			consts.ARG_WHERE: &graphql.ArgumentConfig{
				Type: Cache.WhereExp(entity.Name()),
			},
		},
		//Resolve: entity.QueryResolve(),
	}
	(*feilds)[consts.DELETE+name+consts.BY_ID] = &graphql.Field{
		Type: Cache.OutputType(entity.Name()),
		Args: graphql.FieldConfigArgument{
			consts.ID: &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		//Resolve: entity.QueryResolve(),
	}
	(*feilds)[consts.UPSERT+name] = &graphql.Field{
		Type: Cache.OutputType(entity.Name()),
		Args: graphql.FieldConfigArgument{
			consts.ARG_OBJECTS: &graphql.ArgumentConfig{
				Type: &graphql.NonNull{
					OfType: &graphql.List{
						OfType: &graphql.NonNull{
							OfType: Cache.SaveInput(entity.Name()),
						},
					},
				},
			},
		},
	}
	//Resolve: entity.QueryResolve(),
	(*feilds)[consts.UPSERT_ONE+name] = &graphql.Field{
		Type: Cache.OutputType(entity.Name()),
		Args: graphql.FieldConfigArgument{
			consts.ARG_OBJECT: &graphql.ArgumentConfig{
				Type: &graphql.NonNull{
					OfType: Cache.SaveInput(entity.Name()),
				},
			},
		},
		Resolve: resolve.PostOneResolveFn(entity),
	}

	updateInput := Cache.UpdateInput(entity.Name())
	if len(updateInput.Fields()) > 0 {
		(*feilds)[consts.UPDATE+name] = &graphql.Field{
			Type: *Cache.MutationResponse(entity.Name()),
			Args: graphql.FieldConfigArgument{
				consts.ARG_OBJECTS: &graphql.ArgumentConfig{
					Type: &graphql.NonNull{
						OfType: &graphql.List{
							OfType: &graphql.NonNull{
								OfType: updateInput,
							},
						},
					},
				},
				consts.ARG_WHERE: &graphql.ArgumentConfig{
					Type: Cache.WhereExp(entity.Name()),
				},
			},
			//Resolve: entity.QueryResolve(),
		}
	}

}
