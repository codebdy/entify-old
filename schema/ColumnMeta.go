package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/schema/comparison"
)

const (
	COLUMN_ID           string = "ID"
	COLUMN_INT          string = "Int"
	COLUMN_FLOAT        string = "Float"
	COLUMN_BOOLEAN      string = "Boolean"
	COLUMN_STRING       string = "String"
	COLUMN_DATE         string = "Date"
	COLUMN_SIMPLE_JSON  string = "SimpleJson"
	COLUMN_SIMPLE_ARRAY string = "simpleArray"
	COLUMN_JSON_ARRAY   string = "JsonArray"
	COLUMN_ENUM         string = "Enum"
)

type ColumnMeta struct {
	Uuid          string `json:"uuid"`
	Type          string `json:"type"`
	Name          string `json:"name"`
	Primary       bool   `json:"primary"`
	Generated     bool   `json:"generated"`
	Nullable      bool   `json:"nullable"`
	Unique        bool   `json:"unique"`
	Index         bool   `json:"index"`
	CreateDate    bool   `json:"createDate"`
	UpdateDate    bool   `json:"updateDate"`
	DeleteDate    bool   `json:"deleteDate"`
	Select        bool   `json:"select"`
	Length        int    `json:"length"`
	TypeEnityUuid string `json:"typeEnityUuid"`
	Description   string `json:"description"`
}

func (column *ColumnMeta) toType() graphql.Output {
	switch column.Type {
	case COLUMN_ID:
		return graphql.String
	case COLUMN_INT:
		return graphql.Int
	case COLUMN_FLOAT:
		return graphql.Float
	case COLUMN_BOOLEAN:
		return graphql.Boolean
	case COLUMN_STRING:
		return graphql.String
		return graphql.String
	case COLUMN_DATE:
		return graphql.DateTime
	case COLUMN_SIMPLE_JSON:
		return graphql.NewScalar(graphql.ScalarConfig{Name: "JSON"})
	case COLUMN_SIMPLE_ARRAY:
		return graphql.NewScalar(graphql.ScalarConfig{Name: "JSON"})
	case COLUMN_JSON_ARRAY:
		return graphql.NewScalar(graphql.ScalarConfig{Name: "JSON"})
	case COLUMN_ENUM:
		return graphql.EnumValueType
	}

	panic("No column type:" + column.Type)
}

func (column *ColumnMeta) ToExp() *graphql.InputObjectFieldConfig {
	switch column.Type {
	case COLUMN_INT:
		return &comparison.IntComparisonExp
	case COLUMN_FLOAT:
		return &comparison.FloatComparisonExp
	case COLUMN_BOOLEAN:
		return &comparison.BooleanComparisonExp
	case COLUMN_STRING:
		return &comparison.StringComparisonExp
	case COLUMN_DATE:
		return &comparison.DateTimeComparisonExp
	case COLUMN_SIMPLE_JSON:
		return nil
	case COLUMN_SIMPLE_ARRAY:
		return nil
	case COLUMN_JSON_ARRAY:
		return nil
	case COLUMN_ENUM:
		return &comparison.EnumComparisonExp
	}

	panic("No column type: " + column.Type)
}

func (column *ColumnMeta) ToOrderBy() *graphql.Enum {
	switch column.Type {
	case COLUMN_SIMPLE_JSON:
		return nil
	case COLUMN_SIMPLE_ARRAY:
		return nil
	case COLUMN_JSON_ARRAY:
		return nil
	}

	return EnumOrderBy
}
