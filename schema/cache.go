package schema

import "github.com/graphql-go/graphql"

//where表达式缓存，query跟mutation都用
var WhereExpMap = make(map[string]*graphql.InputObject)

//类型缓存， query用
var OutputTypeMap = make(map[string]*graphql.Output)
var DistinctOnEnumMap = make(map[string]*graphql.Enum)
var OrderByMap = make(map[string]*graphql.InputObject)

//类型缓存， query mutaion通用
var UpdateInputMap = make(map[string]*graphql.Input)
var PostInputMap = make(map[string]*graphql.Input)

var EnumMap = make(map[string]*graphql.Enum)
var EnumComparisonExpMap = make(map[string]*graphql.InputObjectFieldConfig)
