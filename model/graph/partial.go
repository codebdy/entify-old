package graph

import (
	"rxdrag.com/entify/consts"
	"rxdrag.com/entify/model/domain"
	"rxdrag.com/entify/utils"
)

type Partial struct {
	Entity
}

func NewPartial(c *domain.Class) *Partial {
	return &Partial{
		Entity: Entity{
			Class: *NewClass(c),
		},
	}
}

func (p *Partial) NameWithPartial() string {
	return p.Domain.Name + utils.FirstUpper(p.Domain.PartialName)
}

func (p *Partial) QueryName() string {
	return utils.FirstLower(p.NameWithPartial())
}

func (p *Partial) QueryOneName() string {
	return consts.ONE + p.NameWithPartial()
}

func (p *Partial) QueryAggregateName() string {
	return utils.FirstLower(p.NameWithPartial()) + utils.FirstUpper(consts.AGGREGATE)
}

func (p *Partial) DeleteName() string {
	return consts.DELETE + p.NameWithPartial()
}

func (p *Partial) DeleteByIdName() string {
	return consts.DELETE + p.NameWithPartial() + consts.BY_ID
}

func (p *Partial) UpdateName() string {
	return utils.FirstLower(p.NameWithPartial())
}

func (p *Partial) UpsertName() string {
	return consts.UPSERT + p.NameWithPartial()
}

func (p *Partial) UpsertOneName() string {
	return consts.UPSERT_ONE + p.NameWithPartial()
}

func (p *Partial) AggregateName() string {
	return p.NameWithPartial() + utils.FirstUpper(consts.AGGREGATE)
}
