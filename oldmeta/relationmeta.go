package oldmeta

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
	InnerId             uint64 `json:"innerId"`
	RelationType        string `json:"relationType"`
	SourceId            string `json:"sourceId"`
	TargetId            string `json:"targetId"`
	RoleOfTarget        string `json:"roleOfTarget"`
	RoleOfSource        string `json:"roleOfSource"`
	DescriptionOnSource string `json:"descriptionOnSource"`
	DescriptionOnTarget string `json:"descriptionOnTarget"`
	CascadeType         string `json:"cascadeType"`
	OwnerId             string `json:"ownerId"`
	//多对多关联自定义列
	Columns []ColumnMeta `json:"columns"`
}
