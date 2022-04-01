package model

import (
	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/oldmeta"
)

const (
	META_STATUS_PUBLISHED       string = "published"
	META_STATUS_CANCELLED       string = "cancelled"
	META_STATUS_MIGRATION_ERROR string = "migrationError"
	META_STATUS_ROLLBACK_ERROR  string = "rollbackError"

	META_STATUS_ENUM_UUID string = "META_STATUS_ENUM_UUID"
)

var MetaStatusEnum = oldmeta.EntityMeta{
	Uuid:       META_STATUS_ENUM_UUID,
	Name:       "MetaStatus",
	EntityType: oldmeta.ENTITY_ENUM,
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

var MetaEntity = oldmeta.EntityMeta{
	Uuid:       "META_ENTITY_UUID",
	Name:       consts.META_ENTITY_NAME,
	TableName:  "meta",
	InnerId:    1,
	EntityType: oldmeta.ENTITY_NORMAL,
	Columns: []oldmeta.ColumnMeta{
		{
			Uuid: "META_COLUMN_ID_UUID",
			Type: oldmeta.COLUMN_ID,
			Name: consts.META_ID,
		},
		{
			Uuid: "META_COLUMN_CONTENT_UUID",
			Type: oldmeta.COLUMN_SIMPLE_JSON,
			Name: consts.META_CONTENT,
		},
		{
			Uuid:     "META_COLUMN_STATUS_UUID",
			Type:     oldmeta.COLUMN_ENUM,
			Name:     consts.META_STATUS,
			EnumUuid: META_STATUS_ENUM_UUID,
		},
		{
			Uuid: "META_COLUMN_PUBLISHED_AT_UUID",
			Type: oldmeta.COLUMN_DATE,
			Name: consts.META_PUBLISHEDAT,
		},
		{
			Uuid: "META_COLUMN_CREATED_AT_UUID",
			Type: oldmeta.COLUMN_DATE,
			Name: consts.META_CREATEDAT,
		},
		{
			Uuid: "META_COLUMN_UPDATED_AT_UUID",
			Type: oldmeta.COLUMN_DATE,
			Name: consts.META_UPDATEDAT,
		},
	},
}
