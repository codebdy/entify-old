package model

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entify/model/domain"
	"rxdrag.com/entify/model/graph"
	"rxdrag.com/entify/model/meta"
)

type Model struct {
	Meta   *meta.Model
	Domain *domain.Model
	Graph  *graph.Model
	Schema *graphql.Schema
}

func New(c *meta.MetaContent) *Model {
	metaModel := meta.New(c)
	domainModel := domain.New(metaModel)
	grahpModel := graph.New(domainModel)
	model := Model{
		Meta:   metaModel,
		Domain: domainModel,
		Graph:  grahpModel,
		Schema: nil,
	}
	return &model
}

var GlobalModel *Model
