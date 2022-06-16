package schema

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"rxdrag.com/entify/config"
	"rxdrag.com/entify/consts"
	"rxdrag.com/entify/model"
	"rxdrag.com/entify/model/graph"
	"rxdrag.com/entify/utils"
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

	queryFields := ""
	for _, intf := range model.GlobalModel.Graph.RootInterfaces() {
		queryFields = queryFields + makeInterfaceSDL(intf)
	}

	for _, entity := range model.GlobalModel.Graph.RootEnities() {
		appendEntityToQueryFields(entity, queryFields)
	}

	for _, service := range model.GlobalModel.Graph.RootExternals() {
		appendServiceQueryFields(service, queryFields)
	}

	if config.AuthUrl() == "" {
		appendAuthToQuery(queryFields)
	}
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

func makeInterfaceSDL(intf *graph.Interface) string {
	sdl := ""
	sdl = sdl + fmt.Sprintf(`%s(%s) %s`,
		utils.FirstLower(intf.Name()),
		makeArgsSDL(quryeArgs(intf.Name())),
		fmt.Sprintf("![%s]\n", queryResponseType(intf).Name()),
	)

	sdl = sdl + sdl + fmt.Sprintf(`%s(%s) %s`,
		consts.ONE+intf.Name(),
		makeArgsSDL(quryeArgs(intf.Name())),
		Cache.OutputType(intf.Name()).Name(),
	)

	sdl = sdl + sdl + fmt.Sprintf(`%s(%s) %s`,
		intf.Name()+utils.FirstUpper(consts.AGGREGATE),
		makeArgsSDL(quryeArgs(intf.Name())),
		(*AggregateType(intf)).Name(),
	)

	return sdl
}

func makeArgsSDL(args graphql.FieldConfigArgument) string {
	return ""
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
