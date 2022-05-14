package resolve

import (
	"fmt"
	"time"

	"github.com/graphql-go/graphql"
	"rxdrag.com/entify/consts"
	"rxdrag.com/entify/model"
	"rxdrag.com/entify/model/data"
	"rxdrag.com/entify/model/meta"
	"rxdrag.com/entify/repository"
	"rxdrag.com/entify/utils"
)

func doPublish() error {
	publishedMeta := repository.QueryPublishedMeta()
	nextMeta := repository.QueryNextMeta()
	fmt.Println("Start to publish")
	// fmt.Println("Published Meta ID:", publishedMeta.(utils.Object)["id"])
	// fmt.Println("Next Meta ID:", nextMeta.(utils.Object)["id"])

	if nextMeta == nil {
		panic("Can not find unpublished meta")
	}
	publishedModel := model.New(repository.DecodeContent(publishedMeta))
	nextModel := model.New(repository.DecodeContent(nextMeta))
	nextModel.Graph.Validate()
	diff := model.CreateDiff(publishedModel, nextModel)
	repository.ExcuteDiff(diff)
	fmt.Println("ExcuteDiff success")
	metaObj := nextMeta.(utils.Object)
	metaObj[consts.META_STATUS] = meta.META_STATUS_PUBLISHED
	metaObj[consts.META_PUBLISHEDAT] = time.Now()
	_, err := repository.SaveOne(data.NewInstance(metaObj, model.GlobalModel.Graph.GetMetaEntity()))
	if err != nil {
		return err
	}
	repository.LoadModel()

	return nil
}

func PublishMetaResolve(p graphql.ResolveParams) (interface{}, error) {
	defer utils.PrintErrorStack()
	doPublish()
	return "success", nil
}

func SyncMetaResolve(p graphql.ResolveParams) (interface{}, error) {
	object := p.Args[consts.ARG_OBJECT].(map[string]interface{})
	return repository.InsertOne(data.NewInstance(object, model.GlobalModel.Graph.GetMetaEntity()))
}
