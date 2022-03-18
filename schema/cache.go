package schema

import "github.com/graphql-go/graphql"

type TypeCache struct {
	WhereExpMap map[string]*graphql.InputObject

	//类型缓存， query用
	OutputTypeMap     map[string]*graphql.Output
	DistinctOnEnumMap map[string]*graphql.Enum
	OrderByMap        map[string]*graphql.InputObject

	MutationResponseMap map[string]*graphql.Output

	//类型缓存， query mutaion通用
	UpdateInputMap map[string]*graphql.Input
	PostInputMap   map[string]*graphql.Input

	EnumMap              map[string]*graphql.Enum
	EnumComparisonExpMap map[string]*graphql.InputObjectFieldConfig

	AggregateMap map[string]*graphql.Output
}

//where表达式缓存，query跟mutation都用

func (c *TypeCache) ClearCache() {
	c.WhereExpMap = make(map[string]*graphql.InputObject)
	c.OutputTypeMap = make(map[string]*graphql.Output)
	c.DistinctOnEnumMap = make(map[string]*graphql.Enum)
	c.OrderByMap = make(map[string]*graphql.InputObject)
	c.MutationResponseMap = make(map[string]*graphql.Output)
	c.UpdateInputMap = make(map[string]*graphql.Input)
	c.PostInputMap = make(map[string]*graphql.Input)
	c.EnumMap = make(map[string]*graphql.Enum)
	c.EnumComparisonExpMap = make(map[string]*graphql.InputObjectFieldConfig)
	c.AggregateMap = make(map[string]*graphql.Output)
}
