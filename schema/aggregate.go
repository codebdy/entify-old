package schema

import "github.com/graphql-go/graphql"

func (entity *EntityMeta) toAvgFields() graphql.Fields {
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
	//count(columns: [user_select_column!]distinct: Boolean): Int!

	// "stddevPop": user_stddev_pop_fields
	// "stddevSamp": user_stddev_samp_fields
	// "sum": user_sum_fields
	// "varPop": user_var_pop_fields
	// "varSamp": user_var_samp_fields
	// "variance": user_variance_fields
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
