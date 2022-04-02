package grahp

import (
	"rxdrag.com/entity-engine/model/domain"
	"rxdrag.com/entity-engine/model/table"
)

type Entity struct {
	Class        *domain.Class
	Table        *table.Table
	Fields       []*AttributeField
	Associations []*AssociationField
	Methods      []*MethodField
	Interfaces   []*Interface
}
