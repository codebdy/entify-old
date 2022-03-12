package meta

const (
	INHERIT      string = "inherit"
	ONE_TO_ONE   string = "oneToOne"
	ONE_TO_MANY  string = "oneToMany"
	MANY_TO_ONE  string = "manyToOne"
	MANY_TO_MANY string = "manyToMany"

	ON_SOURCE string = "onSource"
	ON_TARGET string = "onTarget"
)

type Relation struct {
	Uuid string `json:"uuid"`

	RelationType string `json:"relationType"`

	SourceId string `json:"sourceId"`

	TargetId string `json:"targetId"`

	RoleOnSource string `json:"roleOnSource"`

	RoleOnTarget string `json:"roleOnTarget"`

	CascadeOn string `json:"cascadeOn"`
}
