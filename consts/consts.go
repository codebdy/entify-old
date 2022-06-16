package consts

const LOADERS = "loaders"

const (
	ROOT_QUERY_NAME        = "Query"
	ROOT_MUTATION_NAME     = "Mutation"
	ROOT_SUBSCRIPTION_NAME = "Subscription"
	LOGIN                  = "login"
	LOGIN_NAME             = "loginName"
	PASSWORD               = "password"
	LOGOUT                 = "logout"
	ME                     = "me"
	PUBLISH                = "publish"
	ROLLBACK               = "rollback"
	SYNC_META              = "syncMeta"
	NAME                   = "name"
	INSTALLED              = "entifyInstalled"

	ONE          = "one"
	QUERY        = "query"
	AGGREGATE    = "aggregate"
	FIELDS       = "Fields"
	NODES        = "nodes"
	INPUT        = "Input"
	UPDATE_INPUT = "UpdateInput"
	UPSERT       = "upsert"
	UPSERT_ONE   = "upsertOne"
	DELETE       = "delete"
	BY_ID        = "ById"
	UPDATE       = "update"
	HAS_MANY     = "HasMany"
	HAS_ONE      = "HasOne"
	ENTITY       = "Entity"

	ARG_DISTINCTON string = "distinctOn"
	ARG_LIMIT      string = "limit"
	ARG_OFFSET     string = "offset"
	ARG_ORDERBY    string = "orderBy"
	ARG_WHERE      string = "where"

	ARG_ADD     string = "add"
	ARG_DELETE  string = "delete"
	ARG_UPDATE  string = "update"
	ARG_SYNC    string = "sync"
	ARG_CASCADE string = "cascade"

	ARG_AND string = "_and"
	ARG_NOT string = "_not"
	ARG_OR  string = "_or"
)

const (
	ARG_EQ     string = "_eq"
	ARG_GT     string = "_gt"
	ARG_GTE    string = "_gte"
	ARG_IN     string = "_in"
	ARG_ISNULL string = "_isNull"
	ARG_LT     string = "_lt"
	ARG_LTE    string = "_lte"
	ARG_NOTEQ  string = "_notEq"
	ARG_NOTIN  string = "_notIn"

	ARG_ILIKE      string = "_iLike"
	ARG_IREGEX     string = "_iregex"
	ARG_LIKE       string = "_like"
	ARG_NOTILIKE   string = "_notILike"
	ARG_NOTIREGEX  string = "_notIRegex"
	ARG_NOTLIKE    string = "_notLike"
	ARG_NOTREGEX   string = "_notRegex"
	ARG_NOTSIMILAR string = "_notSimilar"
	ARG_REGEX      string = "_regex"
	ARG_SIMILAR    string = "_similar"
)

const (
	ARG_OBJECT            string = "object"
	ARG_OBJECTS           string = "objects"
	RESPONSE_RETURNING    string = "returning"
	RESPONSE_AFFECTEDROWS string = "affectedRows"
)

const (
	UUID    string = "uuid"
	INNERID string = "innerId"
	TYPE    string = "type"
)

/**
* Meta实体用到的常量
**/
const (
	META_ENTITY_NAME string = "Meta"
	META_ID          string = "id"
	META_STATUS      string = "status"
	META_CONTENT     string = "content"
	META_PUBLISHEDAT string = "publishedAt"
	META_CREATEDAT   string = "createdAt"
	META_UPDATEDAT   string = "updatedAt"

	META_ENTITIES  string = "entities"
	META_RELATIONS string = "relations"
)

const (
	ID_SUFFIX     string = "_id"
	PIVOT         string = "pivot"
	INDEX_SUFFIX  string = "_idx"
	SUFFIX_SOURCE string = "_source"
	SUFFIX_TARGET string = "_target"
)

const (
	CREATED_AT string = "createdAt"
	ID         string = "id"
	OF         string = "Of"
)

const (
	DELETED_AT string = "deletedAt"
)

const (
	BOOLEXP           string = "BoolExp"
	ORDERBY           string = "OrderBy"
	DISTINCTEXP       string = "DistinctExp"
	MUTATION_RESPONSE string = "MutationResponse"
)

const ASSOCIATION_OWNER_ID = "owner__rx__id"

const META_USER = "User"
const META_ROLE = "Role"

const SYSTEM = "System"
const CREATEDATE = "createDate"
const UPDATEDATE = "updateDate"

const (
	TOKEN          = "token"
	AUTHORIZATION  = "Authorization"
	BEARER         = "Bearer "
	CONTEXT_VALUES = "values"
)

const ABILITY_UUID = "META_ABILITY_UUID"

const (
	META_INNER_ID                 = 1
	ENTITY_AUTH_SETTINGS_INNER_ID = 2
	Ability_INNER_ID              = 3
	USER_INNER_ID                 = 4
	ROLE_INNER_ID                 = 5
	ROLE_USER_RELATION_INNER_ID   = 101
)

//普通角色的ID永远不会是1
const GUEST_ROLE_ID = "1"
const PREDEFINED_QUERYUSER = "$queryUser"
const PREDEFINED_ME = "$me"
