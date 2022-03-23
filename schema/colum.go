package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/meta"
	"rxdrag.com/entity-engine/repository"
	"rxdrag.com/entity-engine/scalars"
)

func ColumnType(column *meta.ColumnMeta) graphql.Output {
	switch column.Type {
	case meta.COLUMN_ID:
		return graphql.ID
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
	case meta.COLUMN_SIMPLE_JSON, meta.COLUMN_SIMPLE_ARRAY, meta.COLUMN_JSON_ARRAY:
		return scalars.JSONType
	case meta.COLUMN_ENUM:
		enumEntity := repository.GetEntityByUuid(column.EnumUuid)
		if enumEntity == nil {
			panic("Can not find enum entity")
		}
		return Cache.OutputType(enumEntity)
	}

	panic("No column type:" + column.Type)
}

func ColumnExp(column *meta.ColumnMeta) *graphql.InputObjectFieldConfig {
	switch column.Type {
	case meta.COLUMN_INT:
		return &IntComparisonExp
	case meta.COLUMN_FLOAT:
		return &FloatComparisonExp
	case meta.COLUMN_BOOLEAN:
		return &BooleanComparisonExp
	case meta.COLUMN_STRING:
		return &StringComparisonExp
	case meta.COLUMN_DATE:
		return &DateTimeComparisonExp
	case meta.COLUMN_SIMPLE_JSON, meta.COLUMN_SIMPLE_ARRAY, meta.COLUMN_JSON_ARRAY:
		return nil
	case meta.COLUMN_ID:
		return &IdComparisonExp
	case meta.COLUMN_ENUM:
		return EnumComparisonExp(column)
	}

	panic("No column type: " + column.Type)
}

func ColumnOrderBy(column *meta.ColumnMeta) *graphql.Enum {
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
