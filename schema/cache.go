package schema

import "github.com/graphql-go/graphql"

//where表达式缓存，query跟mutation都用
var WhereExpMap map[string]*graphql.InputObject

//类型缓存， query用
var OutputTypeMap map[string]*graphql.Output
var DistinctOnEnumMap map[string]*graphql.Enum
var OrderByMap map[string]*graphql.InputObject

var MutationResponseMap map[string]*graphql.Output

//类型缓存， query mutaion通用
var UpdateInputMap map[string]*graphql.Input
var PostInputMap map[string]*graphql.Input

var EnumMap map[string]*graphql.Enum
var EnumComparisonExpMap map[string]*graphql.InputObjectFieldConfig

func ClearCache() {
	WhereExpMap = make(map[string]*graphql.InputObject)
	OutputTypeMap = make(map[string]*graphql.Output)
	DistinctOnEnumMap = make(map[string]*graphql.Enum)
	OrderByMap = make(map[string]*graphql.InputObject)
	MutationResponseMap = make(map[string]*graphql.Output)
	UpdateInputMap = make(map[string]*graphql.Input)
	PostInputMap = make(map[string]*graphql.Input)
	EnumMap = make(map[string]*graphql.Enum)
	EnumComparisonExpMap = make(map[string]*graphql.InputObjectFieldConfig)
}
