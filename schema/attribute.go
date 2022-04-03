package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/model/graph"
	"rxdrag.com/entity-engine/oldmeta"
	"rxdrag.com/entity-engine/scalars"
)

func AttributeType(attr *graph.Attribute) graphql.Output {
	switch attr.Type {
	case oldmeta.COLUMN_ID:
		return graphql.ID
	case oldmeta.COLUMN_INT:
		return graphql.Int
	case oldmeta.COLUMN_FLOAT:
		return graphql.Float
	case oldmeta.COLUMN_BOOLEAN:
		return graphql.Boolean
	case oldmeta.COLUMN_STRING:
		return graphql.String
	case oldmeta.COLUMN_DATE:
		return graphql.DateTime
	case oldmeta.COLUMN_SIMPLE_JSON, oldmeta.COLUMN_SIMPLE_ARRAY, oldmeta.COLUMN_JSON_ARRAY:
		return scalars.JSONType
	case oldmeta.COLUMN_ENUM:
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
	case oldmeta.COLUMN_INT:
		return &IntComparisonExp
	case oldmeta.COLUMN_FLOAT:
		return &FloatComparisonExp
	case oldmeta.COLUMN_BOOLEAN:
		return &BooleanComparisonExp
	case oldmeta.COLUMN_STRING:
		return &StringComparisonExp
	case oldmeta.COLUMN_DATE:
		return &DateTimeComparisonExp
	case oldmeta.COLUMN_SIMPLE_JSON, oldmeta.COLUMN_SIMPLE_ARRAY, oldmeta.COLUMN_JSON_ARRAY:
		return nil
	case oldmeta.COLUMN_ID:
		return &IdComparisonExp
	case oldmeta.COLUMN_ENUM:
		return EnumComparisonExp(column)
	}

	panic("No column type: " + column.Type)
}

func AttributeOrderBy(column *graph.Attribute) *graphql.Enum {
	switch column.Type {
	case oldmeta.COLUMN_SIMPLE_JSON:
		return nil
	case oldmeta.COLUMN_SIMPLE_ARRAY:
		return nil
	case oldmeta.COLUMN_JSON_ARRAY:
		return nil
	}

	return EnumOrderBy
}
