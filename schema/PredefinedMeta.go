package schema

import "rxdrag.com/entity-engine/model"

var Meta = model.EntityMeta{
	Uuid:       "_META_UUID",
	Name:       "Meta",
	TableName:  "meta",
	EntityType: model.Entity_NORMAL,
	Columns:    []model.ColumnMeta{},
}
