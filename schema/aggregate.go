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

func (entity *EntityMeta) createAggregateFields() graphql.Fields {
	fields := graphql.Fields{}
	fields["avg"] = &graphql.Field{
		Type: graphql.NewObject(
			graphql.ObjectConfig{
				Name:   entity.Name + "AvgFields",
				Fields: entity.toAvgFields(),
			},
		),
	}
	//count(columns: [user_select_column!]distinct: Boolean): Int!
	// "max": user_max_fields
	// "min": user_min_fields
	// "stddev": user_stddev_fields
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

	returnValue = graphql.NewObject(
		graphql.ObjectConfig{
			Name: entity.Name + "Aggregate",
			Fields: graphql.Fields{
				"aggregate": &graphql.Field{
					Type: graphql.NewObject(
						graphql.ObjectConfig{
							Name:   entity.Name + "AggregateFields",
							Fields: entity.createAggregateFields(),
						},
					),
				},
				"nodes": &graphql.Field{
					Type: &graphql.List{
						OfType: entity.toOutputType(),
					},
				},
			},
		},
	)

	return returnValue
}
