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

type EntityRelation struct {
	Name        string
	Relation    *Relation
	OfEntity    *Entity
	TypeEntity  *Entity
	Description string
}

type Entity struct {
	Uuid        string     `json:"uuid"`
	Name        string     `json:"name"`
	TableName   string     `json:"tableName"`
	EntityType  string     `json:"entityType"`
	Columns     []Column   `json:"columns"`
	Eventable   bool       `json:"eventable"`
	Description string     `json:"description"`
	EnumValues  utils.JSON `json:"enumValues"`
	SoftDelete  bool       `json:"softDelete"`
}

func (entity *Entity) ColumnNames() []string {
	names := make([]string, len(entity.Columns))

	for i, column := range entity.Columns {
		names[i] = column.Name
	}
	return names
}

func (entity *Entity) GetColumn(name string) *Column {
	for _, column := range entity.Columns {
		if column.Name == name {
			return &column
		}
	}

	return nil
}

func (entity *Entity) GetTableName() string {
	if (*entity).TableName != "" {
		return (*entity).TableName
	}
	return utils.SnakeString((*entity).Name)
}

func (e *Entity) HasTable() bool {
	return e.EntityType == ENTITY_NORMAL || e.EntityType == "" || e.EntityType == ENTITY_INTERFACE
}

func (r *EntityRelation) IsArray() bool {
	if r.Relation.RelationType == ONE_TO_MANY {
		if r.OfEntity.Uuid == r.Relation.SourceId {
			return true
		}
	} else if r.Relation.RelationType == MANY_TO_ONE {
		if r.OfEntity.Uuid == r.Relation.TargetId {
			return true
		}
	} else if r.Relation.RelationType == MANY_TO_MANY {
		return true
	}
	return false
}

func FindRelationByName(name string, relations []EntityRelation) *EntityRelation {
	for i := range relations {
		if relations[i].Name == name {
			return &relations[i]
		}
	}

	return nil
}
