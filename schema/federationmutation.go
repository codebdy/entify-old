package schema

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"rxdrag.com/entify/config"
	"rxdrag.com/entify/model"
	"rxdrag.com/entify/model/graph"
)

var mutationFieldSDL = "\t%s(%s) : %s \n"

func mutationSDL() (string, string) {
	queryFields := ""
	types := ""
	if config.AuthUrl() == "" {
		queryFields = queryFields + makeLoginSDL()
	}

	for _, entity := range model.GlobalModel.Graph.Entities {
		types = types + objectToSDL(Cache.EntityeOutputType(entity.Name()))
	}
	for _, entity := range model.GlobalModel.Graph.RootEnities() {
		queryFields = queryFields + makeEntityMutationSDL(entity)
	}

	// for _, exteneral := range model.GlobalModel.Graph.RootExternals() {
	// 	queryFields = queryFields + makeExteneralSDL(exteneral)
	// 	//types = types + objectToSDL(Cache.EntityeOutputType(exteneral.Name()))
	// }

	for _, responseType := range Cache.MutationResponseMap {
		types = types + objectToSDL(responseType)
	}

	return queryFields, types
}

func makeEntityMutationSDL(entity *graph.Entity) string {
	sdl := ""
	sdl = sdl + fmt.Sprintf(mutationFieldSDL,
		entity.DeleteName(),
		makeArgsSDL(deleteArgs(entity)),
		Cache.MutationResponse(entity.Name()).Name(),
	)

	sdl = sdl + fmt.Sprintf(mutationFieldSDL,
		entity.DeleteByIdName(),
		makeArgsSDL(deleteByIdArgs()),
		Cache.OutputType(entity.Name()).String(),
	)

	updateInput := Cache.UpdateInput(entity.Name())
	if len(updateInput.Fields()) > 0 {
		sdl = sdl + fmt.Sprintf(mutationFieldSDL,
			entity.UpdateName(),
			makeArgsSDL(updateArgs(entity)),
			Cache.MutationResponse(entity.Name()).String(),
		)
	}

	sdl = sdl + fmt.Sprintf(mutationFieldSDL,
		entity.UpsertName(),
		makeArgsSDL(upsertArgs(entity)),
		(&graphql.List{OfType: Cache.OutputType(entity.Name())}).String(),
	)

	sdl = sdl + fmt.Sprintf(mutationFieldSDL,
		entity.UpsertOneName(),
		makeArgsSDL(upsertOneArgs(entity)),
		Cache.OutputType(entity.Name()).String(),
	)

	return sdl
}

func makeLoginSDL() string {
	return `	login(loginName: String!password: String!): String
	logout: Boolean`
}
