package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/model/graph"
	"rxdrag.com/entity-engine/model/meta"
	"rxdrag.com/entity-engine/scalars"
)

func MethodType(method *graph.Method) graphql.Output {
	switch method.Method.Type {
	case meta.ID:
		return graphql.ID
	case meta.INT:
		return graphql.Int
	case meta.FLOAT:
		return graphql.Float
	case meta.BOOLEAN:
		return graphql.Boolean
	case meta.STRING:
		return graphql.String
	case meta.DATE:
		return graphql.DateTime
	case meta.VALUE_OBJECT,
		meta.ENTITY,
		meta.ID_ARRAY,
		meta.INT_ARRAY,
		meta.FLOAT_ARRAY,
		meta.STRING_ARRAY,
		meta.DATE_ARRAY,
		meta.ENUM_ARRAY,
		meta.VALUE_OBJECT_ARRAY,
		meta.ENTITY_ARRAY:
		return scalars.JSONType
	case meta.ENUM:
		enum := method.EumnType
		if enum == nil {
			panic("Can not find enum entity")
		}
		return Cache.EnumType(enum.Name)
	}
	panic("No column type:" + method.Method.Type)
}
