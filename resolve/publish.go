package resolve

import (
	"time"

	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/meta"
	"rxdrag.com/entity-engine/model"
	"rxdrag.com/entity-engine/repository"
	"rxdrag.com/entity-engine/utils"
)

func PublishMetaResolve(p graphql.ResolveParams) (interface{}, error) {
	publishedMeta := repository.QueryPublishedMeta()
	nextMeta := repository.QueryNextMeta()

	// fmt.Println("Published Meta ID:", publishedMeta.(utils.Object)["id"])
	// fmt.Println("Next Meta ID:", nextMeta.(utils.Object)["id"])

	if nextMeta == nil {
		panic("Can not find unpublished meta")
	}
	publishedModel := model.NewModel(repository.DecodeContent(publishedMeta))
	nextModel := model.NewModel(repository.DecodeContent(nextMeta))
	nextModel.Validate()
	diff := model.CreateDiff(publishedModel, nextModel)
	repository.ExcuteDiff(diff)
	metaObj := nextMeta.(utils.Object)
	metaObj[consts.META_STATUS] = meta.META_STATUS_PUBLISHED
	metaObj[consts.META_PUBLISHEDAT] = time.Now()
	repository.SaveOne(metaObj, &meta.MetaEntity)
	repository.LoadModel()
	return nil, nil
}

func SyncMetaResolve(p graphql.ResolveParams) (interface{}, error) {
	object := p.Args[consts.ARG_OBJECT].(map[string]interface{})
	return repository.InsertOne(object, &meta.MetaEntity)
}
