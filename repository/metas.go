package repository

import "rxdrag.com/entity-engine/meta"

var Entities = &[]*meta.Entity{
	&meta.MetaStatusEnum,
	&meta.MetaEntity,
}

func GetEntityByUuid(uuid string) *meta.Entity {
	for _, entity := range *Entities {
		if entity.Uuid == uuid {
			return entity
		}
	}

	return nil
}
