package schema

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/schema/comparisons"
	"rxdrag.com/entity-engine/utils"
)

const (
	DISTINCTEXP string = "DistinctExp"
	BOOLEXP     string = "BoolExp"
)

func createFieldType(column *ColumnMeta) graphql.Output {
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

func createEnumEntityType() {

}

func createFieldExp(column *ColumnMeta) *graphql.InputObjectFieldConfig {
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
		// case COLUMN_SIMPLE_JSON:
		// 	return graphql.NewScalar(graphql.ScalarConfig{Name: "JSON"})
		// case COLUMN_SIMPLE_ARRAY:
		// 	return graphql.NewScalar(graphql.ScalarConfig{Name: "JSON"})
		// case COLUMN_JSON_ARRAY:
		// 	return graphql.NewScalar(graphql.ScalarConfig{Name: "JSON"})
		// case COLUMN_ENUM:
		// 	return graphql.InputObjectFieldConfig{
		// 		Type:graphql.NewEnum()
		// 	}
	}

	panic("No column type: " + column.Type)
}

func appendEntityFieldExps(entity *EntityMeta, fieldsMap *graphql.InputObjectConfigFieldMap) {

}

func CreateEntityFields(entity *EntityMeta) *graphql.Fields {
	fields := &graphql.Fields{}
	for _, column := range entity.Columns {
		(*fields)[column.Name] = &graphql.Field{
			Type: createFieldType(&column),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				fmt.Println(p.Context.Value("data"))
				return "world", nil
			},
		}
	}
	return fields
}

func AppendEntityToQueryFields(entity *EntityMeta, feilds *graphql.Fields) {
	metaType := graphql.NewObject(graphql.ObjectConfig{Name: entity.Name, Fields: *CreateEntityFields(entity)})
	metaDistinctType := graphql.NewEnum(graphql.EnumConfig{
		Name: entity.Name + DISTINCTEXP,
		Values: graphql.EnumValueConfigMap{
			"name": &graphql.EnumValueConfig{
				Value: "name",
			},
		},
	})

	andExp := graphql.InputObjectFieldConfig{}
	notExp := graphql.InputObjectFieldConfig{}
	orExp := graphql.InputObjectFieldConfig{}

	metaBoolExp := graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name: entity.Name + BOOLEXP,
			Fields: graphql.InputObjectConfigFieldMap{
				"and": &andExp,
				"not": &notExp,
				"or":  &orExp,
				"id":  &comparisons.IntComparisonExp,
			},
		},
	)
	andExp.Type = metaBoolExp
	notExp.Type = metaBoolExp
	orExp.Type = metaBoolExp

	(*feilds)[utils.FirstLower(entity.Name)] = &graphql.Field{
		Type: metaType,
		Args: graphql.FieldConfigArgument{
			"distinctOn": &graphql.ArgumentConfig{
				Type: metaDistinctType,
			},
			"limit": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"offset": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"orderBy": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"where": &graphql.ArgumentConfig{
				Type: metaBoolExp,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			fmt.Println(p.Context.Value("data"))
			return "world", nil
		},
	}
}
