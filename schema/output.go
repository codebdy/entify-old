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

func OutputFields(entity *meta.Entity, parents []*meta.Entity) graphql.Fields {
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

	relations := meta.Metas.EntityRelations(entity)
	newParents := append(parents, entity)
	for i := range relations {
		relation := relations[i]
		if !findParent(relation.TypeEntity.Uuid, newParents) {
			fields[relation.Name] = &graphql.Field{
				Type: *OutputType(relation.TypeEntity, newParents),
				// Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// 	fmt.Println(p.Context.Value("data"))
				// 	return "world", nil
				// },
			}
		}
	}
	return fields
}

func OutputType(entity *meta.Entity, parents []*meta.Entity) *graphql.Output {
	name := entity.Name + parentsSuffix(parents)
	if Cache.OutputTypeMap[name] != nil {
		return Cache.OutputTypeMap[name]
	}
	var returnValue graphql.Output

	if entity.EntityType == meta.ENTITY_ENUM {
		returnValue = EnumType(entity)
	} else {
		returnValue = graphql.NewObject(
			graphql.ObjectConfig{
				Name:   name,
				Fields: OutputFields(entity, parents),
			},
		)
	}

	Cache.OutputTypeMap[name] = &returnValue
	return &returnValue
}

func WhereExp(entity *meta.Entity) *graphql.InputObject {
	expName := entity.Name + BOOLEXP
	if Cache.WhereExpMap[expName] != nil {
		return Cache.WhereExpMap[expName]
	}

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

	for _, column := range entity.Columns {
		columnExp := ColumnExp(&column)

		if columnExp != nil {
			fields[column.Name] = columnExp
		}
	}
	Cache.WhereExpMap[expName] = boolExp
	return boolExp
}

func OrderBy(entity *meta.Entity) *graphql.InputObject {
	if Cache.OrderByMap[entity.Name] != nil {
		return Cache.OrderByMap[entity.Name]
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
			Name:   entity.Name + DISTINCTEXP,
			Values: enumValueConfigMap,
		},
	)
	Cache.DistinctOnEnumMap[entity.Name] = entEnum
	return entEnum
}

func parentsSuffix(parents []*meta.Entity) string {
	suffix := ""
	for i := len(parents) - 1; i >= 0; i-- {
		parent := parents[i]
		if suffix != "" {
			suffix = consts.CONST_OF + parent.Name + consts.CONST_OF + suffix
		} else {
			suffix = consts.CONST_OF + parent.Name
		}
	}
	return suffix
}

func findParent(uuid string, parents []*meta.Entity) bool {
	for _, entity := range parents {
		if entity.Uuid == uuid {
			return true
		}
	}
	return false
}
