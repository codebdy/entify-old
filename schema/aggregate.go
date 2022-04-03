package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/model/graph"
	"rxdrag.com/entity-engine/model/meta"
	"rxdrag.com/entity-engine/utils"
)

func AvgFields(entity *graph.Entity) graphql.Fields {
	fields := graphql.Fields{}
	for _, column := range entity.Attributes() {
		if column.Type == meta.INT || column.Type == meta.FLOAT {
			fields[column.Name] = &graphql.Field{
				Type: AttributeType(column),
				// Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// 	fmt.Println(p.Context.Value("data"))
				// 	return "world", nil
				// },
			}
		}

	}
	return fields
}

func MaxFields(entity *graph.Entity) graphql.Fields {
	fields := graphql.Fields{}
	for _, attr := range entity.Attributes() {
		if attr.Type == meta.INT || attr.Type == meta.FLOAT {
			fields[attr.Name] = &graphql.Field{
				Type: AttributeType(attr),
				// Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// 	fmt.Println(p.Context.Value("data"))
				// 	return "world", nil
				// },
			}
		}

	}
	return fields
}

func MinFields(entity *graph.Entity) graphql.Fields {
	fields := graphql.Fields{}
	for _, attr := range entity.Attributes() {
		if attr.Type == meta.INT || attr.Type == meta.FLOAT {
			fields[attr.Name] = &graphql.Field{
				Type: AttributeType(attr),
				// Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// 	fmt.Println(p.Context.Value("data"))
				// 	return "world", nil
				// },
			}
		}

	}
	return fields
}

func SelectFields(entity *graph.Entity) graphql.InputObjectConfigFieldMap {
	fields := graphql.InputObjectConfigFieldMap{}
	for _, attr := range entity.Attributes() {
		fields[attr.Name] = &graphql.InputObjectFieldConfig{
			Type: AttributeType(attr),
		}
	}

	return fields
}

func StddevFields(entity *graph.Entity) graphql.Fields {
	fields := graphql.Fields{}
	for _, attr := range entity.Attributes() {
		if attr.Type == meta.INT || attr.Type == meta.FLOAT {
			fields[attr.Name] = &graphql.Field{
				Type: AttributeType(attr),
				// Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// 	fmt.Println(p.Context.Value("data"))
				// 	return "world", nil
				// },
			}
		}

	}
	return fields
}

func StddevPopFields(entity *graph.Entity) graphql.Fields {
	fields := graphql.Fields{}
	for _, attr := range entity.Attributes() {
		if attr.Type == meta.INT || attr.Type == meta.FLOAT {
			fields[attr.Name] = &graphql.Field{
				Type: AttributeType(attr),
				// Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// 	fmt.Println(p.Context.Value("data"))
				// 	return "world", nil
				// },
			}
		}
	}
	return fields
}

func StddevSampFields(entity *graph.Entity) graphql.Fields {
	fields := graphql.Fields{}
	for _, attr := range entity.Attributes() {
		if attr.Type == meta.INT || attr.Type == meta.FLOAT {
			fields[attr.Name] = &graphql.Field{
				Type: AttributeType(attr),
				// Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// 	fmt.Println(p.Context.Value("data"))
				// 	return "world", nil
				// },
			}
		}

	}
	return fields
}

func SumFields(entity *graph.Entity) graphql.Fields {
	fields := graphql.Fields{}
	for _, attr := range entity.Attributes() {
		if attr.Type == meta.INT || attr.Type == meta.FLOAT {
			fields[attr.Name] = &graphql.Field{
				Type: AttributeType(attr),
				// Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// 	fmt.Println(p.Context.Value("data"))
				// 	return "world", nil
				// },
			}
		}

	}
	return fields
}

func VarPopFields(entity *graph.Entity) graphql.Fields {
	fields := graphql.Fields{}
	for _, attr := range entity.Attributes() {
		if attr.Type == meta.INT || attr.Type == meta.FLOAT {
			fields[attr.Name] = &graphql.Field{
				Type: AttributeType(attr),
				// Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// 	fmt.Println(p.Context.Value("data"))
				// 	return "world", nil
				// },
			}
		}

	}
	return fields
}

func VarSampFields(entity *graph.Entity) graphql.Fields {
	fields := graphql.Fields{}
	for _, attr := range entity.Attributes() {
		if attr.Type == meta.INT || attr.Type == meta.FLOAT {
			fields[attr.Name] = &graphql.Field{
				Type: AttributeType(attr),
				// Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// 	fmt.Println(p.Context.Value("data"))
				// 	return "world", nil
				// },
			}
		}

	}
	return fields
}

func VarianceFields(entity *graph.Entity) graphql.Fields {
	fields := graphql.Fields{}
	for _, attr := range entity.Attributes() {
		if attr.Type == meta.INT || attr.Type == meta.FLOAT {
			fields[attr.Name] = &graphql.Field{
				Type: AttributeType(attr),
				// Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// 	fmt.Println(p.Context.Value("data"))
				// 	return "world", nil
				// },
			}
		}

	}
	return fields
}

func AggregateFields(entity *graph.Entity) graphql.Fields {
	fields := graphql.Fields{}
	avgFields := AvgFields(entity)
	if len(avgFields) > 0 {
		fields["avg"] = &graphql.Field{
			Type: graphql.NewObject(
				graphql.ObjectConfig{
					Name:   entity.Name() + "AvgFields",
					Fields: avgFields,
				},
			),
		}
	}

	maxFields := MaxFields(entity)
	if len(maxFields) > 0 {
		fields["max"] = &graphql.Field{
			Type: graphql.NewObject(
				graphql.ObjectConfig{
					Name:   entity.Name() + "MaxFields",
					Fields: maxFields,
				},
			),
		}
	}

	minFields := MinFields(entity)
	if len(minFields) > 0 {
		fields["min"] = &graphql.Field{
			Type: graphql.NewObject(
				graphql.ObjectConfig{
					Name:   entity.Name() + "MinFields",
					Fields: minFields,
				},
			),
		}
	}

	countFields := SelectFields(entity)
	if len(countFields) > 0 {
		fields["count"] = &graphql.Field{
			Args: graphql.FieldConfigArgument{
				"columns": &graphql.ArgumentConfig{
					Type: graphql.NewInputObject(
						graphql.InputObjectConfig{
							Name:   entity.Name() + "SelectColumn",
							Fields: countFields,
						},
					),
				},
				"distinct": &graphql.ArgumentConfig{
					Type: graphql.Boolean,
				},
			},
			Type: graphql.Int,
		}
	}

	stddevFields := StddevFields(entity)
	if len(stddevFields) > 0 {
		fields["stddev"] = &graphql.Field{
			Type: graphql.NewObject(
				graphql.ObjectConfig{
					Name:   entity.Name() + "StddevFields",
					Fields: stddevFields,
				},
			),
		}
	}

	stddevPopFields := StddevPopFields(entity)
	if len(stddevPopFields) > 0 {
		fields["stddevPop"] = &graphql.Field{
			Type: graphql.NewObject(
				graphql.ObjectConfig{
					Name:   entity.Name() + "StddevPopFields",
					Fields: stddevPopFields,
				},
			),
		}
	}

	stddevSampFields := StddevSampFields(entity)
	if len(stddevSampFields) > 0 {
		fields["stddevSamp"] = &graphql.Field{
			Type: graphql.NewObject(
				graphql.ObjectConfig{
					Name:   entity.Name() + "StddevSampFields",
					Fields: stddevSampFields,
				},
			),
		}
	}

	sumFields := SumFields(entity)
	if len(sumFields) > 0 {
		fields["sum"] = &graphql.Field{
			Type: graphql.NewObject(
				graphql.ObjectConfig{
					Name:   entity.Name() + "SumFields",
					Fields: sumFields,
				},
			),
		}
	}
	varPopFields := VarPopFields(entity)
	if len(varPopFields) > 0 {
		fields["varPop"] = &graphql.Field{
			Type: graphql.NewObject(
				graphql.ObjectConfig{
					Name:   entity.Name() + "VarPopFields",
					Fields: varPopFields,
				},
			),
		}
	}
	varSampFields := VarSampFields(entity)
	if len(varSampFields) > 0 {
		fields["varSamp"] = &graphql.Field{
			Type: graphql.NewObject(
				graphql.ObjectConfig{
					Name:   entity.Name() + "VarSampFields",
					Fields: varSampFields,
				},
			),
		}
	}
	varianceFields := VarianceFields(entity)
	if len(varianceFields) > 0 {
		fields["variance"] = &graphql.Field{
			Type: graphql.NewObject(
				graphql.ObjectConfig{
					Name:   entity.Name() + "VarianceFields",
					Fields: varianceFields,
				},
			),
		}
	}
	return fields
}

func AggregateType(entity *graph.Entity, parents []*graph.Entity) *graphql.Output {
	name := entity.Name() + utils.FirstUpper(consts.AGGREGATE)
	if Cache.AggregateMap[name] != nil {
		return Cache.AggregateMap[name]
	}

	var returnValue graphql.Output

	fields := graphql.Fields{
		consts.NODES: &graphql.Field{
			Type: &graphql.List{
				OfType: Cache.OutputType(entity.Name()),
			},
		},
	}

	aggregateFields := AggregateFields(entity)

	if len(aggregateFields) > 0 {
		fields[consts.AGGREGATE] = &graphql.Field{
			Type: graphql.NewObject(
				graphql.ObjectConfig{
					Name:   entity.Name() + utils.FirstUpper(consts.AGGREGATE) + consts.FIELDS,
					Fields: aggregateFields,
				},
			),
		}
	}

	returnValue = graphql.NewObject(
		graphql.ObjectConfig{
			Name:   name,
			Fields: fields,
		},
	)

	Cache.AggregateMap[name] = &returnValue
	return &returnValue
}
