package entity

type User struct {
	Id        uint64 `json:"id"`
	Name      string `json:"name"`
	LoginName string `json:"loginName"`
	Roles     []Role `json:"roles"`
}
