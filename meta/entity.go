package meta

import (
	"rxdrag.com/entity-engine/utils"
)

const (
	Entity_NORMAL    string = "Normal"
	Entity_ENUM      string = "Enum"
	Entity_INTERFACE string = "Interface"
)

type Entity struct {
	Uuid        string   `json:"uuid"`
	Name        string   `json:"name"`
	TableName   string   `json:"tableName"`
	EntityType  string   `json:"entityType"`
	Columns     []Column `json:"columns"`
	Eventable   bool     `json:"eventable"`
	Description string   `json:"description"`
	EnumValues  []byte   `json:"enumValues"`
}

func (entity *Entity) GetTableName() string {
	if (*entity).TableName != "" {
		return (*entity).TableName
	}
	return utils.SnakeString((*entity).Name)
}
