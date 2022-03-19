package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/meta"
	"rxdrag.com/entity-engine/resolve"
	"rxdrag.com/entity-engine/utils"
)

const (
	BOOLEXP     string = "BoolExp"
	ORDERBY     string = "OrderBy"
	DISTINCTEXP string = "DistinctExp"
)

func OutputFields(entity *meta.Entity, parents []*meta.Entity) graphql.Fields {
	fields := graphql.Fields{}
	for _, column := range meta.Metas.EntityAllColumns(entity) {
		fields[column.Name] = &graphql.Field{
			Type: ColumnType(&column),
			// Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// 	fmt.Println(p.Context.Value("data"))
			// 	return "world", nil
			// },
		}
	}

	relations := meta.Metas.EntityAllRelations(entity)
	newParents := append(parents, entity)
	for i := range relations {
		relation := relations[i]
		if !findParent(relation.TypeEntity.Uuid, newParents) {
			relationType := OutputType(relation.TypeEntity, newParents)
			if relation.IsArray() {
				fields[relation.Name] = &graphql.Field{
					Type: &graphql.NonNull{
						OfType: &graphql.List{
							OfType: *relationType,
						},
					},
					// Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// 	fmt.Println(p.Context.Value("data"))
					// 	return "world", nil
					// },
				}
				fields[relation.Name+utils.FirstUpper(consts.AGGREGATE)] = &graphql.Field{
					Type: *AggregateType(relation.TypeEntity, newParents),
					Args: graphql.FieldConfigArgument{
						consts.ARG_DISTINCTON: &graphql.ArgumentConfig{
							Type: DistinctOnEnum(relation.TypeEntity),
						},
						consts.ARG_LIMIT: &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
						consts.ARG_OFFSET: &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
						consts.ARG_ORDERBY: &graphql.ArgumentConfig{
							Type: OrderBy(relation.TypeEntity),
						},
						consts.ARG_WHERE: &graphql.ArgumentConfig{
							Type: WhereExp(relation.TypeEntity, newParents),
						},
					},
					//Resolve: resolve.QueryResolveFn(entity),
				}
			} else {
				fields[relation.Name] = &graphql.Field{
					Type:    *relationType,
					Resolve: resolve.RelationResolveFn(&relation),
				}
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

func WhereExp(entity *meta.Entity, parents []*meta.Entity) *graphql.InputObject {
	expName := entity.Name + parentsSuffix(parents) + BOOLEXP
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
	relations := meta.Metas.EntityRelations(entity)
	newParents := append(parents, entity)
	for i := range relations {
		relation := relations[i]
		if !findParent(relation.TypeEntity.Uuid, newParents) {
			fields[relation.Name] = &graphql.InputObjectFieldConfig{
				Type: WhereExp(relation.TypeEntity, newParents),
			}
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
	for i := range parents {
		parent := parents[i]
		suffix = consts.OF + parent.Name + suffix
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
