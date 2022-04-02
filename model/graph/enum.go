package graph

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/model/domain"
)

type Enum struct {
	domain.Enum
}

func (e *Enum) GqlType() *graphql.Enum {
	return nil
}
