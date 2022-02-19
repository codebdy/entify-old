package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/utils"
)

func (entity *EntityMeta) AppendToMutationFields(feilds *graphql.Fields) {
	//如果是枚举
	if entity.EntityType == Entity_ENUM {
		return
	}

	name := utils.FirstUpper(entity.Name)

	(*feilds)["delete"+name] = &graphql.Field{
		Type: entity.toOutputType(),
		Args: graphql.FieldConfigArgument{
			"where": &graphql.ArgumentConfig{
				Type: entity.toWhereExp(),
			},
		},
		//Resolve: entity.QueryResolve(),
	}
	(*feilds)["delete"+name+"ById"] = &graphql.Field{
		Type: entity.toOutputType(),
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		//Resolve: entity.QueryResolve(),
	}
	(*feilds)["insert"+name] = &graphql.Field{
		Type: entity.toOutputType(),
		Args: graphql.FieldConfigArgument{
			"objects": &graphql.ArgumentConfig{
				Type: &graphql.NonNull{
					OfType: &graphql.List{
						OfType: &graphql.NonNull{
							OfType: entity.toInsertInput(),
						},
					},
				},
			},
		},
	}
	//Resolve: entity.QueryResolve(),
	(*feilds)["insertOne"+name] = &graphql.Field{
		Type: entity.toOutputType(),
		Args: graphql.FieldConfigArgument{
			"object": &graphql.ArgumentConfig{
				Type: &graphql.NonNull{
					OfType: entity.toInsertInput(),
				},
			},
		},
		//Resolve: entity.QueryResolve(),
	}
	(*feilds)["update"+name] = &graphql.Field{
		Type: entity.toOutputType(),
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		//Resolve: entity.QueryResolve(),
	}

	(*feilds)["update"+name+"ById"] = &graphql.Field{
		Type: entity.toOutputType(),
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		//Resolve: entity.QueryResolve(),
	}
}
