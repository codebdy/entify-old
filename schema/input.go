package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/meta"
)

//mutition类型缓存， mutaion用
var mutationResponseMap = make(map[string]*graphql.Output)

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
	if UpdateInputMap[entity.Name] != nil {
		return UpdateInputMap[entity.Name]
	}
	var returnValue graphql.Input

	returnValue = graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name:   entity.Name + "UpdateInput",
			Fields: InputFields(entity, false),
		},
	)
	UpdateInputMap[entity.Name] = &returnValue
	return &returnValue
}

func PostInput(entity *meta.Entity) *graphql.Input {
	if PostInputMap[entity.Name] != nil {
		return PostInputMap[entity.Name]
	}
	var returnValue graphql.Input

	returnValue = graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name:   entity.Name + "PostInput",
			Fields: InputFields(entity, true),
		},
	)
	PostInputMap[entity.Name] = &returnValue
	return &returnValue
}

func MutationResponseType(entity *meta.Entity) graphql.Output {
	if mutationResponseMap[entity.Name] != nil {
		return *mutationResponseMap[entity.Name]
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
							OfType: OutputType(entity),
						},
					},
				},
			},
		},
	)

	mutationResponseMap[entity.Name] = &returnValue
	return returnValue
}
