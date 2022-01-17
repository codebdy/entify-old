package model

type EntityType int32

const (
	Entity_NORMAL    string = "Normal"
	Entity_ENUM      string = "Enum"
	Entity_ABSTRACT  string = "Abstract"
	Entity_INTERFACE string = "Interface"
)

type EntityMeta struct {
	uuid       string
	name       string
	tableName  string
	entityType string
	columns    []ColumnMeta
	eventable  bool
}
