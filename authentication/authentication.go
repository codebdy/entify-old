package authentication

import (
	"errors"
	"fmt"

	"rxdrag.com/entify/config"
	"rxdrag.com/entify/repository"
)

func Login(loginName, pwd string) (string, error) {
	con, err := repository.Open(config.GetDbConfig())
	defer con.Close()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	var password string
	err = con.Dbx.QueryRow("select password from rx_user where loginName = ?", loginName).Scan(&password)
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
