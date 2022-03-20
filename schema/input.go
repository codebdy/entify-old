package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/meta"
)

func inputFields(entity *meta.Entity, isPost bool) graphql.InputObjectConfigFieldMap {
	fields := graphql.InputObjectConfigFieldMap{}
	for _, column := range meta.Metas.EntityAllColumns(entity) {
		if column.Name != "id" || isPost {
			fields[column.Name] = &graphql.InputObjectFieldConfig{
				Type: ColumnType(&column),
			}
		}
	}
	// relations := meta.Metas.EntityAllRelations(entity)
	// newParents := append(parents, entity)

	// for i := range relations {
	// 	relation := relations[i]
	// 	typeInput := PostInput(relation.TypeEntity, newParents)
	// 	if relation.IsArray() {
	// 		fields[relation.Name] = &graphql.InputObjectFieldConfig{
	// 			Type: &graphql.List{
	// 				OfType: *typeInput,
	// 			},
	// 		}
	// 	} else {
	// 		fields[relation.Name] = &graphql.InputObjectFieldConfig{
	// 			Type: *typeInput,
	// 		}
	// 	}
	// }
	return fields
}
func (c *TypeCache) makeInputs() {
	for i := range meta.Metas.Entities {
		entity := meta.Metas.Entities[i]
		if entity.EntityType != meta.ENTITY_ENUM {
			c.UpdateInputMap[entity.Name] = makeUpdateInput(&entity)
			c.SaveInputMap[entity.Name] = makeSaveInput(&entity)
			c.MutationResponseMap[entity.Name] = makeMutationResponseType(&entity)
		}
	}
}

func makeSaveInput(entity *meta.Entity) *graphql.Input {
	name := entity.Name + consts.INPUT
	var returnValue graphql.Input

	returnValue = graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name:   name,
			Fields: inputFields(entity, true),
		},
	)
	return &returnValue
}

func makeUpdateInput(entity *meta.Entity) *graphql.Input {
	var returnValue graphql.Input

	returnValue = graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name:   entity.Name + consts.UPDATE_INPUT,
			Fields: inputFields(entity, false),
		},
	)
	return &returnValue
}

func makeMutationResponseType(entity *meta.Entity) *graphql.Output {
	var returnValue graphql.Output

	returnValue = graphql.NewObject(
		graphql.ObjectConfig{
			Name: entity.Name + consts.MUTATION_RESPONSE,
			Fields: graphql.Fields{
				consts.RESPONSE_AFFECTEDROWS: &graphql.Field{
					Type: graphql.Int,
				},
				consts.RESPONSE_RETURNING: &graphql.Field{
					Type: &graphql.NonNull{
						OfType: &graphql.List{
							OfType: Cache.OutputType(entity),
						},
					},
				},
			},
		},
	)

	return &returnValue
}
