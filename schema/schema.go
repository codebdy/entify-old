package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entify/consts"
	"rxdrag.com/entify/model"
	"rxdrag.com/entify/repository"
	"rxdrag.com/entify/scalars"
)

var EntityType *graphql.Union

func MakeSchema() {
	Cache.MakeCache()

	EntityType = graphql.NewUnion(
		graphql.UnionConfig{
			Name:        consts.ENTITY_TYPE,
			Types:       Cache.EntityObjects(),
			ResolveType: resolveTypeFn,
		},
	)

	schemaConfig := graphql.SchemaConfig{
		Query:        rootQuery(),
		Mutation:     rootMutation(),
		Subscription: RootSubscription(),
		Directives: []*graphql.Directive{
			graphql.NewDirective(graphql.DirectiveConfig{
				Name:      consts.EXTERNAL,
				Locations: []string{graphql.DirectiveLocationField},
			}),
			graphql.NewDirective(graphql.DirectiveConfig{
				Name:      consts.REQUIRES,
				Locations: []string{graphql.DirectiveLocationField},
				Args: graphql.FieldConfigArgument{
					consts.ARG_FIELDS: &graphql.ArgumentConfig{
						Type: &graphql.NonNull{
							OfType: scalars.FieldSetType,
						},
					},
				},
			}),
			graphql.NewDirective(graphql.DirectiveConfig{
				Name:      consts.PROVIDES,
				Locations: []string{graphql.DirectiveLocationField},
				Args: graphql.FieldConfigArgument{
					consts.ARG_FIELDS: &graphql.ArgumentConfig{
						Type: &graphql.NonNull{
							OfType: scalars.FieldSetType,
						},
					},
				},
			}),
			graphql.NewDirective(graphql.DirectiveConfig{
				Name:      consts.KEY,
				Locations: []string{graphql.DirectiveLocationObject, graphql.DirectiveLocationInterface},
				Args: graphql.FieldConfigArgument{
					consts.ARG_FIELDS: &graphql.ArgumentConfig{
						Type: &graphql.NonNull{
							OfType: scalars.FieldSetType,
						},
					},
				},
			}),
			graphql.NewDirective(graphql.DirectiveConfig{
				Name:      consts.EXTENDS,
				Locations: []string{graphql.DirectiveLocationObject, graphql.DirectiveLocationInterface},
			}),
		},
		Types: append(Cache.EntityTypes(), scalars.FieldSetType),
	}
	theSchema, err := graphql.NewSchema(schemaConfig)

	if err != nil {
		panic(err)
		//log.Fatalf("failed to create new schema, error: %v", err)
	}
	model.GlobalModel.Schema = &theSchema
}

func ResolveSchema() *graphql.Schema {
	return model.GlobalModel.Schema
}

func InitSchema() {
	repository.InitGlobalModel()
	repository.LoadModel()
	LoadModel()
	MakeSchema()
}
