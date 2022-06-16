package schema

import (
	"fmt"
	"strings"

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
		queryFields = queryFields + makeEntitySDL(entity)
	}

	for _, exteneral := range model.GlobalModel.Graph.RootExternals() {
		queryFields = queryFields + makeExteneralSDL(exteneral)
	}

	if config.AuthUrl() == "" {
		queryFields = queryFields + makeAuthSDL()
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
	sdl = sdl + fmt.Sprintf("%s(%s) %s \n",
		utils.FirstLower(intf.Name()),
		makeArgsSDL(quryeArgs(intf.Name())),
		queryResponseType(intf).Name(),
	)

	sdl = sdl + fmt.Sprintf(`%s(%s) %s`,
		consts.ONE+intf.Name(),
		makeArgsSDL(quryeArgs(intf.Name())),
		Cache.OutputType(intf.Name()).Name(),
	)

	sdl = sdl + fmt.Sprintf(`%s(%s) %s`,
		intf.Name()+utils.FirstUpper(consts.AGGREGATE),
		makeArgsSDL(quryeArgs(intf.Name())),
		(*AggregateType(intf)).Name(),
	)

	return sdl
}

func makeEntitySDL(entity *graph.Entity) string {
	sdl := ""
	sdl = sdl + fmt.Sprintf("%s(%s) %s \n",
		utils.FirstLower(entity.Name()),
		makeArgsSDL(quryeArgs(entity.Name())),
		queryResponseType(entity).Name(),
	)

	sdl = sdl + fmt.Sprintf("%s(%s) %s \n",
		consts.ONE+entity.Name(),
		makeArgsSDL(quryeArgs(entity.Name())),
		Cache.OutputType(entity.Name()).Name(),
	)

	sdl = sdl + fmt.Sprintf("%s(%s) %s \n",
		entity.Name()+utils.FirstUpper(consts.AGGREGATE),
		makeArgsSDL(quryeArgs(entity.Name())),
		(*AggregateType(entity)).Name(),
	)

	return sdl
}

func makeExteneralSDL(entity *graph.Entity) string {
	sdl := ""
	sdl = sdl + fmt.Sprintf("%s(%s) %s \n",
		utils.FirstLower(entity.Name()),
		makeArgsSDL(quryeArgs(entity.Name())),
		queryResponseType(entity).Name(),
	)

	sdl = sdl + fmt.Sprintf("%s(%s) %s \n",
		consts.ONE+entity.Name(),
		makeArgsSDL(quryeArgs(entity.Name())),
		Cache.OutputType(entity.Name()).Name(),
	)

	sdl = sdl + fmt.Sprintf("%s(%s) %s \n",
		entity.Name()+utils.FirstUpper(consts.AGGREGATE),
		makeArgsSDL(quryeArgs(entity.Name())),
		(*AggregateType(entity)).Name(),
	)

	return sdl
}

func makeArgsSDL(args graphql.FieldConfigArgument) string {
	var sdls []string
	for key := range args {
		sdls = append(sdls, key+":"+args[key].Type.Name())
	}
	return strings.Join(sdls, ",")
}

func makeAuthSDL() string {
	return fmt.Sprintf("\n me %s \n", baseUserType.Name())
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
