package migration

import (
	"fmt"

	"rxdrag.com/entity-engine/meta"
)

func ValidateNextMeta(metaConent *meta.MetaContent) {
	for _, entity := range metaConent.Entities {
		if len(entity.Columns) <= 1 && entity.EntityType == meta.Entity_NORMAL {
			panic(fmt.Sprintf("Entity %s should have one normal field at least", entity.Name))
		}
	}
}
