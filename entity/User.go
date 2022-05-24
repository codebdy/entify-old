package entity

type User struct {
	Id        uint64 `json:"id"`
	LoginName string `json:"loginName"`
	Roles     []Role
}
