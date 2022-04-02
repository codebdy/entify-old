package modleold

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/model/domain"
	"rxdrag.com/entity-engine/model/graph"
	"rxdrag.com/entity-engine/model/meta"
	"rxdrag.com/entity-engine/model/table"
)

type Model struct {
	Meta   *meta.Model
	Domain *domain.Model
	Grahp  *graph.Model
	Tables []*table.Table
	Schema *graphql.Schema
}

func NewModel(c *meta.MetaContent) *Model {
	metaModel := c.ToModel()
	model := Model{
		Meta:   metaModel,
		Domain: &domain.Model{},
		Grahp:  &graph.Model{},
		Tables: []*table.Table{},
		Schema: nil,
	}
	return &model
}
