package meta

import (
	"rxdrag.com/entity-engine/utils"
)

const (
	ENTITY_NORMAL    string = "Normal"
	ENTITY_ENUM      string = "Enum"
	ENTITY_INTERFACE string = "Interface"
	//留待以后版本支持Union
	//ENTITY_UNION string = "Union"
)

type EntityMeta struct {
	Uuid        string       `json:"uuid"`
	Name        string       `json:"name"`
	TableName   string       `json:"tableName"`
	EntityType  string       `json:"entityType"`
	Columns     []ColumnMeta `json:"columns"`
	Eventable   bool         `json:"eventable"`
	Description string       `json:"description"`
	EnumValues  utils.JSON   `json:"enumValues"`
	SoftDelete  bool         `json:"softDelete"`
}

func (entity *EntityMeta) ColumnNames() []string {
	names := make([]string, len(entity.Columns))

	for i, column := range entity.Columns {
		names[i] = column.Name
	}
	return names
}

func (entity *EntityMeta) GetColumn(name string) *ColumnMeta {
	for _, column := range entity.Columns {
		if column.Name == name {
			return &column
		}
	}

	return nil
}

func (entity *EntityMeta) GetTableName() string {
	if (*entity).TableName != "" {
		return (*entity).TableName
	}
	return utils.SnakeString((*entity).Name)
}
