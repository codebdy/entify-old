package meta

import "rxdrag.com/entity-engine/consts"

const (
	IMPLEMENTS   string = "implements"
	ONE_TO_ONE   string = "oneToOne"
	ONE_TO_MANY  string = "oneToMany"
	MANY_TO_ONE  string = "manyToOne"
	MANY_TO_MANY string = "manyToMany"

	ON_SOURCE string = "onSource"
	ON_TARGET string = "onTarget"
)

type RelationMeta struct {
	Uuid                string `json:"uuid"`
	RelationType        string `json:"relationType"`
	SourceId            string `json:"sourceId"`
	TargetId            string `json:"targetId"`
	RoleOnSource        string `json:"roleOnSource"`
	RoleOnTarget        string `json:"roleOnTarget"`
	DescriptionOnSource string `json:"descriptionOnSource"`
	DescriptionOnTarget string `json:"descriptionOnTarget"`
	CascadeType         string `json:"cascadeType"`
	OwnerId             string `json:"ownerId"`
	//多对多关联自定义列
	Columns []ColumnMeta `json:"columns"`
}

func (r *RelationMeta) RelationSourceColumnName() string {
	return r.RoleOnSource + consts.ID_SUFFIX
}

func (r *RelationMeta) RelationTargetColumnName() string {
	return r.RoleOnTarget + consts.ID_SUFFIX
}
