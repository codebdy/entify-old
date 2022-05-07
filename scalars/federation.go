package scalars

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/consts"
)

var AnyType = graphql.NewScalar(
	graphql.ScalarConfig{
		Name:        consts.SCALAR_ANY,
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

var FieldSetType = graphql.NewScalar(
	graphql.ScalarConfig{
		Name:        consts.SCALAR_FIELDSET,
		Description: "The `_FieldSet` scalar type represents _FieldSet values as specified by [Federation](https://www.apollographql.com/docs/federation/federation-spec)",
		Serialize: func(value interface{}) interface{} {
			return value
		},
		ParseValue: func(value interface{}) interface{} {
			return value
		},
		ParseLiteral: parseLiteral,
	},
)
