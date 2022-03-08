package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/meta"
)

const (
	BOOLEXP     string = "BoolExp"
	ORDERBY     string = "OrderBy"
	DISTINCTEXP string = "DistinctExp"
)

func OutputFields(entity *meta.Entity) graphql.Fields {
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

func OutputType(entity *meta.Entity) graphql.Output {
	if OutputTypeMap[entity.Name] != nil {
		return *OutputTypeMap[entity.Name]
	}
	var returnValue graphql.Output

	if entity.EntityType == meta.Entity_ENUM {
		return EnumType(entity)
	} else {
		returnValue = graphql.NewObject(
			graphql.ObjectConfig{
				Name:   entity.Name,
				Fields: OutputFields(entity),
			},
		)
	}
	OutputTypeMap[entity.Name] = &returnValue
	return returnValue
}

func WhereExp(entity *meta.Entity) *graphql.InputObject {
	expName := entity.Name + BOOLEXP
	if WhereExpMap[expName] != nil {
		return WhereExpMap[expName]
	}

	andExp := graphql.InputObjectFieldConfig{}
	notExp := graphql.InputObjectFieldConfig{}
	orExp := graphql.InputObjectFieldConfig{}

	fields := graphql.InputObjectConfigFieldMap{
		ARG_AND: &andExp,
		ARG_NOT: &notExp,
		ARG_OR:  &orExp,
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
	WhereExpMap[expName] = boolExp
	return boolExp
}

func OrderBy(entity *meta.Entity) *graphql.InputObject {
	if OrderByMap[entity.Name] != nil {
		return OrderByMap[entity.Name]
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

	OrderByMap[entity.Name] = orderByExp
	return orderByExp
}

func DistinctOnEnum(entity *meta.Entity) *graphql.Enum {
	if DistinctOnEnumMap[entity.Name] != nil {
		return DistinctOnEnumMap[entity.Name]
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
	DistinctOnEnumMap[entity.Name] = entEnum
	return entEnum
}
