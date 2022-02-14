package schema

import "rxdrag.com/entity-engine/model"

var MetaEntity = model.EntityMeta{
	Uuid:       "_META_UUID",
	Name:       "Meta",
	TableName:  "meta",
	EntityType: model.Entity_NORMAL,
	Columns: []model.ColumnMeta{
		{
			Uuid: "_META_COLUMN_ID_UUID",
			Type: model.COLUMN_NUMBER,
			Name: "id",
		},
		{
			Uuid: "_META_COLUMN_VERSION_UUID",
			Type: model.COLUMN_STRING,
			Name: "version",
		},
		{
			Uuid: "_META_COLUMN_CONTENT_UUID",
			Type: model.COLUMN_STRING,
			Name: "content",
		},
	},
}
