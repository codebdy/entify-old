package meta

const (
	META_STATUS_PUBLISHED       string = "published"
	META_STATUS_CANCELLED       string = "cancelled"
	META_STATUS_MIGRATION_ERROR string = "migrationError"
	META_STATUS_ROLLBACK_ERROR  string = "rollbackError"

	META_STATUS_ENUM_UUID string = "META_STATUS_ENUM_UUID"
)

var MetaStatusEnum = Entity{
	Uuid:       META_STATUS_ENUM_UUID,
	Name:       "MetaStatus",
	EntityType: Entity_ENUM,
	EnumValues: map[string]interface{}{
		META_STATUS_PUBLISHED: map[string]string{
			"value": META_STATUS_PUBLISHED,
		},
		META_STATUS_CANCELLED: map[string]string{
			"value": META_STATUS_CANCELLED,
		},
		META_STATUS_MIGRATION_ERROR: map[string]string{
			"value": META_STATUS_MIGRATION_ERROR,
		},
		META_STATUS_ROLLBACK_ERROR: map[string]string{
			"value": META_STATUS_ROLLBACK_ERROR,
		},
	},
}

var MetaEntity = Entity{
	Uuid:       "META_ENTITY_UUID",
	Name:       "Meta",
	TableName:  "meta",
	EntityType: Entity_NORMAL,
	Columns: []Column{
		{
			Uuid: "META_COLUMN_ID_UUID",
			Type: COLUMN_ID,
			Name: "id",
		},
		{
			Uuid: "META_COLUMN_VERSION_UUID",
			Type: COLUMN_STRING,
			Name: "status",
		},
		{
			Uuid: "META_COLUMN_CONTENT_UUID",
			Type: COLUMN_SIMPLE_JSON,
			Name: "content",
		},
		{
			Uuid:          "META_COLUMN_VERSION_UUID",
			Type:          COLUMN_ENUM,
			Name:          "status",
			TypeEnityUuid: META_STATUS_ENUM_UUID,
		},
		{
			Uuid: "META_COLUMN_PUBLISHED_AT_UUID",
			Type: COLUMN_DATE,
			Name: "publishedAt",
		},
		{
			Uuid: "META_COLUMN_CREATED_AT_UUID",
			Type: COLUMN_DATE,
			Name: "createdAt",
		},
		{
			Uuid: "META_COLUMN_UPDATED_AT_UUID",
			Type: COLUMN_DATE,
			Name: "updatedAt",
		},
	},
}
