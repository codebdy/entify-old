package grahp

import "github.com/graphql-go/graphql"

type Enum struct {
}

func (e *Enum) GqlType() *graphql.Enum {
	return nil
}
