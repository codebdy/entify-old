package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/meta"
	"rxdrag.com/entity-engine/resolve"
	"rxdrag.com/entity-engine/utils"
)

func AppendToMutationFields(entity *meta.Entity, feilds *graphql.Fields) {
	//如果是枚举
	if entity.EntityType == meta.ENTITY_ENUM {
		return
	}

	name := utils.FirstUpper(entity.Name)

	(*feilds)["delete"+name] = &graphql.Field{
		Type: *MutationResponseType(entity),
		Args: graphql.FieldConfigArgument{
			"where": &graphql.ArgumentConfig{
				Type: WhereExp(entity),
			},
		},
		//Resolve: entity.QueryResolve(),
	}
	(*feilds)["delete"+name+"ById"] = &graphql.Field{
		Type: *OutputType(entity, []*meta.Entity{}),
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		//Resolve: entity.QueryResolve(),
	}
	(*feilds)["post"+name] = &graphql.Field{
		Type: *OutputType(entity, []*meta.Entity{}),
		Args: graphql.FieldConfigArgument{
			consts.ARG_OBJECTS: &graphql.ArgumentConfig{
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
		Type: *OutputType(entity, []*meta.Entity{}),
		Args: graphql.FieldConfigArgument{
			consts.ARG_OBJECT: &graphql.ArgumentConfig{
				Type: &graphql.NonNull{
					OfType: *PostInput(entity),
				},
			},
		},
		Resolve: resolve.PostOneResolveFn(entity),
	}

	(*feilds)["update"+name] = &graphql.Field{
		Type: *MutationResponseType(entity),
		Args: graphql.FieldConfigArgument{
			consts.ARG_OBJECTS: &graphql.ArgumentConfig{
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
