package schema

import "rxdrag.com/entity-engine/meta"

var MetaEntity = meta.Entity{
	Uuid:       "META_ENTITY_UUID",
	Name:       "Meta",
	TableName:  "meta",
	EntityType: meta.Entity_NORMAL,
	Columns: []meta.Column{
		{
			Uuid: "META_COLUMN_ID_UUID",
			Type: meta.COLUMN_ID,
			Name: "id",
		},
		{
			Uuid: "META_COLUMN_VERSION_UUID",
			Type: meta.COLUMN_STRING,
			Name: "status",
		},
		{
			Uuid: "META_COLUMN_CONTENT_UUID",
			Type: meta.COLUMN_SIMPLE_JSON,
			Name: "content",
		},
		{
			Uuid: "META_COLUMN_VERSION_UUID",
			Type: meta.COLUMN_STRING,
			Name: "status",
		},
		// {
		// 	Uuid: "META_COLUMN_INT_TEST_UUID",
		// 	Type: meta.COLUMN_INT,
		// 	Name: "int_test",
		// },
		// {
		// 	Uuid: "META_COLUMN_FLOAT_TEST_UUID",
		// 	Type: meta.COLUMN_FLOAT,
		// 	Name: "float_test",
		// },
		// {
		// 	Uuid: "_META_COLUMN_PUBLISHED_UUID",
		// 	Type: COLUMN_BOOLEAN,
		// 	Name: "published",
		// },
		{
			Uuid: "META_COLUMN_PUBLISHED_AT_UUID",
			Type: meta.COLUMN_DATE,
			Name: "publishedAt",
		},
		{
			Uuid: "META_COLUMN_CREATED_AT_UUID",
			Type: meta.COLUMN_DATE,
			Name: "createdAt",
		},
		{
			Uuid: "META_COLUMN_UPDATED_AT_UUID",
			Type: meta.COLUMN_DATE,
			Name: "updatedAt",
		},
	},
}
