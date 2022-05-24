package entity

import "strconv"

type User struct {
	Id        string `json:"id"` //ID类型，查询时需要用字符串接收
	Name      string `json:"name"`
	LoginName string `json:"loginName"`
	Roles     []Role `json:"roles"`
}

func (u User) Uint64Id() uint64 {
	number, err := strconv.ParseUint(u.Id, 10, 64)
	if err != nil {
		panic(err.Error())
	}
	return number
}
