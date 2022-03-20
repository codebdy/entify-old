package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/meta"
)

func makeWhereExp(entity *meta.Entity) *graphql.InputObject {
	expName := entity.WhereExpName()
	andExp := graphql.InputObjectFieldConfig{}
	notExp := graphql.InputObjectFieldConfig{}
	orExp := graphql.InputObjectFieldConfig{}

	fields := graphql.InputObjectConfigFieldMap{
		consts.ARG_AND: &andExp,
		consts.ARG_NOT: &notExp,
		consts.ARG_OR:  &orExp,
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

	for _, column := range meta.Metas.EntityAllColumns(entity) {
		columnExp := ColumnExp(&column)

		if columnExp != nil {
			fields[column.Name] = columnExp
		}
	}
	return boolExp
}

func OrderBy(entity *meta.Entity) *graphql.InputObject {
	if Cache.OrderByMap[entity.Name] != nil {
		return Cache.OrderByMap[entity.Name]
	}
	fields := graphql.InputObjectConfigFieldMap{}

	orderByExp := graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name:   entity.OrderByName(),
			Fields: fields,
		},
	)

	for _, column := range entity.Columns {
		columnOrderBy := ColumnOrderBy(&column)

		if columnOrderBy != nil {
			fields[column.Name] = &graphql.InputObjectFieldConfig{Type: columnOrderBy}
		}
	}

	Cache.OrderByMap[entity.Name] = orderByExp
	return orderByExp
}

func DistinctOnEnum(entity *meta.Entity) *graphql.Enum {
	if Cache.DistinctOnEnumMap[entity.Name] != nil {
		return Cache.DistinctOnEnumMap[entity.Name]
	}
	enumValueConfigMap := graphql.EnumValueConfigMap{}
	for _, column := range entity.Columns {
		enumValueConfigMap[column.Name] = &graphql.EnumValueConfig{
			Value: column.Name,
		}
	}

	entEnum := graphql.NewEnum(
		graphql.EnumConfig{
			Name:   entity.DistinctExpname(),
			Values: enumValueConfigMap,
		},
	)
	Cache.DistinctOnEnumMap[entity.Name] = entEnum
	return entEnum
}

func findParent(uuid string, parents []*meta.Entity) bool {
	for _, entity := range parents {
		if entity.Uuid == uuid {
			return true
		}
	}
	return false
}
