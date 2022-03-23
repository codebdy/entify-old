package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/authentication"
	"rxdrag.com/entity-engine/authentication/jwt"
	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/meta"
	"rxdrag.com/entity-engine/model"
	"rxdrag.com/entity-engine/resolve"
	"rxdrag.com/entity-engine/utils"
)

func rootMutation() *graphql.Object {
	metaEntity := model.MetaEntity()
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
			Type:    Cache.OutputType(metaEntity),
			Resolve: publishResolve,
		},
		consts.ROLLBACK: &graphql.Field{
			Type:    Cache.OutputType(metaEntity),
			Resolve: resolve.SyncMetaResolve,
		},
		consts.SYNC_META: &graphql.Field{
			Type:    Cache.OutputType(metaEntity),
			Resolve: resolve.SyncMetaResolve,
		},
	}

	for _, entity := range model.TheModel.Entities {
		appendToMutationFields(entity, &mutationFields)
	}

	rootMutation := graphql.ObjectConfig{Name: consts.ROOT_MUTATION_NAME, Fields: mutationFields}

	return graphql.NewObject(rootMutation)
}

func appendToMutationFields(entity *model.Entity, feilds *graphql.Fields) {
	//如果是枚举
	if entity.EntityType == meta.ENTITY_ENUM {
		return
	}

	name := utils.FirstUpper(entity.Name)

	(*feilds)[consts.DELETE+name] = &graphql.Field{
		Type: *Cache.MutationResponse(entity),
		Args: graphql.FieldConfigArgument{
			consts.ARG_WHERE: &graphql.ArgumentConfig{
				Type: Cache.WhereExp(entity),
			},
		},
		//Resolve: entity.QueryResolve(),
	}
	(*feilds)[consts.DELETE+name+consts.BY_ID] = &graphql.Field{
		Type: Cache.OutputType(entity),
		Args: graphql.FieldConfigArgument{
			consts.ID: &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		//Resolve: entity.QueryResolve(),
	}
	(*feilds)[consts.SAVE+name] = &graphql.Field{
		Type: Cache.OutputType(entity),
		Args: graphql.FieldConfigArgument{
			consts.ARG_OBJECTS: &graphql.ArgumentConfig{
				Type: &graphql.NonNull{
					OfType: &graphql.List{
						OfType: &graphql.NonNull{
							OfType: Cache.SaveInput(entity),
						},
					},
				},
			},
		},
	}
	//Resolve: entity.QueryResolve(),
	(*feilds)[consts.SAVE_ONE+name] = &graphql.Field{
		Type: Cache.OutputType(entity),
		Args: graphql.FieldConfigArgument{
			consts.ARG_OBJECT: &graphql.ArgumentConfig{
				Type: &graphql.NonNull{
					OfType: Cache.SaveInput(entity),
				},
			},
		},
		Resolve: resolve.PostOneResolveFn(entity),
	}

	(*feilds)[consts.UPDATE+name] = &graphql.Field{
		Type: *Cache.MutationResponse(entity),
		Args: graphql.FieldConfigArgument{
			consts.ARG_OBJECTS: &graphql.ArgumentConfig{
				Type: &graphql.NonNull{
					OfType: &graphql.List{
						OfType: &graphql.NonNull{
							OfType: Cache.UpdateInput(entity),
						},
					},
				},
			},
			consts.ARG_WHERE: &graphql.ArgumentConfig{
				Type: Cache.WhereExp(entity),
			},
		},
		//Resolve: entity.QueryResolve(),
	}
}
