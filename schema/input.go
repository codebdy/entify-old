package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/meta"
)

func InputFields(entity *meta.Entity, parents []*meta.Entity, isPost bool) graphql.InputObjectConfigFieldMap {
	fields := graphql.InputObjectConfigFieldMap{}
	for _, column := range entity.Columns {
		if column.Name != "id" || isPost {
			fields[column.Name] = &graphql.InputObjectFieldConfig{
				Type: ColumnType(&column),
			}
		}
	}
	relations := meta.Metas.EntityRelations(entity)
	newParents := append(parents, entity)
	for i := range relations {
		relation := relations[i]
		if !findParent(relation.TypeEntity.Uuid, newParents) {
			typeInput := PostInput(relation.TypeEntity, newParents)
			if relation.IsArray() {
				fields[relation.Name] = &graphql.InputObjectFieldConfig{
					Type: &graphql.List{
						OfType: *typeInput,
					},
				}
			} else {
				fields[relation.Name] = &graphql.InputObjectFieldConfig{
					Type: *typeInput,
				}
			}
		}
	}
	return fields
}

func UpdateInput(entity *meta.Entity, parents []*meta.Entity) *graphql.Input {
	if Cache.UpdateInputMap[entity.Name] != nil {
		return Cache.UpdateInputMap[entity.Name]
	}
	var returnValue graphql.Input

	returnValue = graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name:   entity.Name + "UpdateInput",
			Fields: InputFields(entity, parents, false),
		},
	)
	Cache.UpdateInputMap[entity.Name] = &returnValue
	return &returnValue
}

func PostInput(entity *meta.Entity, parents []*meta.Entity) *graphql.Input {
	if Cache.PostInputMap[entity.Name] != nil {
		return Cache.PostInputMap[entity.Name]
	}
	var returnValue graphql.Input

	returnValue = graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name:   entity.Name + "PostInput",
			Fields: InputFields(entity, parents, true),
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
							OfType: *OutputType(entity, []*meta.Entity{}),
						},
					},
				},
			},
		},
	)

	Cache.MutationResponseMap[entity.Name] = &returnValue
	return &returnValue
}
