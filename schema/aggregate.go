package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/meta"
)

func toAvgFields(entity *meta.EntityMeta) graphql.Fields {
	fields := graphql.Fields{}
	for _, column := range entity.Columns {
		if column.Type == meta.COLUMN_INT || column.Type == meta.COLUMN_FLOAT {
			fields[column.Name] = &graphql.Field{
				Type: column.toType(),
				// Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// 	fmt.Println(p.Context.Value("data"))
				// 	return "world", nil
				// },
			}
		}

	}
	return fields
}

func (entity *EntityMeta) toMaxFields() graphql.Fields {
	fields := graphql.Fields{}
	for _, column := range entity.Columns {
		if column.Type == COLUMN_INT || column.Type == COLUMN_FLOAT {
			fields[column.Name] = &graphql.Field{
				Type: column.toType(),
				// Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// 	fmt.Println(p.Context.Value("data"))
				// 	return "world", nil
				// },
			}
		}

	}
	return fields
}

func (entity *EntityMeta) toMinFields() graphql.Fields {
	fields := graphql.Fields{}
	for _, column := range entity.Columns {
		if column.Type == COLUMN_INT || column.Type == COLUMN_FLOAT {
			fields[column.Name] = &graphql.Field{
				Type: column.toType(),
				// Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// 	fmt.Println(p.Context.Value("data"))
				// 	return "world", nil
				// },
			}
		}

	}
	return fields
}

func (entity *EntityMeta) toSelectFields() graphql.InputObjectConfigFieldMap {
	fields := graphql.InputObjectConfigFieldMap{}
	for _, column := range entity.Columns {
		fields[column.Name] = &graphql.InputObjectFieldConfig{
			Type: column.toType(),
		}
	}

	return fields
}

func (entity *EntityMeta) toStddevFields() graphql.Fields {
	fields := graphql.Fields{}
	for _, column := range entity.Columns {
		if column.Type == COLUMN_INT || column.Type == COLUMN_FLOAT {
			fields[column.Name] = &graphql.Field{
				Type: column.toType(),
				// Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// 	fmt.Println(p.Context.Value("data"))
				// 	return "world", nil
				// },
			}
		}

	}
	return fields
}

func (entity *EntityMeta) toStddevPopFields() graphql.Fields {
	fields := graphql.Fields{}
	for _, column := range entity.Columns {
		if column.Type == COLUMN_INT || column.Type == COLUMN_FLOAT {
			fields[column.Name] = &graphql.Field{
				Type: column.toType(),
				// Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// 	fmt.Println(p.Context.Value("data"))
				// 	return "world", nil
				// },
			}
		}
	}
	return fields
}

func (entity *EntityMeta) toStddevSampFields() graphql.Fields {
	fields := graphql.Fields{}
	for _, column := range entity.Columns {
		if column.Type == COLUMN_INT || column.Type == COLUMN_FLOAT {
			fields[column.Name] = &graphql.Field{
				Type: column.toType(),
				// Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// 	fmt.Println(p.Context.Value("data"))
				// 	return "world", nil
				// },
			}
		}

	}
	return fields
}

func (entity *EntityMeta) toSumFields() graphql.Fields {
	fields := graphql.Fields{}
	for _, column := range entity.Columns {
		if column.Type == COLUMN_INT || column.Type == COLUMN_FLOAT {
			fields[column.Name] = &graphql.Field{
				Type: column.toType(),
				// Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// 	fmt.Println(p.Context.Value("data"))
				// 	return "world", nil
				// },
			}
		}

	}
	return fields
}

func (entity *EntityMeta) toVarPopFields() graphql.Fields {
	fields := graphql.Fields{}
	for _, column := range entity.Columns {
		if column.Type == COLUMN_INT || column.Type == COLUMN_FLOAT {
			fields[column.Name] = &graphql.Field{
				Type: column.toType(),
				// Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// 	fmt.Println(p.Context.Value("data"))
				// 	return "world", nil
				// },
			}
		}

	}
	return fields
}

func (entity *EntityMeta) toVarSampFields() graphql.Fields {
	fields := graphql.Fields{}
	for _, column := range entity.Columns {
		if column.Type == COLUMN_INT || column.Type == COLUMN_FLOAT {
			fields[column.Name] = &graphql.Field{
				Type: column.toType(),
				// Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// 	fmt.Println(p.Context.Value("data"))
				// 	return "world", nil
				// },
			}
		}

	}
	return fields
}

func (entity *EntityMeta) toVarianceFields() graphql.Fields {
	fields := graphql.Fields{}
	for _, column := range entity.Columns {
		if column.Type == COLUMN_INT || column.Type == COLUMN_FLOAT {
			fields[column.Name] = &graphql.Field{
				Type: column.toType(),
				// Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// 	fmt.Println(p.Context.Value("data"))
				// 	return "world", nil
				// },
			}
		}

	}
	return fields
}

func (entity *EntityMeta) createAggregateFields() graphql.Fields {
	fields := graphql.Fields{}
	avgFields := entity.toAvgFields()
	if len(avgFields) > 0 {
		fields["avg"] = &graphql.Field{
			Type: graphql.NewObject(
				graphql.ObjectConfig{
					Name:   entity.Name + "AvgFields",
					Fields: avgFields,
				},
			),
		}
	}

	maxFields := entity.toMaxFields()
	if len(maxFields) > 0 {
		fields["max"] = &graphql.Field{
			Type: graphql.NewObject(
				graphql.ObjectConfig{
					Name:   entity.Name + "MaxFields",
					Fields: maxFields,
				},
			),
		}
	}

	minFields := entity.toMinFields()
	if len(minFields) > 0 {
		fields["min"] = &graphql.Field{
			Type: graphql.NewObject(
				graphql.ObjectConfig{
					Name:   entity.Name + "MinFields",
					Fields: minFields,
				},
			),
		}
	}

	countFields := entity.toSelectFields()
	if len(countFields) > 0 {
		fields["count"] = &graphql.Field{
			Args: graphql.FieldConfigArgument{
				"columns": &graphql.ArgumentConfig{
					Type: graphql.NewInputObject(
						graphql.InputObjectConfig{
							Name:   entity.Name + "SelectColumn",
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

	stddevFields := entity.toStddevFields()
	if len(stddevFields) > 0 {
		fields["stddev"] = &graphql.Field{
			Type: graphql.NewObject(
				graphql.ObjectConfig{
					Name:   entity.Name + "StddevFields",
					Fields: stddevFields,
				},
			),
		}
	}

	stddevPopFields := entity.toStddevPopFields()
	if len(stddevPopFields) > 0 {
		fields["stddevPop"] = &graphql.Field{
			Type: graphql.NewObject(
				graphql.ObjectConfig{
					Name:   entity.Name + "StddevPopFields",
					Fields: stddevPopFields,
				},
			),
		}
	}

	stddevSampFields := entity.toStddevSampFields()
	if len(stddevSampFields) > 0 {
		fields["stddevSamp"] = &graphql.Field{
			Type: graphql.NewObject(
				graphql.ObjectConfig{
					Name:   entity.Name + "StddevSampFields",
					Fields: stddevSampFields,
				},
			),
		}
	}

	sumFields := entity.toSumFields()
	if len(sumFields) > 0 {
		fields["sum"] = &graphql.Field{
			Type: graphql.NewObject(
				graphql.ObjectConfig{
					Name:   entity.Name + "SumFields",
					Fields: sumFields,
				},
			),
		}
	}
	varPopFields := entity.toVarPopFields()
	if len(varPopFields) > 0 {
		fields["varPop"] = &graphql.Field{
			Type: graphql.NewObject(
				graphql.ObjectConfig{
					Name:   entity.Name + "VarPopFields",
					Fields: varPopFields,
				},
			),
		}
	}
	varSampFields := entity.toVarSampFields()
	if len(varSampFields) > 0 {
		fields["varSamp"] = &graphql.Field{
			Type: graphql.NewObject(
				graphql.ObjectConfig{
					Name:   entity.Name + "VarSampFields",
					Fields: varSampFields,
				},
			),
		}
	}
	varianceFields := entity.toVarianceFields()
	if len(varianceFields) > 0 {
		fields["variance"] = &graphql.Field{
			Type: graphql.NewObject(
				graphql.ObjectConfig{
					Name:   entity.Name + "VarianceFields",
					Fields: varianceFields,
				},
			),
		}
	}
	return fields
}

func (entity *EntityMeta) toAggregateType() graphql.Output {
	var returnValue graphql.Output

	fields := graphql.Fields{
		"nodes": &graphql.Field{
			Type: &graphql.List{
				OfType: entity.toOutputType(),
			},
		},
	}

	aggregateFields := entity.createAggregateFields()

	if len(aggregateFields) > 0 {
		fields["aggregate"] = &graphql.Field{
			Type: graphql.NewObject(
				graphql.ObjectConfig{
					Name:   entity.Name + "AggregateFields",
					Fields: aggregateFields,
				},
			),
		}
	}

	returnValue = graphql.NewObject(
		graphql.ObjectConfig{
			Name:   entity.Name + "Aggregate",
			Fields: fields,
		},
	)

	return returnValue
}
