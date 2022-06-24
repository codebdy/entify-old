package schema

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"rxdrag.com/entify/consts"
	"rxdrag.com/entify/model"
	"rxdrag.com/entify/model/graph"
)

func (c *TypeCache) makeInputs() {
	for i := range model.GlobalModel.Graph.Entities {
		entity := model.GlobalModel.Graph.Entities[i]
		c.SetInputMap[entity.Name()] = makeEntitySetInput(entity)
		c.SaveInputMap[entity.Name()] = makeEntitySaveInput(entity)
		c.MutationResponseMap[entity.Name()] = makeEntityMutationResponseType(entity)
	}

	for i := range model.GlobalModel.Graph.Entities {
		entity := model.GlobalModel.Graph.Entities[i]
		c.HasManyInputMap[entity.Name()] = c.makeHasManyInput(entity)
		c.HasOneInputMap[entity.Name()] = c.makeHasOneInput(entity)
	}

	for i := range model.GlobalModel.Graph.Partials {
		partail := model.GlobalModel.Graph.Partials[i]
		c.SetInputMap[partail.Name()] = makePartialSetInput(partail)
		c.SaveInputMap[partail.Name()] = makePartailSaveInput(partail)
		c.MutationResponseMap[partail.Name()] = makePartialMutationResponseType(partail)
	}

	for i := range model.GlobalModel.Graph.Partials {
		partial := model.GlobalModel.Graph.Partials[i]
		c.HasManyInputMap[partial.Name()] = c.makeHasManyInput(&partial.Entity)
		c.HasOneInputMap[partial.Name()] = c.makeHasOneInput(&partial.Entity)
	}

	for i := range model.GlobalModel.Graph.Externals {
		external := model.GlobalModel.Graph.Externals[i]
		c.HasManyInputMap[external.Name()] = c.makeHasManyInput(&external.Entity)
		c.HasOneInputMap[external.Name()] = c.makeHasOneInput(&external.Entity)
	}
	c.makeEntityInputRelations()
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

func (c *TypeCache) makeEntityInputRelations() {
	for i := range model.GlobalModel.Graph.Entities {
		entity := model.GlobalModel.Graph.Entities[i]

		input := c.SetInputMap[entity.Name()]
		update := c.SaveInputMap[entity.Name()]

		associas := entity.AllAssociations()

		for i := range associas {
			assoc := associas[i]
			if !assoc.IsAbstract() {
				typeInput := c.SaveInput(assoc.Owner().Name())
				if typeInput == nil {
					panic("can not find save input:" + assoc.Owner().Name())
				}
				if len(typeInput.Fields()) == 0 {
					fmt.Println("Fields == 0")
					continue
				}

				arrayType := c.getAssociationType(assoc)
				//如果是虚类，并且没有子类
				if assoc.TypeInterface() != nil && len(assoc.TypeInterface().Children) == 0 {
					continue
				}
				if arrayType == nil {
					panic("Can not get association type:" + assoc.Owner().Name() + "." + assoc.Name())
				}
				input.AddFieldConfig(assoc.Name(), &graphql.InputObjectFieldConfig{
					Type:        arrayType,
					Description: assoc.Description(),
				})
				update.AddFieldConfig(assoc.Name(), &graphql.InputObjectFieldConfig{
					Type:        arrayType,
					Description: assoc.Description(),
				})
			} else {
				c.makeDevrivedInputRelations(assoc, entity, input, update)
			}
		}
	}
}

func (c *TypeCache) makeDevrivedInputRelations(association *graph.Association,
	entity *graph.Entity,
	input *graphql.InputObject,
	update *graphql.InputObject,
) {
	derivedAssociations := association.DerivedAssociationsByOwnerUuid(entity.Uuid())
	for i := range derivedAssociations {
		derivedAssociation := derivedAssociations[i]
		arrayType := c.getDerivedAssociationType(derivedAssociation)
		if arrayType == nil {
			panic("Can not get derived association type:" + derivedAssociation.OwnerClassUuid + "." + derivedAssociation.DerivedFrom.Name())
		}
		input.AddFieldConfig(derivedAssociation.Name(), &graphql.InputObjectFieldConfig{
			Type:        arrayType,
			Description: association.Description(),
		})
		update.AddFieldConfig(derivedAssociation.Name(), &graphql.InputObjectFieldConfig{
			Type:        arrayType,
			Description: association.Description(),
		})
	}
}

func (c *TypeCache) getAssociationType(association *graph.Association) *graphql.InputObject {
	if association.IsArray() {
		return c.HasManyInput(association.TypeClass().Name())
	} else {
		return c.HasOneInput(association.TypeClass().Name())
	}
}

func (c *TypeCache) getDerivedAssociationType(association *graph.DerivedAssociation) *graphql.InputObject {
	if association.DerivedFrom.IsArray() {
		return c.HasManyInput(association.TypeEntity().Name())
	} else {
		return c.HasOneInput(association.TypeEntity().Name())
	}
}

func inputFields(entity *graph.Entity, isPost bool) graphql.InputObjectConfigFieldMap {
	fields := graphql.InputObjectConfigFieldMap{}
	for _, column := range entity.AllAttributes() {
		if column.Name != consts.ID || isPost {
			fields[column.Name] = &graphql.InputObjectFieldConfig{
				Type:        AttributeType(column),
				Description: column.Description,
			}
		}
	}
	return fields
}

func makeEntitySaveInput(entity *graph.Entity) *graphql.InputObject {
	name := entity.Name() + consts.INPUT
	return graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name:   name,
			Fields: inputFields(entity, true),
		},
	)
}

func makePartailSaveInput(partail *graph.Partial) *graphql.InputObject {
	name := partail.NameWithPartial() + consts.INPUT
	return graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name:   name,
			Fields: inputFields(&partail.Entity, true),
		},
	)
}

func makeEntitySetInput(entity *graph.Entity) *graphql.InputObject {
	return graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name:   entity.Name() + consts.SET_INPUT,
			Fields: inputFields(entity, false),
		},
	)
}

func makePartialSetInput(partial *graph.Partial) *graphql.InputObject {
	return graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name:   partial.NameWithPartial() + consts.SET_INPUT,
			Fields: inputFields(&partial.Entity, false),
		},
	)
}

func makeEntityMutationResponseType(entity *graph.Entity) *graphql.Object {
	var returnValue *graphql.Object

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

	return returnValue
}

func makePartialMutationResponseType(partial *graph.Partial) *graphql.Object {
	var returnValue *graphql.Object

	returnValue = graphql.NewObject(
		graphql.ObjectConfig{
			Name: partial.NameWithPartial() + consts.MUTATION_RESPONSE,
			Fields: graphql.Fields{
				consts.RESPONSE_AFFECTEDROWS: &graphql.Field{
					Type: graphql.Int,
				},
				consts.RESPONSE_RETURNING: &graphql.Field{
					Type: &graphql.NonNull{
						OfType: &graphql.List{
							OfType: Cache.OutputType(partial.Name()),
						},
					},
				},
			},
		},
	)

	return returnValue
}
