package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/meta"
)

//mutition类型缓存， mutaion用

func InputFields(entity *meta.Entity, isPost bool) graphql.InputObjectConfigFieldMap {
	fields := graphql.InputObjectConfigFieldMap{}
	for _, column := range entity.Columns {
		if column.Name != "id" || isPost {
			fields[column.Name] = &graphql.InputObjectFieldConfig{
				Type: ColumnType(&column),
			}
		}
	}

	return fields
}

func UpdateInput(entity *meta.Entity) *graphql.Input {
	if Cache.UpdateInputMap[entity.Name] != nil {
		return Cache.UpdateInputMap[entity.Name]
	}
	var returnValue graphql.Input

	returnValue = graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name:   entity.Name + "UpdateInput",
			Fields: InputFields(entity, false),
		},
	)
	Cache.UpdateInputMap[entity.Name] = &returnValue
	return &returnValue
}

func PostInput(entity *meta.Entity) *graphql.Input {
	if Cache.PostInputMap[entity.Name] != nil {
		return Cache.PostInputMap[entity.Name]
	}
	var returnValue graphql.Input

	returnValue = graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name:   entity.Name + "PostInput",
			Fields: InputFields(entity, true),
		},
	)
	Cache.PostInputMap[entity.Name] = &returnValue
	return &returnValue
}

func MutationResponseType(entity *meta.Entity) *graphql.Output {
	if Cache.MutationResponseMap[entity.Name] != nil {
		return Cache.MutationResponseMap[entity.Name]
	}
	var returnValue graphql.Output

	returnValue = graphql.NewObject(
		graphql.ObjectConfig{
			Name: entity.Name + "MutationResponse",
			Fields: graphql.Fields{
				consts.RESPONSE_AFFECTEDROWS: &graphql.Field{
					Type: graphql.Int,
				},
				consts.RESPONSE_RETURNING: &graphql.Field{
					Type: &graphql.NonNull{
						OfType: &graphql.List{
							OfType: *OutputType(entity),
						},
					},
				},
			},
		},
	)

	Cache.MutationResponseMap[entity.Name] = &returnValue
	return &returnValue
}
