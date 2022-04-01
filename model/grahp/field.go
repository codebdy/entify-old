package grahp

import (
	"rxdrag.com/entity-engine/model/domain"
	"rxdrag.com/entity-engine/model/table"
)

type AttributeField struct {
}

type AssociationField struct {
	Table *table.Table
}

type MethodField struct {
	Method *domain.Method
}
