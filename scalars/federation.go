package scalars

import (
	"github.com/graphql-go/graphql"
)

var _AnyType = graphql.NewScalar(
	graphql.ScalarConfig{
		Name:        "_Any",
		Description: "The `_Any` scalar type represents _Any values as specified by [Federation](https://www.apollographql.com/docs/federation/federation-spec)",
		Serialize: func(value interface{}) interface{} {
			return value
		},
		ParseValue: func(value interface{}) interface{} {
			return value
		},
		ParseLiteral: parseLiteral,
	},
)

var _FieldSetType = graphql.NewScalar(
	graphql.ScalarConfig{
		Name:        "_FieldSet",
		Description: "The `_Any` scalar type represents _Any values as specified by [Federation](https://www.apollographql.com/docs/federation/federation-spec)",
		Serialize: func(value interface{}) interface{} {
			return value
		},
		ParseValue: func(value interface{}) interface{} {
			return value
		},
		ParseLiteral: parseLiteral,
	},
)
