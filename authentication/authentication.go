package authentication

import (
	"errors"
	"fmt"

	"rxdrag.com/entity-engine/repository"
)

func Login(loginName, pwd string) (string, error) {
	con, err := repository.OpenConnection()
	defer con.Close()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	var password string
	err = con.QueryRow("select password from rx_user where loginName = ?", loginName).Scan(&password)
	if err != nil {
		fmt.Println(err)
		return "", errors.New("Login failed!")
	}

	// err = bcrypt.CompareHashAndPassword([]byte(pwd), []byte(password)) //验证（对比）
	// if err != nil {
	// 	fmt.Println(err)
	// 	return "", errors.New("Password error!")
	// }
	return loginName, err
}

func Logout() {

}
