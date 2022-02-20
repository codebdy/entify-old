package schema

import (
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
)

const (
	BOOLEXP     string = "BoolExp"
	ORDERBY     string = "OrderBy"
	DISTINCTEXP string = "DistinctExp"
)

const (
	Entity_NORMAL    string = "Normal"
	Entity_ENUM      string = "Enum"
	Entity_INTERFACE string = "Interface"
)

type EntityMeta struct {
	Uuid        string       `json:"uuid"`
	Name        string       `json:"name"`
	TableName   string       `json:"tableName"`
	EntityType  string       `json:"entityType"`
	Columns     []ColumnMeta `json:"columns"`
	Eventable   bool         `json:"eventable"`
	Description string       `json:"description"`
	EnumValues  []byte       `json:"enumValues"`
}

//where表达式缓存，query跟mutation都用
var WhereExpMap = make(map[string]*graphql.InputObject)

//类型缓存， query用
var OutputTypeMap = make(map[string]*graphql.Output)

//mutition类型缓存， mutaion用
var MutationResponseMap = make(map[string]*graphql.Output)

func (entity *EntityMeta) createQueryFields() graphql.Fields {
	fields := graphql.Fields{}
	for _, column := range entity.Columns {
		fields[column.Name] = &graphql.Field{
			Type: column.toOutputType(),
			// Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// 	fmt.Println(p.Context.Value("data"))
			// 	return "world", nil
			// },
		}
	}
	return fields
}

func (entity *EntityMeta) toOutputType() graphql.Output {
	if OutputTypeMap[entity.Name] != nil {
		return *OutputTypeMap[entity.Name]
	}
	var returnValue graphql.Output

	if entity.EntityType == Entity_ENUM {
		enumValues := make(map[string]interface{})
		json.Unmarshal(entity.EnumValues, &enumValues)
		enumValueConfigMap := graphql.EnumValueConfigMap{}
		for enumName, enumValue := range enumValues {
			var value, ok = enumValue.(string)
			if !ok {
				value = enumValue.(map[string]string)["value"]
			}
			enumValueConfigMap[enumName] = &graphql.EnumValueConfig{
				Value: value,
			}
		}
		returnValue = graphql.NewEnum(
			graphql.EnumConfig{
				Name:   entity.Name,
				Values: enumValueConfigMap,
			},
		)
	} else {
		returnValue = graphql.NewObject(
			graphql.ObjectConfig{
				Name:   entity.Name,
				Fields: entity.createQueryFields(),
			},
		)
	}
	OutputTypeMap[entity.Name] = &returnValue
	return returnValue
}

func (entity *EntityMeta) toWhereExp() *graphql.InputObject {
	expName := entity.Name + BOOLEXP
	if WhereExpMap[expName] != nil {
		return WhereExpMap[expName]
	}

	andExp := graphql.InputObjectFieldConfig{}
	notExp := graphql.InputObjectFieldConfig{}
	orExp := graphql.InputObjectFieldConfig{}

	fields := graphql.InputObjectConfigFieldMap{
		"and": &andExp,
		"not": &notExp,
		"or":  &orExp,
	}

	boolExp := graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name:   expName,
			Fields: fields,
		},
	)
	andExp.Type = boolExp
	notExp.Type = boolExp
	orExp.Type = boolExp

	for _, column := range entity.Columns {
		columnExp := column.ToExp()

		if columnExp != nil {
			fields[column.Name] = columnExp
		}
	}
	WhereExpMap[expName] = boolExp
	return boolExp
}

func (entity *EntityMeta) toOrderBy() *graphql.InputObject {
	fields := graphql.InputObjectConfigFieldMap{}

	orderByExp := graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name:   entity.Name + ORDERBY,
			Fields: fields,
		},
	)

	for _, column := range entity.Columns {
		columnOrderBy := column.ToOrderBy()

		if columnOrderBy != nil {
			fields[column.Name] = &graphql.InputObjectFieldConfig{Type: columnOrderBy}
		}
	}

	return orderByExp
}

func (entity *EntityMeta) toDistinctOnEnum() *graphql.Enum {
	enumValueConfigMap := graphql.EnumValueConfigMap{}
	for _, column := range entity.Columns {
		enumValueConfigMap[column.Name] = &graphql.EnumValueConfig{
			Value: column.Name,
		}
	}

	return graphql.NewEnum(
		graphql.EnumConfig{
			Name:   entity.Name + DISTINCTEXP,
			Values: enumValueConfigMap,
		},
	)
}

func (entity *EntityMeta) QueryResolve() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		fmt.Println("Resolve entity:" + entity.Name)
		fmt.Println(p.Args)
		fmt.Println(p.Context.Value("data"))
		rtValue := make(map[string]interface{})
		rtValue["content"] = "content"
		return rtValue, nil
	}
}

func (entity *EntityMeta) toMutationResponseType() graphql.Output {
	if MutationResponseMap[entity.Name] != nil {
		return *MutationResponseMap[entity.Name]
	}
	var returnValue graphql.Output

	returnValue = graphql.NewObject(
		graphql.ObjectConfig{
			Name: entity.Name + "MutationResponse",
			Fields: graphql.Fields{
				"affectedRows": &graphql.Field{
					Type: graphql.Int,
				},
				"returning": &graphql.Field{
					Type: &graphql.NonNull{
						OfType: &graphql.List{
							OfType: entity.toOutputType(),
						},
					},
				},
			},
		},
	)

	MutationResponseMap[entity.Name] = &returnValue
	return returnValue
}
