package graph

import "github.com/graphql-go/graphql"

type GqlNode interface {
	OutputType() graphql.Type
	InputType() *graphql.InputObject
}
