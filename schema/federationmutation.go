package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entify/config"
	"rxdrag.com/entify/consts"
	"rxdrag.com/entify/model"
)

func mutationSDL() (string, string) {
	queryFields := ""
	types := ""
	if config.AuthUrl() == "" {
		queryFields = queryFields + makeAuthSDL()
		types = types + objectToSDL(baseRoleTye)
		types = types + objectToSDL(baseUserType)
	}

	for _, enum := range model.GlobalModel.Graph.Enums {
		types = types + enumToSDL(Cache.EnumType(enum.Name))
	}

	for _, enum := range Cache.DistinctOnEnums() {
		types = types + enumToSDL(enum)
	}

	types = types + enumToSDL(EnumOrderBy)
	for _, orderBy := range Cache.OrderByMap {
		types = types + inputToSDL(orderBy)
	}

	types = types + comparisonToSDL(&BooleanComparisonExp)
	types = types + comparisonToSDL(&DateTimeComparisonExp)
	types = types + comparisonToSDL(&FloatComparisonExp)
	types = types + comparisonToSDL(&IntComparisonExp)
	types = types + comparisonToSDL(&IdComparisonExp)
	types = types + comparisonToSDL(&StringComparisonExp)

	for _, comparision := range Cache.EnumComparisonExpMap {
		types = types + comparisonToSDL(comparision)
	}

	for _, where := range Cache.WhereExpMap {
		types = types + inputToSDL(where)
	}

	for _, intf := range model.GlobalModel.Graph.Interfaces {
		types = types + interfaceToSDL(Cache.InterfaceOutputType(intf.Name()))
	}

	for _, intf := range model.GlobalModel.Graph.RootInterfaces() {
		queryFields = queryFields + makeInterfaceSDL(intf)
	}

	for _, entity := range model.GlobalModel.Graph.Entities {
		types = types + objectToSDL(Cache.EntityeOutputType(entity.Name()))
	}
	for _, entity := range model.GlobalModel.Graph.RootEnities() {
		queryFields = queryFields + makeEntitySDL(entity)
	}

	for _, exteneral := range model.GlobalModel.Graph.RootExternals() {
		queryFields = queryFields + makeExteneralSDL(exteneral)
		//types = types + objectToSDL(Cache.EntityeOutputType(exteneral.Name()))
	}

	for _, aggregate := range Cache.AggregateMap {
		types = types + objectToSDL(aggregate)
		fieldsType := aggregate.Fields()[consts.AGGREGATE].Type.(*graphql.Object)
		types = types + objectToSDL(fieldsType)
	}

	for _, selectColumn := range Cache.SelectColumnsMap {
		types = types + inputToSDL(selectColumn)
	}
	return queryFields, types
}
