package scalars

import "github.com/graphql-go/graphql"

// Upload type
var UploadType = graphql.NewScalar(
	graphql.ScalarConfig{
		Name:        "Upload",
		Description: "The `Upload` scalar type ",
		Serialize: func(value interface{}) interface{} {
			return value
		},
		ParseValue: func(value interface{}) interface{} {
			return value
		},
		ParseLiteral: parseLiteral,
	},
)
