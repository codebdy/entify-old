package schema

import (
	"github.com/graphql-go/graphql"
)

func (entity *EntityMeta) AppendToMutationFields(feilds *graphql.Fields) {
	//如果是枚举
	if entity.EntityType == Entity_ENUM {
		return
	}

	(*feilds)["delete"+entity.Name] = &graphql.Field{
		Type: entity.toOutputType(),
		Args: graphql.FieldConfigArgument{
			"where": &graphql.ArgumentConfig{
				Type: entity.toWhereExp(),
			},
		},
		//Resolve: entity.QueryResolve(),
	}
	(*feilds)["delete"+entity.Name+"ByPK"] = &graphql.Field{
		Type: entity.toOutputType(),
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		//Resolve: entity.QueryResolve(),
	}
}
