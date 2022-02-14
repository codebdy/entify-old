package schema

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/model"
)

func AppendEntityToQueryFields(entity model.EntityMeta, feilds *graphql.Fields) {
	(*feilds)[entity.Name] = &graphql.Field{
		Type: graphql.String,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			fmt.Println(p.Context.Value("data"))
			return "world", nil
		},
	}
}
