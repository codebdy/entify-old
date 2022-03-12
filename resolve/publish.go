package resolve

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/meta"
	"rxdrag.com/entity-engine/migration"
	"rxdrag.com/entity-engine/repository"
	"rxdrag.com/entity-engine/utils"
)

func PublishMetaResolve(p graphql.ResolveParams) (interface{}, error) {
	publishedMeta, err := repository.QueryOne(&meta.MetaEntity, repository.QueryArg{
		consts.ARG_WHERE: repository.QueryArg{
			consts.META_STATUS: repository.QueryArg{
				consts.AEG_EQ: meta.META_STATUS_PUBLISHED,
			},
		},
	})
	if err != nil {
		panic("Read published meta error" + err.Error())
	}
	nextMeta, err := repository.QueryOne(&meta.MetaEntity, repository.QueryArg{
		consts.ARG_WHERE: repository.QueryArg{
			consts.META_STATUS: repository.QueryArg{
				consts.ARG_ISNULL: true,
			},
		},
	})
	if err != nil {
		panic("Read next meta error" + err.Error())
	}

	if nextMeta == nil {
		panic("Can not find unpublished meta")
	}
	var publishedContent map[string]interface{}
	var nextContent map[string]interface{}
	if publishedMeta != nil {
		publishedContent = publishedMeta.(utils.Object)[consts.META_CONTENT].(utils.Object)
	}
	diff := migration.CreateDiff(publishedContent, nextContent)
	migration.ExcuteDiff(diff)
	return nil, nil
}

func SyncMetaResolve(p graphql.ResolveParams) (interface{}, error) {
	object := p.Args[consts.ARG_OBJECT].(map[string]interface{})
	return repository.InsertOne(object, &meta.MetaEntity)
}
