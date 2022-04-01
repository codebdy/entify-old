package oldmeta

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
	InnerId     uint64       `json:"innerId"`
	Name        string       `json:"name"`
	TableName   string       `json:"tableName"`
	EntityType  string       `json:"entityType"`
	Columns     []ColumnMeta `json:"columns"`
	Eventable   bool         `json:"eventable"`
	Description string       `json:"description"`
	EnumValues  utils.JSON   `json:"enumValues"`
	SoftDelete  bool         `json:"softDelete"`
}
