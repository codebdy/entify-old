package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/meta"
)

func (c *TypeCache) makeInputs() {
	for i := range meta.Metas.Entities {
		entity := &meta.Metas.Entities[i]
		if entity.EntityType != meta.ENTITY_ENUM {
			c.UpdateInputMap[entity.Name] = makeUpdateInput(entity)
			c.SaveInputMap[entity.Name] = makeSaveInput(entity)
			c.MutationResponseMap[entity.Name] = makeMutationResponseType(entity)
		}
	}

	c.makeInputRelations()
}

func (c *TypeCache) makeInputRelations() {
	for i := range meta.Metas.Entities {
		entity := &meta.Metas.Entities[i]
		if entity.EntityType != meta.ENTITY_ENUM {
			input := c.UpdateInputMap[entity.Name]
			update := c.SaveInputMap[entity.Name]

			relations := meta.Metas.EntityAllRelations(entity)

			for i := range relations {
				relation := relations[i]
				typeInput := c.SaveInput(relation.TypeEntity)
				if relation.IsArray() {
					input.AddFieldConfig(relation.Name, &graphql.InputObjectFieldConfig{
						Type: &graphql.List{
							OfType: typeInput,
						},
					})
					update.AddFieldConfig(relation.Name, &graphql.InputObjectFieldConfig{
						Type: &graphql.List{
							OfType: typeInput,
						},
					})
				} else {
					input.AddFieldConfig(relation.Name, &graphql.InputObjectFieldConfig{
						Type: typeInput,
					})
					update.AddFieldConfig(relation.Name, &graphql.InputObjectFieldConfig{
						Type: typeInput,
					})
				}
			}
		}
	}
}

func inputFields(entity *meta.Entity, isPost bool) graphql.InputObjectConfigFieldMap {
	fields := graphql.InputObjectConfigFieldMap{}
	for _, column := range meta.Metas.EntityAllColumns(entity) {
		if column.Name != "id" || isPost {
			fields[column.Name] = &graphql.InputObjectFieldConfig{
				Type: ColumnType(&column),
			}
		}
	}
	return fields
}

func makeSaveInput(entity *meta.Entity) *graphql.InputObject {
	name := entity.Name + consts.INPUT

	return graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name:   name,
			Fields: inputFields(entity, true),
		},
	)
}

func makeUpdateInput(entity *meta.Entity) *graphql.InputObject {
	return graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name:   entity.Name + consts.UPDATE_INPUT,
			Fields: inputFields(entity, false),
		},
	)
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
