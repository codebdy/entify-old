package schema

import (
	"github.com/graphql-go/graphql"
)

//类型缓存， query mutaion通用
var UpdateInputMap = make(map[string]*graphql.Input)
var PostInputMap = make(map[string]*graphql.Input)

//mutition类型缓存， mutaion用
var mutationResponseMap = make(map[string]*graphql.Output)

func (entity *EntityMeta) createInputFields(isPost bool) graphql.InputObjectConfigFieldMap {
	fields := graphql.InputObjectConfigFieldMap{}
	for _, column := range entity.Columns {
		if column.Name != "id" || isPost {
			fields[column.Name] = &graphql.InputObjectFieldConfig{
				Type: column.toInputType(),
			}
		}
	}

	return fields
}

func (entity *EntityMeta) toUpdateInput() *graphql.Input {
	if UpdateInputMap[entity.Name] != nil {
		return UpdateInputMap[entity.Name]
	}
	var returnValue graphql.Input

	returnValue = graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name:   entity.Name + "UpdateInput",
			Fields: entity.createInputFields(false),
		},
	)
	UpdateInputMap[entity.Name] = &returnValue
	return &returnValue
}

func (entity *EntityMeta) toPostInput() *graphql.Input {
	if PostInputMap[entity.Name] != nil {
		return PostInputMap[entity.Name]
	}
	var returnValue graphql.Input

	returnValue = graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name:   entity.Name + "PostInput",
			Fields: entity.createInputFields(true),
		},
	)
	PostInputMap[entity.Name] = &returnValue
	return &returnValue
}

func (entity *EntityMeta) toMutationResponseType() graphql.Output {
	if mutationResponseMap[entity.Name] != nil {
		return *mutationResponseMap[entity.Name]
	}
	var returnValue graphql.Output

	returnValue = graphql.NewObject(
		graphql.ObjectConfig{
			Name: entity.Name + "MutationResponse",
			Fields: graphql.Fields{
				"affectedRows": &graphql.Field{
					Type: graphql.Int,
				},
				"returning": &graphql.Field{
					Type: &graphql.NonNull{
						OfType: &graphql.List{
							OfType: entity.toOutputType(),
						},
					},
				},
			},
		},
	)

	mutationResponseMap[entity.Name] = &returnValue
	return returnValue
}
