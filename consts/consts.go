package consts

const (
	ROOT_QUERY_NAME        = "Query"
	ROOT_MUTATION_NAME     = "Mutation"
	ROOT_SUBSCRIPTION_NAME = "Subscription"

	LOGIN      = "login"
	LOGIN_NAME = "loginName"
	PASSWORD   = "password"
	LOGOUT     = "logout"
	PUBLISH    = "_publish"
	ROLLBACK   = "_rollback"
	SYNC_META  = "_syncMeta"

	ONE       = "one"
	QUERY     = "query"
	AGGREGATE = "aggregate"

	ARG_DISTINCTON string = "_distinctOn"
	ARG_LIMIT      string = "_limit"
	ARG_OFFSET     string = "_offset"
	ARG_ORDERBY    string = "_orderBy"
	ARG_WHERE      string = "_where"

	ARG_AND string = "_and"
	ARG_NOT string = "_not"
	ARG_OR  string = "_or"

	AEG_EQ string = "_eq"
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
	SUFFIX_PIVOT  string = "__pivot"
	PARENT_ID     string = "parent_part_id"
	INDEX_SUFFIX  string = "_idx"
	SUFFIX_SOURCE string = "_source"
	SUFFIX_TARGET string = "_target"
)

const (
	CREATED_AT string = "createdAt"
	CONST_ID   string = "id"
)
