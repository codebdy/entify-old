package entity

import "strconv"

type Role struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (r Role) Uint64Id() uint64 {
	number, err := strconv.ParseUint(r.Id, 10, 64)
	if err != nil {
		panic(err.Error())
	}
	return number
}
