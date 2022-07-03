package schema

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"rxdrag.com/entify/config"
	"rxdrag.com/entify/consts"
	"rxdrag.com/entify/model"
	"rxdrag.com/entify/model/graph"
	"rxdrag.com/entify/model/meta"
)

var mutationFieldSDL = "\t%s(%s) : %s \n"

func mutationSDL() (string, string) {
	queryFields := ""
	types := ""
	if config.AuthUrl() == "" {
		queryFields = queryFields + makeLoginSDL()
	}

	for _, entity := range model.GlobalModel.Graph.RootEnities() {
		if notSystemEntity(entity) {
			queryFields = queryFields + makeEntityMutationSDL(entity)
			types = types + objectToSDL(Cache.MutationResponse(entity.Name()), false)
		}
	}
	for _, partial := range model.GlobalModel.Graph.RootPartails() {
		queryFields = queryFields + makePartialMutationSDL(partial)
		types = types + objectToSDL(Cache.MutationResponse(partial.Name()), false)
	}

	// for _, exteneral := range model.GlobalModel.Graph.RootExternals() {
	// 	queryFields = queryFields + makeExteneralSDL(exteneral)
	// 	//types = types + objectToSDL(Cache.EntityeOutputType(exteneral.Name()))
	// }
	for _, input := range Cache.SetInputMap {
		if len(input.Fields()) > 0 &&
			input.Name() != meta.MetaClass.Name &&
			input.Name() != meta.MetaClass.Name+consts.SET &&
			input.Name() != meta.AbilityClass.Name &&
			input.Name() != meta.EntityAuthSettingsClass.Name {
			types = types + inputToSDL(input)
		}

	}
	for _, input := range Cache.SaveInputMap {
		types = types + inputToSDL(input)
	}
	for _, input := range Cache.HasManyInputMap {
		types = types + inputToSDL(input)
	}
	for _, input := range Cache.HasOneInputMap {
		types = types + inputToSDL(input)
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

	updateInput := Cache.SetInput(entity.Name())
	if len(updateInput.Fields()) > 0 {
		sdl = sdl + fmt.Sprintf(mutationFieldSDL,
			entity.SetName(),
			makeArgsSDL(setArgs(entity)),
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

func makePartialMutationSDL(partial *graph.Partial) string {
	sdl := ""
	sdl = sdl + fmt.Sprintf(mutationFieldSDL,
		partial.DeleteName(),
		makeArgsSDL(deleteArgs(&partial.Entity)),
		Cache.MutationResponse(partial.Name()).Name(),
	)

	sdl = sdl + fmt.Sprintf(mutationFieldSDL,
		partial.DeleteByIdName(),
		makeArgsSDL(deleteByIdArgs()),
		Cache.OutputType(partial.Name()).String(),
	)

	sdl = sdl + fmt.Sprintf(mutationFieldSDL,
		partial.InsertName(),
		makeArgsSDL(upsertArgs(&partial.Entity)),
		(&graphql.List{OfType: Cache.OutputType(partial.Name())}).String(),
	)

	sdl = sdl + fmt.Sprintf(mutationFieldSDL,
		partial.InsertOneName(),
		makeArgsSDL(upsertOneArgs(&partial.Entity)),
		Cache.OutputType(partial.Name()).String(),
	)

	updateInput := Cache.SetInput(partial.Name())
	if len(updateInput.Fields()) > 0 {
		sdl = sdl + fmt.Sprintf(mutationFieldSDL,
			partial.SetName(),
			makeArgsSDL(setArgs(&partial.Entity)),
			Cache.MutationResponse(partial.Name()).String(),
		)
	}

	sdl = sdl + fmt.Sprintf(mutationFieldSDL,
		partial.UpdateName(),
		makeArgsSDL(upsertArgs(&partial.Entity)),
		(&graphql.List{OfType: Cache.OutputType(partial.Name())}).String(),
	)

	sdl = sdl + fmt.Sprintf(mutationFieldSDL,
		partial.UpdateOneName(),
		makeArgsSDL(upsertOneArgs(&partial.Entity)),
		Cache.OutputType(partial.Name()).String(),
	)

	return sdl
}

func makeLoginSDL() string {
	return `	login(loginName: String!password: String!): String
	logout: Boolean`
}
