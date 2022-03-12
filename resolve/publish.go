package resolve

import (
	"github.com/graphql-go/graphql"
	"github.com/mitchellh/mapstructure"
	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/meta"
	"rxdrag.com/entity-engine/migration"
	"rxdrag.com/entity-engine/repository"
	"rxdrag.com/entity-engine/utils"
)

func decodeContent(obj interface{}) *meta.MetaContent {
	content := meta.MetaContent{}
	if obj != nil {
		err := mapstructure.Decode(obj.(utils.Object)[consts.META_CONTENT], &content)
		if err != nil {
			panic("Decode content failure:" + err.Error())
		}
	}
	return &content
}

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
	publishedContent := decodeContent(publishedMeta)
	nextContent := decodeContent(nextMeta)
	diff := migration.CreateDiff(publishedContent, nextContent)
	migration.ExcuteDiff(diff)
	return nil, nil
}

func SyncMetaResolve(p graphql.ResolveParams) (interface{}, error) {
	object := p.Args[consts.ARG_OBJECT].(map[string]interface{})
	return repository.InsertOne(object, &meta.MetaEntity)
}
