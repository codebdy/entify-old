package meta

import "rxdrag.com/entity-engine/consts"

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
	EntityType: ENTITY_ENUM,
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
	Name:       consts.META_ENTITY_NAME,
	TableName:  "meta",
	EntityType: ENTITY_NORMAL,
	Columns: []Column{
		{
			Uuid: "META_COLUMN_ID_UUID",
			Type: COLUMN_ID,
			Name: consts.META_ID,
		},
		{
			Uuid: "META_COLUMN_CONTENT_UUID",
			Type: COLUMN_SIMPLE_JSON,
			Name: consts.META_CONTENT,
		},
		{
			Uuid:          "META_COLUMN_STATUS_UUID",
			Type:          COLUMN_ENUM,
			Name:          consts.META_STATUS,
			TypeEnityUuid: META_STATUS_ENUM_UUID,
		},
		{
			Uuid: "META_COLUMN_PUBLISHED_AT_UUID",
			Type: COLUMN_DATE,
			Name: consts.META_PUBLISHEDAT,
		},
		{
			Uuid: "META_COLUMN_CREATED_AT_UUID",
			Type: COLUMN_DATE,
			Name: consts.META_CREATEDAT,
		},
		{
			Uuid: "META_COLUMN_UPDATED_AT_UUID",
			Type: COLUMN_DATE,
			Name: consts.META_UPDATEDAT,
		},
	},
}

var Metas *MetaContent
