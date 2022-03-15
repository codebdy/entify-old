package meta

type MetaContent struct {
	Entities  []Entity      `json:"entities"`
	Relations []Relation    `json:"relations"`
	Diagrams  []interface{} `json:"diagrams"`
	X6Nodes   []interface{} `json:"x6Nodes"`
	X6Edges   []interface{} `json:"x6Edges"`
}

func (entity *MetaContent) EntityRelations(entityUuid string) []EntityRelation {
	return []EntityRelation{}
}
