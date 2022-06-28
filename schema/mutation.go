package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entify/authentication"
	"rxdrag.com/entify/config"
	"rxdrag.com/entify/consts"
	"rxdrag.com/entify/model"
	"rxdrag.com/entify/model/graph"
	"rxdrag.com/entify/resolve"
	"rxdrag.com/entify/scalars"
	"rxdrag.com/entify/utils"
)

const INPUT = "input"

func appendAuthMutation(fields graphql.Fields) {
	fields[consts.LOGIN] = &graphql.Field{
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
			defer utils.PrintErrorStack()
			return authentication.Login(p.Args[consts.LOGIN_NAME].(string), p.Args[consts.PASSWORD].(string))
		},
	}

	fields[consts.LOGOUT] = &graphql.Field{
		Type:    graphql.Boolean,
		Resolve: resolve.Logout,
	}
}

func rootMutation() *graphql.Object {
	metaEntity := model.GlobalModel.Graph.GetMetaEntity()
	mutationFields := graphql.Fields{
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

	if config.Storage() != "" {
		mutationFields[consts.SINGLE_UPLOAD] = &graphql.Field{
			Type: Cache.OutputType(consts.MEDIA_ENTITY_NAME),
			Args: graphql.FieldConfigArgument{
				consts.ARG_FILE: &graphql.ArgumentConfig{
					Type: &graphql.NonNull{OfType: scalars.UploadType},
				},
			},
			Resolve: resolve.SingleUploadResolve,
		}
	}

	for _, entity := range model.GlobalModel.Graph.RootEnities() {
		if entity.Domain.Root {
			appendEntityMutationToFields(entity, mutationFields)
		}
	}

	for _, partial := range model.GlobalModel.Graph.RootPartails() {
		if partial.Domain.Root {
			appendPartialMutationToFields(partial, mutationFields)
		}
	}

	// for _, service := range model.GlobalModel.Graph.RootExternals() {
	// 	appendServiceMutationFields(service, mutationFields)
	// }

	if config.AuthUrl() == "" {
		appendAuthMutation(mutationFields)
	}

	rootMutation := graphql.ObjectConfig{
		Name:        consts.ROOT_MUTATION_NAME,
		Fields:      mutationFields,
		Description: "Root mutation of entity engine.",
	}

	return graphql.NewObject(rootMutation)
}

func deleteArgs(entity *graph.Entity) graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		consts.ARG_WHERE: &graphql.ArgumentConfig{
			Type: Cache.WhereExp(entity.Name()),
		},
	}
}

func deleteByIdArgs() graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		consts.ID: &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	}
}

func upsertArgs(entity *graph.Entity) graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		consts.ARG_OBJECTS: &graphql.ArgumentConfig{
			Type: &graphql.NonNull{
				OfType: &graphql.List{
					OfType: &graphql.NonNull{
						OfType: Cache.SaveInput(entity.Name()),
					},
				},
			},
		},
	}
}

func upsertOneArgs(entity *graph.Entity) graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		consts.ARG_OBJECT: &graphql.ArgumentConfig{
			Type: &graphql.NonNull{
				OfType: Cache.SaveInput(entity.Name()),
			},
		},
	}
}

func setArgs(entity *graph.Entity) graphql.FieldConfigArgument {
	updateInput := Cache.SetInput(entity.Name())
	return graphql.FieldConfigArgument{
		consts.ARG_SET: &graphql.ArgumentConfig{
			Type: &graphql.NonNull{
				OfType: updateInput,
			},
		},
		consts.ARG_WHERE: &graphql.ArgumentConfig{
			Type: Cache.WhereExp(entity.Name()),
		},
	}
}

func appendEntityMutationToFields(entity *graph.Entity, feilds graphql.Fields) {
	(feilds)[entity.DeleteName()] = &graphql.Field{
		Type: Cache.MutationResponse(entity.Name()),
		Args: deleteArgs(entity),
		//Resolve: entity.QueryResolve(),
	}
	(feilds)[entity.DeleteByIdName()] = &graphql.Field{
		Type: Cache.OutputType(entity.Name()),
		Args: deleteByIdArgs(),
		//Resolve: entity.QueryResolve(),
	}
	(feilds)[entity.UpsertName()] = &graphql.Field{
		Type: &graphql.List{OfType: Cache.OutputType(entity.Name())},
		Args: upsertArgs(entity),
	}
	//Resolve: entity.QueryResolve(),
	(feilds)[entity.UpsertOneName()] = &graphql.Field{
		Type:    Cache.OutputType(entity.Name()),
		Args:    upsertOneArgs(entity),
		Resolve: resolve.PostOneResolveFn(entity),
	}

	updateInput := Cache.SetInput(entity.Name())
	if len(updateInput.Fields()) > 0 {
		(feilds)[entity.SetName()] = &graphql.Field{
			Type: Cache.MutationResponse(entity.Name()),
			Args: setArgs(entity),
			//Resolve: entity.QueryResolve(),
		}
	}

}

func appendPartialMutationToFields(partial *graph.Partial, feilds graphql.Fields) {

	(feilds)[partial.DeleteName()] = &graphql.Field{
		Type: Cache.MutationResponse(partial.Name()),
		Args: deleteArgs(&partial.Entity),
		//Resolve: entity.QueryResolve(),
	}
	(feilds)[partial.DeleteByIdName()] = &graphql.Field{
		Type: Cache.OutputType(partial.Name()),
		Args: deleteByIdArgs(),
		//Resolve: entity.QueryResolve(),
	}
	(feilds)[partial.InsertName()] = &graphql.Field{
		Type: &graphql.List{OfType: Cache.OutputType(partial.Name())},
		Args: upsertArgs(&partial.Entity),
	}
	//Resolve: entity.QueryResolve(),
	(feilds)[partial.InsertOneName()] = &graphql.Field{
		Type: Cache.OutputType(partial.Name()),
		Args: upsertOneArgs(&partial.Entity),
		//Resolve: resolve.PostOneResolveFn(&partial.Entity),
	}
	(feilds)[partial.UpdateName()] = &graphql.Field{
		Type: &graphql.List{OfType: Cache.OutputType(partial.Name())},
		Args: upsertArgs(&partial.Entity),
	}
	//Resolve: entity.QueryResolve(),
	(feilds)[partial.UpdateOneName()] = &graphql.Field{
		Type: Cache.OutputType(partial.Name()),
		Args: upsertOneArgs(&partial.Entity),
		//Resolve: resolve.PostOneResolveFn(&partial.Entity),
	}

	updateInput := Cache.SetInput(partial.Name())
	if len(updateInput.Fields()) > 0 {
		(feilds)[partial.SetName()] = &graphql.Field{
			Type: Cache.MutationResponse(partial.Name()),
			Args: setArgs(&partial.Entity),
			//Resolve: entity.QueryResolve(),
		}
	}

}
