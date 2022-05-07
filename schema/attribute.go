package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entify/model/graph"
	"rxdrag.com/entify/model/meta"
	"rxdrag.com/entify/scalars"
)

func AttributeType(attr *graph.Attribute) graphql.Output {
	switch attr.Type {
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
		enum := attr.EumnType
		if enum == nil {
			panic("Can not find enum entity")
		}
		return Cache.EnumType(enum.Name)
	}

	panic("No column type:" + attr.Type)
}

func AttributeExp(column *graph.Attribute) *graphql.InputObjectFieldConfig {
	switch column.Type {
	case meta.INT:
		return &IntComparisonExp
	case meta.FLOAT:
		return &FloatComparisonExp
	case meta.BOOLEAN:
		return &BooleanComparisonExp
	case meta.STRING:
		return &StringComparisonExp
	case meta.DATE:
		return &DateTimeComparisonExp
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
		return nil
	case meta.ID:
		return &IdComparisonExp
	case meta.ENUM:
		return EnumComparisonExp(column)
	}

	panic("No column type: " + column.Type)
}

func AttributeOrderBy(column *graph.Attribute) *graphql.Enum {
	switch column.Type {
	case meta.VALUE_OBJECT,
		meta.BOOLEAN,
		meta.ENTITY,
		meta.ID_ARRAY,
		meta.INT_ARRAY,
		meta.FLOAT_ARRAY,
		meta.STRING_ARRAY,
		meta.DATE_ARRAY,
		meta.ENUM_ARRAY,
		meta.VALUE_OBJECT_ARRAY,
		meta.ENTITY_ARRAY:
		return nil
	}

	return EnumOrderBy
}
