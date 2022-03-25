package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/model"
)

func (c *TypeCache) makeInputs() {
	for i := range model.TheModel.Entities {
		entity := model.TheModel.Entities[i]
		c.UpdateInputMap[entity.Name] = makeUpdateInput(entity)
		c.SaveInputMap[entity.Name] = makeSaveInput(entity)
		c.MutationResponseMap[entity.Name] = makeMutationResponseType(entity)

	}

	c.makeInputRelations()
}

func (c *TypeCache) makeInputRelations() {
	for i := range model.TheModel.Entities {
		entity := model.TheModel.Entities[i]

		input := c.UpdateInputMap[entity.Name]
		update := c.SaveInputMap[entity.Name]

		assocs := entity.Associations

		for i := range assocs {
			relation := assocs[i]
			typeInput := c.SaveInput(relation.TypeEntity)
			if relation.IsArray() {
				input.AddFieldConfig(relation.Name, &graphql.InputObjectFieldConfig{
					Type: &graphql.List{
						OfType: typeInput,
					},
					Description: relation.Description,
				})
				update.AddFieldConfig(relation.Name, &graphql.InputObjectFieldConfig{
					Type: &graphql.List{
						OfType: typeInput,
					},
					Description: relation.Description,
				})
			} else {
				input.AddFieldConfig(relation.Name, &graphql.InputObjectFieldConfig{
					Type:        typeInput,
					Description: relation.Description,
				})
				update.AddFieldConfig(relation.Name, &graphql.InputObjectFieldConfig{
					Type:        typeInput,
					Description: relation.Description,
				})
			}
		}

	}
}

func inputFields(entity *model.Entity, isPost bool) graphql.InputObjectConfigFieldMap {
	fields := graphql.InputObjectConfigFieldMap{}
	for _, column := range entity.Columns {
		if column.Name != consts.ID || isPost {
			fields[column.Name] = &graphql.InputObjectFieldConfig{
				Type:        ColumnType(column),
				Description: column.Description,
			}
		}
	}
	return fields
}

func makeSaveInput(entity *model.Entity) *graphql.InputObject {
	name := entity.Name + consts.INPUT

	return graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name:   name,
			Fields: inputFields(entity, true),
		},
	)
}

func makeUpdateInput(entity *model.Entity) *graphql.InputObject {
	return graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name:   entity.Name + consts.UPDATE_INPUT,
			Fields: inputFields(entity, false),
		},
	)
}

func makeMutationResponseType(entity *model.Entity) *graphql.Output {
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
							OfType: Cache.OutputObjectType(entity),
						},
					},
				},
			},
		},
	)

	return &returnValue
}
