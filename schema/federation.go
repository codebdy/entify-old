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

var allSDL = `
extend schema
@link(url: "https://specs.apollo.dev/federation/v2.0",
	import: ["@key", "@shareable"])

scalar JSON

extend type Query {
%s
}

extend type Mutation {
%s
}
%s
`

var queryFieldSDL = "\t%s(%s) : %s \n"

var objectSDL = `
type %s%s{
	%s
}
`

var enumSDL = `
enum %s{
	%s
}
`

var interfaceSDL = `
interface %s{
	%s
}
`

var inputSDL = `
input %s{
	%s
}
`

var comparisonSDL = `
input %s{
	%s
}
`

func makeFederationSDL() string {
	sdl := allSDL

	queryFields := ""
	mutationFields := "review(date: String review: String): Result"
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

	for _, intf := range model.GlobalModel.Graph.RootInterfaces() {
		queryFields = queryFields + makeInterfaceSDL(intf)

		types = types + interfaceToSDL(Cache.InterfaceOutputType(intf.Name()))
	}

	for _, entity := range model.GlobalModel.Graph.RootEnities() {
		queryFields = queryFields + makeEntitySDL(entity)

		types = types + objectToSDL(Cache.EntityeOutputType(entity.Name()))
	}

	for _, exteneral := range model.GlobalModel.Graph.RootExternals() {
		queryFields = queryFields + makeExteneralSDL(exteneral)
		//types = types + objectToSDL(Cache.EntityeOutputType(exteneral.Name()))
	}

	for _, aggregate := range Cache.AggregateMap {
		types = types + objectToSDL(aggregate)
		fieldsType := aggregate.Fields()[consts.AGGREGATE].Type
		types = types + objectToSDL(fieldsType.(*graphql.Object))
	}

	return fmt.Sprintf(sdl, queryFields, mutationFields, types)
}

func makeInterfaceSDL(intf *graph.Interface) string {
	sdl := ""
	sdl = sdl + fmt.Sprintf(queryFieldSDL,
		intf.QueryName(),
		makeArgsSDL(quryeArgs(intf.Name())),
		queryResponseType(intf).String(),
	)

	sdl = sdl + fmt.Sprintf(queryFieldSDL,
		intf.QueryOneName(),
		makeArgsSDL(quryeArgs(intf.Name())),
		Cache.OutputType(intf.Name()).String(),
	)

	sdl = sdl + fmt.Sprintf(queryFieldSDL,
		intf.QueryAggregateName(),
		makeArgsSDL(quryeArgs(intf.Name())),
		(*AggregateType(intf)).String(),
	)

	return sdl
}

func makeEntitySDL(entity *graph.Entity) string {
	sdl := ""
	sdl = sdl + fmt.Sprintf(queryFieldSDL,
		entity.QueryName(),
		makeArgsSDL(quryeArgs(entity.Name())),
		queryResponseType(entity).String(),
	)

	sdl = sdl + fmt.Sprintf(queryFieldSDL,
		entity.QueryOneName(),
		makeArgsSDL(quryeArgs(entity.Name())),
		Cache.OutputType(entity.Name()).String(),
	)

	sdl = sdl + fmt.Sprintf(queryFieldSDL,
		entity.QueryAggregateName(),
		makeArgsSDL(quryeArgs(entity.Name())),
		(*AggregateType(entity)).String(),
	)

	return sdl
}

func makeExteneralSDL(entity *graph.Entity) string {
	sdl := ""
	sdl = sdl + fmt.Sprintf(queryFieldSDL,
		entity.QueryName(),
		makeArgsSDL(quryeArgs(entity.Name())),
		queryResponseType(entity).String(),
	)

	sdl = sdl + fmt.Sprintf(queryFieldSDL,
		consts.ONE+entity.Name(),
		makeArgsSDL(quryeArgs(entity.Name())),
		Cache.OutputType(entity.Name()).String(),
	)

	sdl = sdl + fmt.Sprintf(queryFieldSDL,
		entity.Name()+utils.FirstUpper(consts.AGGREGATE),
		makeArgsSDL(quryeArgs(entity.Name())),
		(*AggregateType(entity)).String(),
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

func makeArgArraySDL(args []*graphql.Argument) string {
	var sdls []string
	for _, arg := range args {
		sdls = append(sdls, arg.Name()+":"+arg.Type.Name())
	}
	return strings.Join(sdls, ",")
}

func makeAuthSDL() string {
	return fmt.Sprintf("\tme : %s \n", baseUserType.Name())
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

func objectToSDL(obj *graphql.Object) string {
	var intfNames []string
	implString := ""

	for _, intf := range obj.Interfaces() {
		intfNames = append(intfNames, intf.Name())
	}
	if len(intfNames) > 0 {
		implString = " implements " + strings.Join(intfNames, " & ")
	}

	sdl := objectSDL
	return fmt.Sprintf(sdl, obj.Name(), implString, fieldsToSDL(obj.Fields()))
}

func enumToSDL(enum *graphql.Enum) string {
	var values []string

	sdl := enumSDL
	for _, value := range enum.Values() {
		values = append(values, value.Name)
	}
	return fmt.Sprintf(sdl, enum.Name(), strings.Join(values, "\n\t"))
}

func interfaceToSDL(intf *graphql.Interface) string {
	sdl := interfaceSDL
	return fmt.Sprintf(sdl, intf.Name(), fieldsToSDL(intf.Fields()))
}

func inputToSDL(input *graphql.InputObject) string {
	sdl := inputSDL
	return fmt.Sprintf(sdl, input.Name(), inputFieldsToSDL(input.Fields()))
}

func inputFieldsToSDL(fields graphql.InputObjectFieldMap) string {
	var fieldsStrings []string
	for key := range fields {
		field := fields[key]
		fieldsStrings = append(fieldsStrings, key+":"+field.Type.String())
	}

	return strings.Join(fieldsStrings, "\n\t")
}

func comparisonToSDL(comarison *graphql.InputObjectFieldConfig) string {
	sdl := comparisonSDL
	var comType *graphql.InputObject
	comType = comarison.Type.(*graphql.InputObject)
	return fmt.Sprintf(sdl, comType.Name(), inputFieldsToSDL(comType.Fields()))
}

func fieldsToSDL(fields graphql.FieldDefinitionMap) string {
	var fieldsStrings []string
	for i := range fields {
		field := fields[i]
		if len(field.Args) > 0 {
			fieldsStrings = append(fieldsStrings, fmt.Sprintf("%s(%s):%s", field.Name, makeArgArraySDL(field.Args), field.Type.String()))
		} else {
			fieldsStrings = append(fieldsStrings, field.Name+":"+field.Type.String())
		}
	}

	return strings.Join(fieldsStrings, "\n\t")
}
