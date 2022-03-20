package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/meta"
)

const (
	BOOLEXP     string = "BoolExp"
	ORDERBY     string = "OrderBy"
	DISTINCTEXP string = "DistinctExp"
)

func makeWhereExp(entity *meta.Entity) *graphql.InputObject {
	expName := entity.Name + BOOLEXP
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

	columns := meta.Metas.EntityAllColumns(entity)

	for i := range columns {
		column := &columns[i]
		columnExp := ColumnExp(column)

		if columnExp != nil {
			fields[column.Name] = columnExp
		}
	}
	return boolExp
}

func makeOrderBy(entity *meta.Entity) *graphql.InputObject {
	fields := graphql.InputObjectConfigFieldMap{}

	orderByExp := graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name:   entity.Name + ORDERBY,
			Fields: fields,
		},
	)

	columns := meta.Metas.EntityAllColumns(entity)
	for i := range columns {
		column := &columns[i]
		columnOrderBy := ColumnOrderBy(column)
		if columnOrderBy != nil {
			fields[column.Name] = &graphql.InputObjectFieldConfig{Type: columnOrderBy}
		}
	}
	return orderByExp
}

func makeDistinctOnEnum(entity *meta.Entity) *graphql.Enum {
	enumValueConfigMap := graphql.EnumValueConfigMap{}
	columns := meta.Metas.EntityAllColumns(entity)
	for i := range columns {
		column := &columns[i]
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
	return entEnum
}
