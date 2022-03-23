package repository

import (
	"github.com/mitchellh/mapstructure"
	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/meta"
	"rxdrag.com/entity-engine/model"
	"rxdrag.com/entity-engine/utils"
)

func QueryPublishedMeta() interface{} {
	publishedMeta, err := QueryOne(&meta.MetaEntity, QueryArg{
		consts.ARG_WHERE: QueryArg{
			consts.META_STATUS: QueryArg{
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
	nextMeta, err := QueryOne(&meta.MetaEntity, QueryArg{
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

func GetEntityByUuid(uuid string) *model.Entity {
	for _, entity := range model.TheModel.Entities {
		if entity.Uuid == uuid {
			return entity
		}
	}

	return nil
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
	publishedMeta := QueryPublishedMeta()
	publishedContent := DecodeContent(publishedMeta)

	publishedContent.Entities = append(publishedContent.Entities, meta.MetaStatusEnum)
	publishedContent.Entities = append(publishedContent.Entities, meta.MetaEntity)

	model.TheModel = model.NewModel(publishedContent)
}

func init() {
	LoadModel()
}
