package schema

import (
	"encoding/json"

	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/meta"
	"rxdrag.com/entity-engine/repository"
	"rxdrag.com/entity-engine/utils"
)

const (
	BOOLEXP     string = "BoolExp"
	ORDERBY     string = "OrderBy"
	DISTINCTEXP string = "DistinctExp"
)

//where表达式缓存，query跟mutation都用
var whereExpMap = make(map[string]*graphql.InputObject)

//类型缓存， query用
var outputTypeMap = make(map[string]*graphql.Output)

var distinctOnEnumMap = make(map[string]*graphql.Enum)

var orderByMap = make(map[string]*graphql.InputObject)

func createQueryFields(entity *meta.EntityMeta) graphql.Fields {
	fields := graphql.Fields{}
	for _, column := range entity.Columns {
		fields[column.Name] = &graphql.Field{
			Type: ColumnType(&column),
			// Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// 	fmt.Println(p.Context.Value("data"))
			// 	return "world", nil
			// },
		}
	}
	return fields
}

func OutputType(entity *meta.EntityMeta) graphql.Output {
	if outputTypeMap[entity.Name] != nil {
		return *outputTypeMap[entity.Name]
	}
	var returnValue graphql.Output

	if entity.EntityType == meta.Entity_ENUM {
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
				Fields: createQueryFields(entity),
			},
		)
	}
	outputTypeMap[entity.Name] = &returnValue
	return returnValue
}

func WhereExp(entity *meta.EntityMeta) *graphql.InputObject {
	expName := entity.Name + BOOLEXP
	if whereExpMap[expName] != nil {
		return whereExpMap[expName]
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
	andExp.Type = &graphql.List{
		OfType: &graphql.NonNull{
			OfType: boolExp,
		},
	}
	notExp.Type = boolExp
	orExp.Type = &graphql.List{
		OfType: &graphql.NonNull{
			OfType: boolExp,
		},
	}

	for _, column := range entity.Columns {
		columnExp := ColumnExp(&column)

		if columnExp != nil {
			fields[column.Name] = columnExp
		}
	}
	whereExpMap[expName] = boolExp
	return boolExp
}

func OrderBy(entity *meta.EntityMeta) *graphql.InputObject {
	if orderByMap[entity.Name] != nil {
		return orderByMap[entity.Name]
	}
	fields := graphql.InputObjectConfigFieldMap{}

	orderByExp := graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name:   entity.Name + ORDERBY,
			Fields: fields,
		},
	)

	for _, column := range entity.Columns {
		columnOrderBy := ColumnOrderBy(&column)

		if columnOrderBy != nil {
			fields[column.Name] = &graphql.InputObjectFieldConfig{Type: columnOrderBy}
		}
	}

	orderByMap[entity.Name] = orderByExp
	return orderByExp
}

func DistinctOnEnum(entity *meta.EntityMeta) *graphql.Enum {
	if distinctOnEnumMap[entity.Name] != nil {
		return distinctOnEnumMap[entity.Name]
	}
	enumValueConfigMap := graphql.EnumValueConfigMap{}
	for _, column := range entity.Columns {
		enumValueConfigMap[column.Name] = &graphql.EnumValueConfig{
			Value: column.Name,
		}
	}

	entEnum := graphql.NewEnum(
		graphql.EnumConfig{
			Name:   entity.Name + DISTINCTEXP,
			Values: enumValueConfigMap,
		},
	)
	distinctOnEnumMap[entity.Name] = entEnum
	return entEnum
}

func AppendToQueryFields(entity *meta.EntityMeta, feilds *graphql.Fields) {
	//如果是枚举
	if entity.EntityType == meta.Entity_ENUM {
		return
	}

	(*feilds)[utils.FirstLower(entity.Name)] = &graphql.Field{
		Type: &graphql.NonNull{
			OfType: &graphql.List{
				OfType: OutputType(entity),
			},
		},
		Args: graphql.FieldConfigArgument{
			"distinctOn": &graphql.ArgumentConfig{
				Type: DistinctOnEnum(entity),
			},
			"limit": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"offset": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"orderBy": &graphql.ArgumentConfig{
				Type: OrderBy(entity),
			},
			"where": &graphql.ArgumentConfig{
				Type: WhereExp(entity),
			},
		},
		Resolve: repository.QueryResolveFn(entity),
	}
	(*feilds)[utils.FirstLower(entity.Name)+"ById"] = &graphql.Field{
		Type: OutputType(entity),
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: repository.QueryResolveFn(entity),
	}

	(*feilds)[utils.FirstLower(entity.Name)+"Aggregate"] = &graphql.Field{
		Type: AggregateType(entity),
		Args: graphql.FieldConfigArgument{
			"distinctOn": &graphql.ArgumentConfig{
				Type: DistinctOnEnum(entity),
			},
			"limit": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"offset": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"orderBy": &graphql.ArgumentConfig{
				Type: OrderBy(entity),
			},
			"where": &graphql.ArgumentConfig{
				Type: WhereExp(entity),
			},
		},
		Resolve: repository.QueryResolveFn(entity),
	}
}
