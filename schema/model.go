package schema

import (
	"github.com/mitchellh/mapstructure"
	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/model"
	"rxdrag.com/entity-engine/model/meta"

	"rxdrag.com/entity-engine/repository"
	"rxdrag.com/entity-engine/utils"
)

func QueryPublishedMeta() interface{} {
	publishedMeta, err := repository.QueryOne(model.GlobalModel.Graph.GetMetaEntity(), repository.QueryArg{
		consts.ARG_WHERE: repository.QueryArg{
			consts.META_STATUS: repository.QueryArg{
				consts.AEG_EQ: meta.META_STATUS_PUBLISHED,
			},
		},
	})
	if err != nil {
		panic("Read published meta error" + err.Error())
	}

	return publishedMeta
}

func QueryNextMeta() interface{} {
	nextMeta, err := repository.QueryOne(model.GlobalModel.Graph.GetMetaEntity(), repository.QueryArg{
		consts.ARG_WHERE: repository.QueryArg{
			consts.META_STATUS: repository.QueryArg{
				consts.ARG_ISNULL: true,
			},
		},
	})
	if err != nil {
		panic("Read next meta error" + err.Error())
	}

	return nextMeta
}

func DecodeContent(obj interface{}) *meta.MetaContent {
	content := meta.MetaContent{}
	if obj != nil {
		err := mapstructure.Decode(obj.(utils.Object)[consts.META_CONTENT], &content)
		if err != nil {
			panic("Decode content failure:" + err.Error())
		}
	}
	return &content
}

func LoadModel() {
	//初始值，用户取meta信息，取完后，换掉该部分内容
	initMeta := meta.MetaContent{
		Classes: []meta.ClassMeta{
			meta.MetaStatusEnum,
			meta.MetaClass,
		},
	}
	model.GlobalModel = model.New(&initMeta)
	publishedMeta := QueryPublishedMeta()
	publishedContent := DecodeContent(publishedMeta)
	publishedContent.Classes = append(publishedContent.Classes, meta.MetaStatusEnum)
	publishedContent.Classes = append(publishedContent.Classes, meta.MetaClass)

	model.GlobalModel = model.New(publishedContent)
}

func init() {
	LoadModel()
}
