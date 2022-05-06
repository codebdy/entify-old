package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/model"
	"rxdrag.com/entity-engine/resolve"
)

var EntityType *graphql.Union

func publishResolve(p graphql.ResolveParams) (interface{}, error) {
	reslult, err := resolve.PublishMetaResolve(p)
	if err != nil {
		return reslult, err
	}

	MakeSchema()
	return reslult, nil
}

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
				Name:      "testDirective",
				Locations: []string{graphql.DirectiveLocationField},
			}),
		},
		Types: Cache.EntityTypes(),
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

func init() {
	LoadModel()
	MakeSchema()
}
