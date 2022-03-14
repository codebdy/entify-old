package repository

import (
	"github.com/mitchellh/mapstructure"
	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/meta"
	"rxdrag.com/entity-engine/utils"
)

var Entities *[]*meta.Entity

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

func GetEntityByUuid(uuid string) *meta.Entity {
	for _, entity := range *Entities {
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

func LoadMetas() {
	publishedMeta := QueryPublishedMeta()
	publishedContent := DecodeContent(publishedMeta)

	theEntities := make([]*meta.Entity, len(publishedContent.Entities)+2)
	theEntities[0] = &meta.MetaStatusEnum
	theEntities[1] = &meta.MetaEntity
	for i := range publishedContent.Entities {
		theEntities[i+2] = &publishedContent.Entities[i]
	}

	Entities = &theEntities
}

func init() {
	LoadMetas()
}
