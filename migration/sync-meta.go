package migration

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/repository"
	"rxdrag.com/entity-engine/schema"
)

const (
	MEA_STATUS_PUBLISHED        string = "published"
	META_STATUS_MIGRATION_ERROR string = "migration-error"
	META_STATUS_ROLLBACK_ERROR  string = "rollback-error"
)

func SyncMetaResolve(p graphql.ResolveParams) (interface{}, error) {
	object := p.Args[consts.ARG_OBJECT].(map[string]interface{})
	return repository.InsertOne(object, &schema.MetaEntity)
}
