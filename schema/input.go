package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/model"
	"rxdrag.com/entity-engine/utils"
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

		associas := entity.Associations

		for i := range associas {
			assoc := associas[i]
			typeInput := c.SaveInput(assoc.TypeEntity.Name)
			if len(typeInput.Fields()) == 0 {
				continue
			}
			arrayType := c.makeAssociationType(assoc)
			input.AddFieldConfig(assoc.Name, &graphql.InputObjectFieldConfig{
				Type:        arrayType,
				Description: assoc.Description,
			})
			update.AddFieldConfig(assoc.Name, &graphql.InputObjectFieldConfig{
				Type:        arrayType,
				Description: assoc.Description,
			})
		}
	}
}

func (c *TypeCache) makeAssociationType(association *model.Association) *graphql.InputObject {
	typeInput := c.SaveInput(association.TypeEntity.Name)
	listType := &graphql.List{
		OfType: typeInput,
	}
	if association.IsArray() {
		return graphql.NewInputObject(graphql.InputObjectConfig{
			Name: association.OfEntity.Name + utils.FirstUpper(association.Name) + "Input",
			Fields: graphql.InputObjectConfigFieldMap{
				consts.ARG_ADD: &graphql.InputObjectFieldConfig{
					Type: listType,
				},
				consts.ARG_DELETE: &graphql.InputObjectFieldConfig{
					Type: listType,
				},
				consts.ARG_MODIFY: &graphql.InputObjectFieldConfig{
					Type: listType,
				},
				consts.ARG_SYNC: &graphql.InputObjectFieldConfig{
					Type: listType,
				},
			},
		})
	} else {
		return graphql.NewInputObject(graphql.InputObjectConfig{
			Name: association.OfEntity.Name + utils.FirstUpper(association.Name) + "Input",
			Fields: graphql.InputObjectConfigFieldMap{
				consts.ARG_DELETE: &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				consts.ARG_SYNC: &graphql.InputObjectFieldConfig{
					Type: typeInput,
				},
			},
		})
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
							OfType: Cache.OutputType(entity.Name),
						},
					},
				},
			},
		},
	)

	return &returnValue
}
