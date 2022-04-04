package repositoryold

import (
	"github.com/mitchellh/mapstructure"
	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/oldmeta"
	"rxdrag.com/entity-engine/utils"
)

func QueryPublishedMeta() interface{} {
	publishedMeta, err := QueryOne(modelold.TheModel.GetMetaEntity(), QueryArg{
		consts.ARG_WHERE: QueryArg{
			consts.META_STATUS: QueryArg{
				consts.AEG_EQ: modelold.META_STATUS_PUBLISHED,
			},
		},
	})
	if err != nil {
		panic("Read published meta error" + err.Error())
	}

	return publishedMeta
}

func QueryNextMeta() interface{} {
	nextMeta, err := QueryOne(modelold.TheModel.GetMetaEntity(), QueryArg{
		consts.ARG_WHERE: QueryArg{
			consts.META_STATUS: QueryArg{
				consts.ARG_ISNULL: true,
			},
		},
	})
	if err != nil {
		panic("Read next meta error" + err.Error())
	}

	return nextMeta
}

func DecodeContent(obj interface{}) *oldmeta.MetaContent {
	content := oldmeta.MetaContent{}
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
	initMeta := oldmeta.MetaContent{
		Entities: []oldmeta.EntityMeta{
			modelold.MetaStatusEnum,
			modelold.MetaEntity,
		},
	}
	modelold.TheModel = modelold.NewModel(&initMeta)
	publishedMeta := QueryPublishedMeta()
	publishedContent := DecodeContent(publishedMeta)
	publishedContent.Entities = append(publishedContent.Entities, modelold.MetaStatusEnum)
	publishedContent.Entities = append(publishedContent.Entities, modelold.MetaEntity)

	modelold.TheModel = modelold.NewModel(publishedContent)
}

func init() {
	LoadModel()
}
