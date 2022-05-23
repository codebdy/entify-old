package authentication

import (
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"rxdrag.com/entify/db/dialect"
	"rxdrag.com/entify/repository"
)

func Login(loginName, pwd string) (string, error) {
	con, err := repository.Open()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	sqlBuilder := dialect.GetSQLBuilder()
	var password string
	err = con.Dbx.QueryRow(sqlBuilder.BuildLoginSQL(), strings.ToUpper(loginName)).Scan(&password)
	if err != nil {
		fmt.Println(err)
		return "", errors.New("Login failed!")
	}

	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(pwd)) //验证（对比）
	if err != nil {
		fmt.Println(err, pwd, password)
		return "", errors.New("Password error!")
	}
	return loginName, err
}

func Logout() {

}
