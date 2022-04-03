package meta

import (
	"rxdrag.com/entity-engine/consts"
)

const (
	META_STATUS_PUBLISHED       string = "published"
	META_STATUS_CANCELLED       string = "cancelled"
	META_STATUS_MIGRATION_ERROR string = "migrationError"
	META_STATUS_ROLLBACK_ERROR  string = "rollbackError"

	META_STATUS_ENUM_UUID string = "META_STATUS_ENUM_UUID"
)

var MetaStatusEnum = ClassMeta{
	Uuid:       META_STATUS_ENUM_UUID,
	Name:       "MetaStatus",
	StereoType: ENUM,
	Attributes: []AttributeMeta{
		{
			Name: META_STATUS_PUBLISHED,
		},
		{
			Name: META_STATUS_CANCELLED,
		},
		{
			Name: META_STATUS_MIGRATION_ERROR,
		},
		{
			Name: META_STATUS_ROLLBACK_ERROR,
		},
	},
}

var MetaClass = ClassMeta{
	Uuid:       "META_ENTITY_UUID",
	Name:       consts.META_ENTITY_NAME,
	InnerId:    1,
	StereoType: CLASSS_ENTITY,
	Attributes: []AttributeMeta{
		{
			Uuid: "META_COLUMN_ID_UUID",
			Type: ID,
			Name: consts.META_ID,
		},
		{
			Uuid: "META_COLUMN_CONTENT_UUID",
			Type: VALUE_OBJECT,
			Name: consts.META_CONTENT,
		},
		{
			Uuid:     "META_COLUMN_STATUS_UUID",
			Type:     ENUM,
			Name:     consts.META_STATUS,
			TypeUuid: META_STATUS_ENUM_UUID,
		},
		{
			Uuid: "META_COLUMN_PUBLISHED_AT_UUID",
			Type: DATE,
			Name: consts.META_PUBLISHEDAT,
		},
		{
			Uuid: "META_COLUMN_CREATED_AT_UUID",
			Type: DATE,
			Name: consts.META_CREATEDAT,
		},
		{
			Uuid: "META_COLUMN_UPDATED_AT_UUID",
			Type: DATE,
			Name: consts.META_UPDATEDAT,
		},
	},
}
