package meta

import (
	"rxdrag.com/entify/consts"
)

const (
	META_STATUS_PUBLISHED       string = "published"
	META_STATUS_CANCELLED       string = "cancelled"
	META_STATUS_MIGRATION_ERROR string = "migrationError"
	META_STATUS_ROLLBACK_ERROR  string = "rollbackError"
	META_STATUS_ENUM_UUID       string = "META_STATUS_ENUM_UUID"

	META_ABILITY_TYPE_CREATE    string = "create"
	META_ABILITY_TYPE_READ      string = "read"
	META_ABILITY_TYPE_UPDATE    string = "update"
	META_ABILITY_TYPE_DELETE    string = "delete"
	META_ABILITY_TYPE_ENUM_UUID string = "META_ABILITY_TYPE_ENUM_UUID"
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
	Root:       true,
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

var EntityAuthSettingsClass = ClassMeta{
	Name:    "EntityAuthSettings",
	Uuid:    "META_ENTITY_AUTH_SETTINGS_UUID",
	InnerId: 2,
	Root:    true,
	System:  true,
	Attributes: []AttributeMeta{
		{
			Name:    consts.ID,
			Type:    ID,
			Uuid:    "RX_ENTITY_AUTH_SETTINGS_ID_UUID",
			Primary: true,
			System:  true,
		},
		{
			Name:   "entityUuid",
			Type:   "String",
			Uuid:   "RX_ENTITY_AUTH_SETTINGS_ENTITY_UUID_UUID",
			System: true,
			Unique: true,
		},
		{
			Name:   "expand",
			Type:   "Boolean",
			Uuid:   "RX_ENTITY_AUTH_SETTINGS_EXPAND_UUID",
			System: true,
		},
	},
	StereoType: "Entity",
}

var AbilityTypeEnum = ClassMeta{
	Uuid:       META_ABILITY_TYPE_ENUM_UUID,
	Name:       "AbilityType",
	StereoType: ENUM,
	Attributes: []AttributeMeta{
		{
			Name: META_ABILITY_TYPE_CREATE,
		},
		{
			Name: META_ABILITY_TYPE_READ,
		},
		{
			Name: META_ABILITY_TYPE_UPDATE,
		},
		{
			Name: META_ABILITY_TYPE_DELETE,
		},
	},
}

var AbilityClass = ClassMeta{
	Name:    "Ability",
	Uuid:    "META_ABILITY_UUID",
	InnerId: 3,
	Root:    true,
	System:  true,
	Attributes: []AttributeMeta{
		{
			Name:    consts.ID,
			Type:    ID,
			Uuid:    "RX_ABILITY_ID_UUID",
			Primary: true,
			System:  true,
		},
		{
			Name:   "entityUuid",
			Type:   "String",
			Uuid:   "RX_ABILITY_ENTITY_UUID_UUID",
			System: true,
		},
		{
			Name:   "columnUuid",
			Type:   "String",
			Uuid:   "RX_ABILITY_COLUMN_UUID_UUID",
			System: true,
		},
		{
			Name:   "can",
			Type:   "Boolean",
			Uuid:   "RX_ABILITY_CAN_UUID",
			System: true,
		},
		{
			Name:   "expression",
			Type:   "String",
			Uuid:   "RX_ABILITY_EXPRESSION_UUID",
			Length: 2000,
			System: true,
		},
		{
			Name:     "abilityType",
			Type:     ENUM,
			Uuid:     "RX_ABILITY_ABILITYTYPE_UUID",
			System:   true,
			TypeUuid: META_ABILITY_TYPE_ENUM_UUID,
		},
		{
			Name:   "roleId",
			Type:   ID,
			Uuid:   "RX_ABILITY_ROLE_ID_UUID",
			System: true,
		},
	},
	StereoType: "Entity",
}
