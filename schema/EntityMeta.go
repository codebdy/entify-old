package schema

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

const (
	Entity_NORMAL    string = "Normal"
	Entity_ENUM      string = "Enum"
	Entity_INTERFACE string = "Interface"
)

type EntityMeta struct {
	Uuid        string       `json:"uuid"`
	Name        string       `json:"name"`
	TableName   string       `json:"tableName"`
	EntityType  string       `json:"entityType"`
	Columns     []ColumnMeta `json:"columns"`
	Eventable   bool         `json:"eventable"`
	Description string       `json:"description"`
}

func (entity *EntityMeta) createQueryFields() graphql.Fields {
	fields := graphql.Fields{}
	for _, column := range entity.Columns {
		fields[column.Name] = &graphql.Field{
			Type: column.toOutputType(),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				fmt.Println(p.Context.Value("data"))
				return "world", nil
			},
		}
	}
	return fields
}

func (entity *EntityMeta) toOutputType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name:   entity.Name,
			Fields: entity.createQueryFields(),
		},
	)
}
