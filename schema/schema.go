package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/authentication"
	"rxdrag.com/entity-engine/authentication/jwt"
	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/meta"
	"rxdrag.com/entity-engine/repository"
	"rxdrag.com/entity-engine/resolve"
)

var Cache TypeCache
var GQLSchema *graphql.Schema

func publishResolve(p graphql.ResolveParams) (interface{}, error) {
	reslult, err := resolve.PublishMetaResolve(p)
	if err != nil {
		return reslult, err
	}

	MakeSchema()
	return reslult, nil
}

func MakeSchema() {
	Cache.ClearCache()

	queryFields := graphql.Fields{}

	for _, entity := range *repository.Entities {
		AppendToQueryFields(entity, &queryFields)
	}

	mutationFields := graphql.Fields{
		consts.LOGIN: &graphql.Field{
			Type: graphql.String,
			Args: graphql.FieldConfigArgument{
				consts.LOGIN_NAME: &graphql.ArgumentConfig{
					Type: &graphql.NonNull{OfType: graphql.String},
				},
				consts.PASSWORD: &graphql.ArgumentConfig{
					Type: &graphql.NonNull{OfType: graphql.String},
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				loginName, err := authentication.Login(p.Args[consts.LOGIN_NAME].(string), p.Args[consts.PASSWORD].(string))
				if err != nil {
					return "", err
				}
				return jwt.GenerateToken(loginName)
			},
		},
		consts.LOGOUT: &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world2", nil
			},
		},
		consts.PUBLISH: &graphql.Field{
			Type:    *OutputType(&meta.MetaEntity),
			Resolve: publishResolve,
		},
		consts.ROLLBACK: &graphql.Field{
			Type:    *OutputType(&meta.MetaEntity),
			Resolve: resolve.SyncMetaResolve,
		},
		consts.SYNC_META: &graphql.Field{
			Type:    *OutputType(&meta.MetaEntity),
			Resolve: resolve.SyncMetaResolve,
		},
	}

	for _, entity := range *repository.Entities {
		AppendToMutationFields(entity, &mutationFields)
	}

	rootQuery := graphql.ObjectConfig{Name: consts.ROOT_QUERY_NAME, Fields: queryFields}
	rootMutation := graphql.ObjectConfig{Name: consts.ROOT_MUTATION_NAME, Fields: mutationFields}
	schemaConfig := graphql.SchemaConfig{
		Query:    graphql.NewObject(rootQuery),
		Mutation: graphql.NewObject(rootMutation),
		Directives: []*graphql.Directive{
			graphql.NewDirective(graphql.DirectiveConfig{
				Name:      "testDirective",
				Locations: []string{graphql.DirectiveLocationField},
			}),
		},
	}
	theSchema, err := graphql.NewSchema(schemaConfig)

	if err != nil {
		panic(err)
		//log.Fatalf("failed to create new schema, error: %v", err)
	}
	GQLSchema = &theSchema
}

func ResolveSchema() *graphql.Schema {
	return GQLSchema
}

func init() {
	MakeSchema()
}
