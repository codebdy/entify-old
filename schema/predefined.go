package schema

var MetaEntity = EntityMeta{
	Uuid:       "_META_ENTITY_UUID",
	Name:       "_meta",
	TableName:  "_meta",
	EntityType: Entity_NORMAL,
	Columns: []ColumnMeta{
		{
			Uuid: "_META_COLUMN_ID_UUID",
			Type: COLUMN_NUMBER,
			Name: "id",
		},
		{
			Uuid: "_META_COLUMN_VERSION_UUID",
			Type: COLUMN_STRING,
			Name: "version",
		},
		{
			Uuid: "_META_COLUMN_CONTENT_UUID",
			Type: COLUMN_STRING,
			Name: "content",
		},
		// {
		// 	Uuid: "_META_COLUMN_PUBLISHED_UUID",
		// 	Type: COLUMN_BOOLEAN,
		// 	Name: "published",
		// },
		{
			Uuid: "_META_COLUMN_PUBLISHED_AT_UUID",
			Type: COLUMN_DATE,
			Name: "publishedAt",
		},
		{
			Uuid: "_META_COLUMN_CREATED_AT_UUID",
			Type: COLUMN_DATE,
			Name: "createdAt",
		},
		{
			Uuid: "_META_COLUMN_UPDATED_AT_UUID",
			Type: COLUMN_DATE,
			Name: "updatedAt",
		},
	},
}
