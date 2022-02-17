package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/schema/comparisons"
)

const (
	COLUMN_NUMBER       string = "Number"
	COLUMN_BOOLEAN      string = "Boolean"
	COLUMN_STRING       string = "String"
	COLUMN_TEXT         string = "Text"
	COLUMN_MEDIUM_TEXT  string = "MediumText"
	COLUMN_LONG_TEXT    string = "LongText"
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

func (column *ColumnMeta) toOutputType() graphql.Output {
	switch column.Type {
	case COLUMN_NUMBER:
		return graphql.Int
	case COLUMN_BOOLEAN:
		return graphql.Boolean
	case COLUMN_STRING:
		return graphql.String
	case COLUMN_TEXT:
		return graphql.String
	case COLUMN_MEDIUM_TEXT:
		return graphql.String
	case COLUMN_LONG_TEXT:
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
	case COLUMN_NUMBER:
		return &comparisons.IntComparisonExp
	case COLUMN_BOOLEAN:
		return &comparisons.BooleanComparisonExp
	case COLUMN_STRING:
		return &comparisons.StringComparisonExp
	case COLUMN_TEXT:
		return &comparisons.StringComparisonExp
	case COLUMN_MEDIUM_TEXT:
		return &comparisons.StringComparisonExp
	case COLUMN_LONG_TEXT:
		return &comparisons.StringComparisonExp
	case COLUMN_DATE:
		return &comparisons.DateTimeComparisonExp
	case COLUMN_SIMPLE_JSON:
		return nil
	case COLUMN_SIMPLE_ARRAY:
		return nil
	case COLUMN_JSON_ARRAY:
		return nil
	case COLUMN_ENUM:
		return &comparisons.EnumComparisonExp
	}

	panic("No column type: " + column.Type)
}
