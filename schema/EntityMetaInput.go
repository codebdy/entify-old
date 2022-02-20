package schema

import (
	"github.com/graphql-go/graphql"
)

//类型缓存， query mutaion通用
var UpdateInputMap = make(map[string]*graphql.Input)
var PostInputMap = make(map[string]*graphql.Input)

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
