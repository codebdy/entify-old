package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/model/graph"
)

func (c *TypeCache) makeInputs() {
	for i := range Model.Graph.Entities {
		entity := Model.Graph.Entities[i]
		c.UpdateInputMap[entity.Name()] = makeUpdateInput(entity)
		c.SaveInputMap[entity.Name()] = makeSaveInput(entity)
		c.MutationResponseMap[entity.Name()] = makeMutationResponseType(entity)
	}
	for i := range Model.Graph.Entities {
		entity := Model.Graph.Entities[i]
		c.HasManyInputMap[entity.Name()] = c.makeHasManyInput(entity)
		c.HasOneInputMap[entity.Name()] = c.makeHasOneInput(entity)
	}
	c.makeInputRelations()
}

func (c *TypeCache) makeHasManyInput(entity *graph.Entity) *graphql.InputObject {
	typeInput := c.SaveInput(entity.Name())
	listType := &graphql.List{
		OfType: typeInput,
	}
	return graphql.NewInputObject(graphql.InputObjectConfig{
		Name: entity.GetHasManyName(),
		Fields: graphql.InputObjectConfigFieldMap{
			consts.ARG_ADD: &graphql.InputObjectFieldConfig{
				Type: listType,
			},
			consts.ARG_DELETE: &graphql.InputObjectFieldConfig{
				Type: listType,
			},
			consts.ARG_UPDATE: &graphql.InputObjectFieldConfig{
				Type: listType,
			},
			consts.ARG_SYNC: &graphql.InputObjectFieldConfig{
				Type: listType,
			},
			consts.ARG_CASCADE: &graphql.InputObjectFieldConfig{
				Type: graphql.Boolean,
			},
		},
	})
}

func (c *TypeCache) makeHasOneInput(entity *graph.Entity) *graphql.InputObject {
	typeInput := c.SaveInput(entity.Name())
	return graphql.NewInputObject(graphql.InputObjectConfig{
		Name: entity.GetHasOneName(),
		Fields: graphql.InputObjectConfigFieldMap{
			consts.ARG_DELETE: &graphql.InputObjectFieldConfig{
				Type: graphql.Boolean,
			},
			consts.ARG_SYNC: &graphql.InputObjectFieldConfig{
				Type: typeInput,
			},
			consts.ARG_CASCADE: &graphql.InputObjectFieldConfig{
				Type: graphql.Boolean,
			},
		},
	})
}

func (c *TypeCache) makeInputRelations() {
	for i := range Model.Graph.Entities {
		entity := Model.Graph.Entities[i]

		input := c.UpdateInputMap[entity.Name()]
		update := c.SaveInputMap[entity.Name()]

		associas := entity.Associations()

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

func (c *TypeCache) makeAssociationType(association *graph.Association) *graphql.InputObject {
	if association.IsArray() {
		return c.HasManyInput(association.TypeClass().Name())
	} else {
		return c.HasOneInput(association.TypeClass().Name())
	}
}

func inputFields(entity *graph.Entity, isPost bool) graphql.InputObjectConfigFieldMap {
	fields := graphql.InputObjectConfigFieldMap{}
	for _, column := range entity.Attributes() {
		if column.Name != consts.ID || isPost {
			fields[column.Name] = &graphql.InputObjectFieldConfig{
				Type:        AttributeType(column),
				Description: column.Description,
			}
		}
	}
	return fields
}

func makeSaveInput(entity *graph.Entity) *graphql.InputObject {
	name := entity.Name() + consts.INPUT

	return graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name:   name,
			Fields: inputFields(entity, true),
		},
	)
}

func makeUpdateInput(entity *graph.Entity) *graphql.InputObject {
	return graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name:   entity.Name() + consts.UPDATE_INPUT,
			Fields: inputFields(entity, false),
		},
	)
}

func makeMutationResponseType(entity *graph.Entity) *graphql.Output {
	var returnValue graphql.Output

	returnValue = graphql.NewObject(
		graphql.ObjectConfig{
			Name: entity.Name() + consts.MUTATION_RESPONSE,
			Fields: graphql.Fields{
				consts.RESPONSE_AFFECTEDROWS: &graphql.Field{
					Type: graphql.Int,
				},
				consts.RESPONSE_RETURNING: &graphql.Field{
					Type: &graphql.NonNull{
						OfType: &graphql.List{
							OfType: Cache.OutputType(entity.Name()),
						},
					},
				},
			},
		},
	)

	return &returnValue
}
