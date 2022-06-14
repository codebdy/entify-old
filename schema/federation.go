package schema

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"rxdrag.com/entify/config"
	"rxdrag.com/entify/consts"
)

func makeFederationSDL() string {
	sdl := `
		extend schema
		@link(url: "https://specs.apollo.dev/federation/v2.0",
			import: ["@key", "@shareable"])

		extend type Query {
			%s
		}

		extend type Mutation {
			%s
		}

		%s
	`

	queryFields := "me: User"
	mutationFields := "review(date: String review: String): Result"
	types := `
	type User {
		id: ID!
		username: String
	}

	type Result {
		success: Boolean
	}
	`

	return fmt.Sprintf(sdl, queryFields, mutationFields, types)
}

func serviceField() *graphql.Field {
	return &graphql.Field{
		Type: _ServiceType,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return map[string]interface{}{
				consts.ID:  config.ServiceId(),
				consts.SDL: makeFederationSDL(),
			}, nil
		},
	}
}
