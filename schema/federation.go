package schema

import (
	"fmt"
)

var allSDL = `
extend schema
	@link(url: "https://specs.apollo.dev/federation/v2.0",
		import: ["@key"])

scalar JSON
scalar DateTime

type Query {
%s
}

type Mutation {
%s
}
%s
`

func makeFederationSDL() string {
	sdl := allSDL
	queryFields, queryTypes := querySDL()
	mutationFields, mutationTypes := mutationSDL()
	return fmt.Sprintf(sdl, queryFields, mutationFields, queryTypes+mutationTypes)
}
