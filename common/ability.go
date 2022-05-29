package common

type Ability struct {
	Id          uint64 `json:"id"`
	EntityUuid  string `json:"entityUuid"`
	ColumnUuid  string `json:"columnUuid"`
	Can         bool   `json:"can"`
	Expression  string `json:"expression"`
	AbilityType string `json:"abilityType"`
	RoleId      uint64 `json:"roleId"`
}
