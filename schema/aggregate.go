package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entify/consts"
	"rxdrag.com/entify/model/graph"
	"rxdrag.com/entify/model/meta"
)

func AggregateInterfaceType(intf *graph.Interface) *graphql.Object {
	return aggregateType(intf.Name(), intf.AggregateName(), intf.AllAttributes())
}

func AggregateEntityType(entity *graph.Entity) *graphql.Object {
	return aggregateType(entity.Name(), entity.AggregateName(), entity.AllAttributes())
}

func AggregatePartialType(partial *graph.Partial) *graphql.Object {
	return aggregateType(partial.Name(), partial.AggregateName(), partial.AllAttributes())
}

func avgFields(attrs []*graph.Attribute) graphql.Fields {
	fields := graphql.Fields{}
	for _, column := range attrs {
		if column.Type == meta.INT || column.Type == meta.FLOAT {
			fields[column.Name] = &graphql.Field{
				Type: PropertyType(column),
			}
		}

	}
	return fields
}

func maxFields(attrs []*graph.Attribute) graphql.Fields {
	fields := graphql.Fields{}
	for _, attr := range attrs {
		if attr.Type == meta.INT || attr.Type == meta.FLOAT {
			fields[attr.Name] = &graphql.Field{
				Type: PropertyType(attr),
			}
		}

	}
	return fields
}

func minFields(attrs []*graph.Attribute) graphql.Fields {
	fields := graphql.Fields{}
	for _, attr := range attrs {
		if attr.Type == meta.INT || attr.Type == meta.FLOAT {
			fields[attr.Name] = &graphql.Field{
				Type: PropertyType(attr),
				// Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// 	fmt.Println(p.Context.Value("data"))
				// 	return "world", nil
				// },
			}
		}

	}
	return fields
}

func selectFields(attrs []*graph.Attribute) graphql.InputObjectConfigFieldMap {
	fields := graphql.InputObjectConfigFieldMap{}
	for _, attr := range attrs {
		if attr.Type != meta.FILE {
			fields[attr.Name] = &graphql.InputObjectFieldConfig{
				Type: InputPropertyType(attr),
			}
		}
	}

	return fields
}

func stddevFields(attrs []*graph.Attribute) graphql.Fields {
	fields := graphql.Fields{}
	for _, attr := range attrs {
		if attr.Type == meta.INT || attr.Type == meta.FLOAT {
			fields[attr.Name] = &graphql.Field{
				Type: PropertyType(attr),
				// Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// 	fmt.Println(p.Context.Value("data"))
				// 	return "world", nil
				// },
			}
		}

	}
	return fields
}

func stddevPopFields(attrs []*graph.Attribute) graphql.Fields {
	fields := graphql.Fields{}
	for _, attr := range attrs {
		if attr.Type == meta.INT || attr.Type == meta.FLOAT {
			fields[attr.Name] = &graphql.Field{
				Type: PropertyType(attr),
				// Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// 	fmt.Println(p.Context.Value("data"))
				// 	return "world", nil
				// },
			}
		}
	}
	return fields
}

func stddevSampFields(attrs []*graph.Attribute) graphql.Fields {
	fields := graphql.Fields{}
	for _, attr := range attrs {
		if attr.Type == meta.INT || attr.Type == meta.FLOAT {
			fields[attr.Name] = &graphql.Field{
				Type: PropertyType(attr),
				// Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// 	fmt.Println(p.Context.Value("data"))
				// 	return "world", nil
				// },
			}
		}

	}
	return fields
}

func sumFields(attrs []*graph.Attribute) graphql.Fields {
	fields := graphql.Fields{}
	for _, attr := range attrs {
		if attr.Type == meta.INT || attr.Type == meta.FLOAT {
			fields[attr.Name] = &graphql.Field{
				Type: PropertyType(attr),
				// Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// 	fmt.Println(p.Context.Value("data"))
				// 	return "world", nil
				// },
			}
		}

	}
	return fields
}

func varPopFields(attrs []*graph.Attribute) graphql.Fields {
	fields := graphql.Fields{}
	for _, attr := range attrs {
		if attr.Type == meta.INT || attr.Type == meta.FLOAT {
			fields[attr.Name] = &graphql.Field{
				Type: PropertyType(attr),
				// Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// 	fmt.Println(p.Context.Value("data"))
				// 	return "world", nil
				// },
			}
		}

	}
	return fields
}

func varSampFields(attrs []*graph.Attribute) graphql.Fields {
	fields := graphql.Fields{}
	for _, attr := range attrs {
		if attr.Type == meta.INT || attr.Type == meta.FLOAT {
			fields[attr.Name] = &graphql.Field{
				Type: PropertyType(attr),
				// Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// 	fmt.Println(p.Context.Value("data"))
				// 	return "world", nil
				// },
			}
		}

	}
	return fields
}

func varianceFields(attrs []*graph.Attribute) graphql.Fields {
	fields := graphql.Fields{}
	for _, attr := range attrs {
		if attr.Type == meta.INT || attr.Type == meta.FLOAT {
			fields[attr.Name] = &graphql.Field{
				Type: PropertyType(attr),
				// Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// 	fmt.Println(p.Context.Value("data"))
				// 	return "world", nil
				// },
			}
		}

	}
	return fields
}

func aggregateFields(name string, attrs []*graph.Attribute) graphql.Fields {
	fields := graphql.Fields{}
	avgFields := avgFields(attrs)
	if len(avgFields) > 0 {
		fields["avg"] = &graphql.Field{
			Type: graphql.NewObject(
				graphql.ObjectConfig{
					Name:   name + "AvgFields",
					Fields: avgFields,
				},
			),
		}
	}

	maxFields := maxFields(attrs)
	if len(maxFields) > 0 {
		fields["max"] = &graphql.Field{
			Type: graphql.NewObject(
				graphql.ObjectConfig{
					Name:   name + "MaxFields",
					Fields: maxFields,
				},
			),
		}
	}

	minFields := minFields(attrs)
	if len(minFields) > 0 {
		fields["min"] = &graphql.Field{
			Type: graphql.NewObject(
				graphql.ObjectConfig{
					Name:   name + "MinFields",
					Fields: minFields,
				},
			),
		}
	}

	countFields := selectFields(attrs)
	if len(countFields) > 0 {
		selectColumnName := name + "SelectColumn"
		selectColumn := graphql.NewInputObject(
			graphql.InputObjectConfig{
				Name:   selectColumnName,
				Fields: countFields,
			},
		)
		Cache.SelectColumnsMap[selectColumnName] = selectColumn
		fields[consts.ARG_COUNT] = &graphql.Field{
			Args: graphql.FieldConfigArgument{
				consts.ARG_COLUMNS: &graphql.ArgumentConfig{
					Type: selectColumn,
				},
				consts.ARG_DISTINCT: &graphql.ArgumentConfig{
					Type: graphql.Boolean,
				},
			},
			Type: graphql.Int,
		}
	}

	stddevFields := stddevFields(attrs)
	if len(stddevFields) > 0 {
		fields["stddev"] = &graphql.Field{
			Type: graphql.NewObject(
				graphql.ObjectConfig{
					Name:   name + "StddevFields",
					Fields: stddevFields,
				},
			),
		}
	}

	stddevPopFields := stddevPopFields(attrs)
	if len(stddevPopFields) > 0 {
		fields["stddevPop"] = &graphql.Field{
			Type: graphql.NewObject(
				graphql.ObjectConfig{
					Name:   name + "StddevPopFields",
					Fields: stddevPopFields,
				},
			),
		}
	}

	stddevSampFields := stddevSampFields(attrs)
	if len(stddevSampFields) > 0 {
		fields["stddevSamp"] = &graphql.Field{
			Type: graphql.NewObject(
				graphql.ObjectConfig{
					Name:   name + "StddevSampFields",
					Fields: stddevSampFields,
				},
			),
		}
	}

	sumFields := sumFields(attrs)
	if len(sumFields) > 0 {
		fields["sum"] = &graphql.Field{
			Type: graphql.NewObject(
				graphql.ObjectConfig{
					Name:   name + "SumFields",
					Fields: sumFields,
				},
			),
		}
	}
	varPopFields := varPopFields(attrs)
	if len(varPopFields) > 0 {
		fields["varPop"] = &graphql.Field{
			Type: graphql.NewObject(
				graphql.ObjectConfig{
					Name:   name + "VarPopFields",
					Fields: varPopFields,
				},
			),
		}
	}
	varSampFields := varSampFields(attrs)
	if len(varSampFields) > 0 {
		fields["varSamp"] = &graphql.Field{
			Type: graphql.NewObject(
				graphql.ObjectConfig{
					Name:   name + "VarSampFields",
					Fields: varSampFields,
				},
			),
		}
	}
	varianceFields := varianceFields(attrs)
	if len(varianceFields) > 0 {
		fields["variance"] = &graphql.Field{
			Type: graphql.NewObject(
				graphql.ObjectConfig{
					Name:   name + "VarianceFields",
					Fields: varianceFields,
				},
			),
		}
	}
	return fields
}

func aggregateType(name string, aggregateName string, attrs []*graph.Attribute) *graphql.Object {
	if Cache.AggregateMap[aggregateName] != nil {
		return Cache.AggregateMap[aggregateName]
	}

	var returnValue *graphql.Object

	fields := graphql.Fields{
		consts.NODES: &graphql.Field{
			Type: &graphql.List{
				OfType: Cache.OutputType(name),
			},
		},
	}

	aggregateFields := aggregateFields(name, attrs)

	if len(aggregateFields) > 0 {
		fields[consts.AGGREGATE] = &graphql.Field{
			Type: graphql.NewObject(
				graphql.ObjectConfig{
					Name:   aggregateName + consts.FIELDS,
					Fields: aggregateFields,
				},
			),
		}
	}

	returnValue = graphql.NewObject(
		graphql.ObjectConfig{
			Name:   aggregateName,
			Fields: fields,
		},
	)

	Cache.AggregateMap[aggregateName] = returnValue
	return returnValue
}
