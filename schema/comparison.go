package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/meta"
	"rxdrag.com/entity-engine/repository"
)

var BooleanComparisonExp = graphql.InputObjectFieldConfig{
	Type: graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name: "BooleanComparisonExp",
			Fields: graphql.InputObjectConfigFieldMap{
				ARG_EQ: &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				ARG_GT: &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				ARG_GTE: &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				ARG_IN: &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(graphql.Boolean),
				},
				ARG_ISNULL: &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				ARG_LT: &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				ARG_LTE: &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				ARG_NOTEQ: &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				ARG_NOTIN: &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(graphql.Boolean),
				},
			},
		},
	),
}

var DateTimeComparisonExp = graphql.InputObjectFieldConfig{
	Type: graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name: "DateTimeComparisonExp",
			Fields: graphql.InputObjectConfigFieldMap{
				ARG_EQ: &graphql.InputObjectFieldConfig{
					Type: graphql.DateTime,
				},
				ARG_GT: &graphql.InputObjectFieldConfig{
					Type: graphql.DateTime,
				},
				ARG_GTE: &graphql.InputObjectFieldConfig{
					Type: graphql.DateTime,
				},
				ARG_IN: &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(graphql.DateTime),
				},
				ARG_ISNULL: &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				ARG_LT: &graphql.InputObjectFieldConfig{
					Type: graphql.DateTime,
				},
				ARG_LTE: &graphql.InputObjectFieldConfig{
					Type: graphql.DateTime,
				},
				ARG_NOTEQ: &graphql.InputObjectFieldConfig{
					Type: graphql.DateTime,
				},
				ARG_NOTIN: &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(graphql.DateTime),
				},
			},
		},
	),
}

var FloatComparisonExp = graphql.InputObjectFieldConfig{
	Type: graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name: "FloatComparisonExp",
			Fields: graphql.InputObjectConfigFieldMap{
				ARG_EQ: &graphql.InputObjectFieldConfig{
					Type: graphql.Float,
				},
				ARG_GT: &graphql.InputObjectFieldConfig{
					Type: graphql.Float,
				},
				ARG_GTE: &graphql.InputObjectFieldConfig{
					Type: graphql.Float,
				},
				ARG_IN: &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(graphql.Float),
				},
				ARG_ISNULL: &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				ARG_LT: &graphql.InputObjectFieldConfig{
					Type: graphql.Float,
				},
				ARG_LTE: &graphql.InputObjectFieldConfig{
					Type: graphql.Float,
				},
				ARG_NOTEQ: &graphql.InputObjectFieldConfig{
					Type: graphql.Float,
				},
				ARG_NOTIN: &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(graphql.Float),
				},
			},
		},
	),
}

var IntComparisonExp = graphql.InputObjectFieldConfig{
	Type: graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name: "IntComparisonExp",
			Fields: graphql.InputObjectConfigFieldMap{
				ARG_EQ: &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				ARG_GT: &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				ARG_GTE: &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				ARG_IN: &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(graphql.Int),
				},
				ARG_ISNULL: &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				ARG_LT: &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				ARG_LTE: &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				ARG_NOTEQ: &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				ARG_NOTIN: &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(graphql.Int),
				},
			},
		},
	),
}

var StringComparisonExp = graphql.InputObjectFieldConfig{
	Type: graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name: "StringComparisonExp",
			Fields: graphql.InputObjectConfigFieldMap{
				ARG_EQ: &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				ARG_GT: &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				ARG_GTE: &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				ARG_ILIKE: &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				ARG_IN: &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(graphql.String),
				},
				ARG_IREGEX: &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				ARG_ISNULL: &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				ARG_LIKE: &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				ARG_LT: &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				ARG_LTE: &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				ARG_NOTEQ: &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				ARG_NOTILIKE: &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				ARG_NOTIN: &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(graphql.String),
				},
				ARG_NOTIREGEX: &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				ARG_NOTLIKE: &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				ARG_NOTREGEX: &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				ARG_NOTSIMILAR: &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				ARG_REGEX: &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				ARG_SIMILAR: &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
			},
		},
	),
}

func EnumComparisonExp(column *meta.Column) *graphql.InputObjectFieldConfig {
	enumEntity := repository.GetEntityByUuid(column.TypeEnityUuid)
	if enumEntity == nil {
		panic("Can not find enum entity")
	}
	if EnumComparisonExpMap[enumEntity.Name] != nil {
		return EnumComparisonExpMap[enumEntity.Name]
	}
	enumType := EnumType(enumEntity)
	enumxp := graphql.InputObjectFieldConfig{
		Type: graphql.NewInputObject(
			graphql.InputObjectConfig{
				Name: "EnumComparisonExp",
				Fields: graphql.InputObjectConfigFieldMap{
					ARG_EQ: &graphql.InputObjectFieldConfig{
						Type: enumType,
					},
					ARG_IN: &graphql.InputObjectFieldConfig{
						Type: graphql.NewList(enumType),
					},
					ARG_ISNULL: &graphql.InputObjectFieldConfig{
						Type: graphql.Boolean,
					},
					ARG_NOTEQ: &graphql.InputObjectFieldConfig{
						Type: enumType,
					},
					ARG_NOTIN: &graphql.InputObjectFieldConfig{
						Type: graphql.NewList(enumType),
					},
				},
			},
		),
	}
	EnumComparisonExpMap[enumEntity.Name] = &enumxp
	return &enumxp
}
