package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/meta"
	"rxdrag.com/entity-engine/scalars"
	"rxdrag.com/entity-engine/schema/comparison"
)

func ColumnType(column *meta.Column) graphql.Output {
	switch column.Type {
	case meta.COLUMN_ID:
		return graphql.String
	case meta.COLUMN_INT:
		return graphql.Int
	case meta.COLUMN_FLOAT:
		return graphql.Float
	case meta.COLUMN_BOOLEAN:
		return graphql.Boolean
	case meta.COLUMN_STRING:
		return graphql.String
	case meta.COLUMN_DATE:
		return graphql.DateTime
	case meta.COLUMN_SIMPLE_JSON:
		return scalars.JSONType
	case meta.COLUMN_SIMPLE_ARRAY:
		return graphql.NewScalar(graphql.ScalarConfig{Name: "JSON"})
	case meta.COLUMN_JSON_ARRAY:
		return graphql.NewScalar(graphql.ScalarConfig{Name: "JSON"})
	case meta.COLUMN_ENUM:
		return graphql.EnumValueType
	}

	panic("No column type:" + column.Type)
}

func ColumnExp(column *meta.Column) *graphql.InputObjectFieldConfig {
	switch column.Type {
	case meta.COLUMN_INT:
		return &comparison.IntComparisonExp
	case meta.COLUMN_FLOAT:
		return &comparison.FloatComparisonExp
	case meta.COLUMN_BOOLEAN:
		return &comparison.BooleanComparisonExp
	case meta.COLUMN_STRING:
		return &comparison.StringComparisonExp
	case meta.COLUMN_DATE:
		return &comparison.DateTimeComparisonExp
	case meta.COLUMN_SIMPLE_JSON:
		return nil
	case meta.COLUMN_SIMPLE_ARRAY:
		return nil
	case meta.COLUMN_JSON_ARRAY:
		return nil
	case meta.COLUMN_ID:
		return nil
	case meta.COLUMN_ENUM:
		return &comparison.EnumComparisonExp
	}

	panic("No column type: " + column.Type)
}

func ColumnOrderBy(column *meta.Column) *graphql.Enum {
	switch column.Type {
	case meta.COLUMN_SIMPLE_JSON:
		return nil
	case meta.COLUMN_SIMPLE_ARRAY:
		return nil
	case meta.COLUMN_JSON_ARRAY:
		return nil
	}

	return EnumOrderBy
}
