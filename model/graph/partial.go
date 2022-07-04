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
	return consts.ONE + utils.FirstUpper(p.NameWithPartial())
}

func (p *Partial) QueryAggregateName() string {
	return utils.FirstLower(p.NameWithPartial()) + utils.FirstUpper(consts.AGGREGATE)
}

func (p *Partial) DeleteName() string {
	return consts.DELETE + utils.FirstUpper(p.NameWithPartial())
}

func (p *Partial) DeleteByIdName() string {
	return consts.DELETE + utils.FirstUpper(p.NameWithPartial()) + consts.BY_ID
}

func (p *Partial) SetName() string {
	return consts.SET + utils.FirstUpper(p.NameWithPartial())
}

func (p *Partial) InsertName() string {
	return consts.INSERT + utils.FirstUpper(p.NameWithPartial())
}

func (p *Partial) InsertOneName() string {
	return consts.INSERT_ONE + utils.FirstUpper(p.NameWithPartial())
}

func (p *Partial) UpdateName() string {
	return consts.UPDATE + utils.FirstUpper(p.NameWithPartial())
}

func (p *Partial) UpdateOneName() string {
	return consts.UPDATE_ONE + utils.FirstUpper(p.NameWithPartial())
}

func (p *Partial) AggregateName() string {
	return p.NameWithPartial() + utils.FirstUpper(consts.AGGREGATE)
}

func (p *Partial) GetHasManyName() string {
	return utils.FirstUpper(consts.SET) + p.NameWithPartial() + consts.HAS_MANY
}

func (p *Partial) GetHasOneName() string {
	return utils.FirstUpper(consts.SET) + p.NameWithPartial() + consts.HAS_ONE
}
