package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entify/consts"
	"rxdrag.com/entify/model/graph"
)

var BooleanComparisonExp = graphql.InputObjectFieldConfig{
	Type: graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name: "BooleanComparisonExp",
			Fields: graphql.InputObjectConfigFieldMap{
				consts.ARG_EQ: &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				consts.ARG_GT: &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				consts.ARG_GTE: &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				consts.ARG_IN: &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(graphql.Boolean),
				},
				consts.ARG_ISNULL: &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				consts.ARG_LT: &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				consts.ARG_LTE: &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				consts.ARG_NOTEQ: &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				consts.ARG_NOTIN: &graphql.InputObjectFieldConfig{
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
				consts.ARG_EQ: &graphql.InputObjectFieldConfig{
					Type: graphql.DateTime,
				},
				consts.ARG_GT: &graphql.InputObjectFieldConfig{
					Type: graphql.DateTime,
				},
				consts.ARG_GTE: &graphql.InputObjectFieldConfig{
					Type: graphql.DateTime,
				},
				consts.ARG_IN: &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(graphql.DateTime),
				},
				consts.ARG_ISNULL: &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				consts.ARG_LT: &graphql.InputObjectFieldConfig{
					Type: graphql.DateTime,
				},
				consts.ARG_LTE: &graphql.InputObjectFieldConfig{
					Type: graphql.DateTime,
				},
				consts.ARG_NOTEQ: &graphql.InputObjectFieldConfig{
					Type: graphql.DateTime,
				},
				consts.ARG_NOTIN: &graphql.InputObjectFieldConfig{
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
				consts.ARG_EQ: &graphql.InputObjectFieldConfig{
					Type: graphql.Float,
				},
				consts.ARG_GT: &graphql.InputObjectFieldConfig{
					Type: graphql.Float,
				},
				consts.ARG_GTE: &graphql.InputObjectFieldConfig{
					Type: graphql.Float,
				},
				consts.ARG_IN: &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(graphql.Float),
				},
				consts.ARG_ISNULL: &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				consts.ARG_LT: &graphql.InputObjectFieldConfig{
					Type: graphql.Float,
				},
				consts.ARG_LTE: &graphql.InputObjectFieldConfig{
					Type: graphql.Float,
				},
				consts.ARG_NOTEQ: &graphql.InputObjectFieldConfig{
					Type: graphql.Float,
				},
				consts.ARG_NOTIN: &graphql.InputObjectFieldConfig{
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
				consts.ARG_EQ: &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				consts.ARG_GT: &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				consts.ARG_GTE: &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				consts.ARG_IN: &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(graphql.Int),
				},
				consts.ARG_ISNULL: &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				consts.ARG_LT: &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				consts.ARG_LTE: &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				consts.ARG_NOTEQ: &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				consts.ARG_NOTIN: &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(graphql.Int),
				},
			},
		},
	),
}

var IdComparisonExp = graphql.InputObjectFieldConfig{
	Type: graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name: "IdComparisonExp",
			Fields: graphql.InputObjectConfigFieldMap{
				consts.ARG_EQ: &graphql.InputObjectFieldConfig{
					Type: graphql.ID,
				},
				consts.ARG_GT: &graphql.InputObjectFieldConfig{
					Type: graphql.ID,
				},
				consts.ARG_GTE: &graphql.InputObjectFieldConfig{
					Type: graphql.ID,
				},
				consts.ARG_IN: &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(graphql.ID),
				},
				consts.ARG_LT: &graphql.InputObjectFieldConfig{
					Type: graphql.ID,
				},
				consts.ARG_LTE: &graphql.InputObjectFieldConfig{
					Type: graphql.ID,
				},
				consts.ARG_NOTEQ: &graphql.InputObjectFieldConfig{
					Type: graphql.ID,
				},
				consts.ARG_NOTIN: &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(graphql.ID),
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
				consts.ARG_EQ: &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				consts.ARG_GT: &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				consts.ARG_GTE: &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				consts.ARG_ILIKE: &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				consts.ARG_IN: &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(graphql.String),
				},
				consts.ARG_IREGEX: &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				consts.ARG_ISNULL: &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				consts.ARG_LIKE: &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				consts.ARG_LT: &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				consts.ARG_LTE: &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				consts.ARG_NOTEQ: &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				consts.ARG_NOTILIKE: &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				consts.ARG_NOTIN: &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(graphql.String),
				},
				consts.ARG_NOTIREGEX: &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				consts.ARG_NOTLIKE: &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				consts.ARG_NOTREGEX: &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				consts.ARG_NOTSIMILAR: &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				consts.ARG_REGEX: &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				consts.ARG_SIMILAR: &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
			},
		},
	),
}

func EnumComparisonExp(attr *graph.Attribute) *graphql.InputObjectFieldConfig {
	enumEntity := attr.EumnType
	if enumEntity == nil {
		panic("Can not find enum entity")
	}
	if Cache.EnumComparisonExpMap[enumEntity.Name] != nil {
		return Cache.EnumComparisonExpMap[enumEntity.Name]
	}
	enumType := Cache.EnumType(enumEntity.Name)
	enumxp := graphql.InputObjectFieldConfig{
		Type: graphql.NewInputObject(
			graphql.InputObjectConfig{
				Name: "EnumComparisonExp",
				Fields: graphql.InputObjectConfigFieldMap{
					consts.ARG_EQ: &graphql.InputObjectFieldConfig{
						Type: enumType,
					},
					consts.ARG_IN: &graphql.InputObjectFieldConfig{
						Type: graphql.NewList(enumType),
					},
					consts.ARG_ISNULL: &graphql.InputObjectFieldConfig{
						Type: graphql.Boolean,
					},
					consts.ARG_NOTEQ: &graphql.InputObjectFieldConfig{
						Type: enumType,
					},
					consts.ARG_NOTIN: &graphql.InputObjectFieldConfig{
						Type: graphql.NewList(enumType),
					},
				},
			},
		),
	}
	Cache.EnumComparisonExpMap[enumEntity.Name] = &enumxp
	return &enumxp
}
