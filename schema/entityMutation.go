package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/meta"
	"rxdrag.com/entity-engine/repository"
	"rxdrag.com/entity-engine/utils"
)

func AppendToMutationFields(entity *meta.EntityMeta, feilds *graphql.Fields) {
	//如果是枚举
	if entity.EntityType == meta.Entity_ENUM {
		return
	}

	name := utils.FirstUpper(entity.Name)

	(*feilds)["delete"+name] = &graphql.Field{
		Type: MutationResponseType(entity),
		Args: graphql.FieldConfigArgument{
			"where": &graphql.ArgumentConfig{
				Type: WhereExp(entity),
			},
		},
		//Resolve: entity.QueryResolve(),
	}
	(*feilds)["delete"+name+"ById"] = &graphql.Field{
		Type: OutputType(entity),
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		//Resolve: entity.QueryResolve(),
	}
	(*feilds)["post"+name] = &graphql.Field{
		Type: OutputType(entity),
		Args: graphql.FieldConfigArgument{
			"objects": &graphql.ArgumentConfig{
				Type: &graphql.NonNull{
					OfType: &graphql.List{
						OfType: &graphql.NonNull{
							OfType: *PostInput(entity),
						},
					},
				},
			},
		},
	}
	//Resolve: entity.QueryResolve(),
	(*feilds)["postOne"+name] = &graphql.Field{
		Type: MutationResponseType(entity),
		Args: graphql.FieldConfigArgument{
			"object": &graphql.ArgumentConfig{
				Type: &graphql.NonNull{
					OfType: *PostInput(entity),
				},
			},
		},
		Resolve: repository.PostOneResolveFn(entity),
	}

	(*feilds)["update"+name] = &graphql.Field{
		Type: MutationResponseType(entity),
		Args: graphql.FieldConfigArgument{
			"objects": &graphql.ArgumentConfig{
				Type: &graphql.NonNull{
					OfType: &graphql.List{
						OfType: &graphql.NonNull{
							OfType: *UpdateInput(entity),
						},
					},
				},
			},
			"where": &graphql.ArgumentConfig{
				Type: WhereExp(entity),
			},
		},
		//Resolve: entity.QueryResolve(),
	}
}
