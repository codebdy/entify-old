package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entify/consts"
	"rxdrag.com/entify/model/graph"
	"rxdrag.com/entify/model/meta"
	"rxdrag.com/entify/utils"
)

func AvgFields(node graph.Noder) graphql.Fields {
	fields := graphql.Fields{}
	for _, column := range node.AllAttributes() {
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

func MaxFields(node graph.Noder) graphql.Fields {
	fields := graphql.Fields{}
	for _, attr := range node.AllAttributes() {
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

func MinFields(node graph.Noder) graphql.Fields {
	fields := graphql.Fields{}
	for _, attr := range node.AllAttributes() {
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

func SelectFields(node graph.Noder) graphql.InputObjectConfigFieldMap {
	fields := graphql.InputObjectConfigFieldMap{}
	for _, attr := range node.AllAttributes() {
		fields[attr.Name] = &graphql.InputObjectFieldConfig{
			Type: AttributeType(attr),
		}
	}

	return fields
}

func StddevFields(node graph.Noder) graphql.Fields {
	fields := graphql.Fields{}
	for _, attr := range node.AllAttributes() {
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

func StddevPopFields(node graph.Noder) graphql.Fields {
	fields := graphql.Fields{}
	for _, attr := range node.AllAttributes() {
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

func StddevSampFields(node graph.Noder) graphql.Fields {
	fields := graphql.Fields{}
	for _, attr := range node.AllAttributes() {
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

func SumFields(node graph.Noder) graphql.Fields {
	fields := graphql.Fields{}
	for _, attr := range node.AllAttributes() {
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

func VarPopFields(node graph.Noder) graphql.Fields {
	fields := graphql.Fields{}
	for _, attr := range node.AllAttributes() {
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

func VarSampFields(node graph.Noder) graphql.Fields {
	fields := graphql.Fields{}
	for _, attr := range node.AllAttributes() {
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

func VarianceFields(node graph.Noder) graphql.Fields {
	fields := graphql.Fields{}
	for _, attr := range node.AllAttributes() {
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

func AggregateFields(node graph.Noder) graphql.Fields {
	fields := graphql.Fields{}
	avgFields := AvgFields(node)
	if len(avgFields) > 0 {
		fields["avg"] = &graphql.Field{
			Type: graphql.NewObject(
				graphql.ObjectConfig{
					Name:   node.Name() + "AvgFields",
					Fields: avgFields,
				},
			),
		}
	}

	maxFields := MaxFields(node)
	if len(maxFields) > 0 {
		fields["max"] = &graphql.Field{
			Type: graphql.NewObject(
				graphql.ObjectConfig{
					Name:   node.Name() + "MaxFields",
					Fields: maxFields,
				},
			),
		}
	}

	minFields := MinFields(node)
	if len(minFields) > 0 {
		fields["min"] = &graphql.Field{
			Type: graphql.NewObject(
				graphql.ObjectConfig{
					Name:   node.Name() + "MinFields",
					Fields: minFields,
				},
			),
		}
	}

	countFields := SelectFields(node)
	if len(countFields) > 0 {
		fields["count"] = &graphql.Field{
			Args: graphql.FieldConfigArgument{
				"columns": &graphql.ArgumentConfig{
					Type: graphql.NewInputObject(
						graphql.InputObjectConfig{
							Name:   node.Name() + "SelectColumn",
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

	stddevFields := StddevFields(node)
	if len(stddevFields) > 0 {
		fields["stddev"] = &graphql.Field{
			Type: graphql.NewObject(
				graphql.ObjectConfig{
					Name:   node.Name() + "StddevFields",
					Fields: stddevFields,
				},
			),
		}
	}

	stddevPopFields := StddevPopFields(node)
	if len(stddevPopFields) > 0 {
		fields["stddevPop"] = &graphql.Field{
			Type: graphql.NewObject(
				graphql.ObjectConfig{
					Name:   node.Name() + "StddevPopFields",
					Fields: stddevPopFields,
				},
			),
		}
	}

	stddevSampFields := StddevSampFields(node)
	if len(stddevSampFields) > 0 {
		fields["stddevSamp"] = &graphql.Field{
			Type: graphql.NewObject(
				graphql.ObjectConfig{
					Name:   node.Name() + "StddevSampFields",
					Fields: stddevSampFields,
				},
			),
		}
	}

	sumFields := SumFields(node)
	if len(sumFields) > 0 {
		fields["sum"] = &graphql.Field{
			Type: graphql.NewObject(
				graphql.ObjectConfig{
					Name:   node.Name() + "SumFields",
					Fields: sumFields,
				},
			),
		}
	}
	varPopFields := VarPopFields(node)
	if len(varPopFields) > 0 {
		fields["varPop"] = &graphql.Field{
			Type: graphql.NewObject(
				graphql.ObjectConfig{
					Name:   node.Name() + "VarPopFields",
					Fields: varPopFields,
				},
			),
		}
	}
	varSampFields := VarSampFields(node)
	if len(varSampFields) > 0 {
		fields["varSamp"] = &graphql.Field{
			Type: graphql.NewObject(
				graphql.ObjectConfig{
					Name:   node.Name() + "VarSampFields",
					Fields: varSampFields,
				},
			),
		}
	}
	varianceFields := VarianceFields(node)
	if len(varianceFields) > 0 {
		fields["variance"] = &graphql.Field{
			Type: graphql.NewObject(
				graphql.ObjectConfig{
					Name:   node.Name() + "VarianceFields",
					Fields: varianceFields,
				},
			),
		}
	}
	return fields
}

func AggregateType(node graph.Noder) *graphql.Output {
	name := node.Name() + utils.FirstUpper(consts.AGGREGATE)
	if Cache.AggregateMap[name] != nil {
		return Cache.AggregateMap[name]
	}

	var returnValue graphql.Output

	fields := graphql.Fields{
		consts.NODES: &graphql.Field{
			Type: &graphql.List{
				OfType: Cache.OutputType(node.Name()),
			},
		},
	}

	aggregateFields := AggregateFields(node)

	if len(aggregateFields) > 0 {
		fields[consts.AGGREGATE] = &graphql.Field{
			Type: graphql.NewObject(
				graphql.ObjectConfig{
					Name:   node.Name() + utils.FirstUpper(consts.AGGREGATE) + consts.FIELDS,
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
