package authentication

import (
	"database/sql"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"rxdrag.com/entify/authentication/jwt"
	"rxdrag.com/entify/common"
	"rxdrag.com/entify/config"
	"rxdrag.com/entify/db/dialect"
	"rxdrag.com/entify/repository"
)

var TokenCache = map[string]*common.User{}

func loadUser(loginName string) *common.User {
	con, err := repository.Open(repository.NewSupperVerifier())
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	var user common.User

	sqlBuilder := dialect.GetSQLBuilder()
	err = con.Dbx.QueryRow(sqlBuilder.BuildMeSQL(), loginName).Scan(&user.Id, &user.Name, &user.LoginName)
	switch {
	case err == sql.ErrNoRows:
		return nil
	case err != nil:
		panic(err.Error())
	}

	rows, err := con.Dbx.Query(sqlBuilder.BuildRolesSQL(), user.Id)
	if err != nil {
		panic(err.Error())
	}
	for rows.Next() {
		var role common.Role
		err = rows.Scan(&role.Id, &role.Name)
		if err != nil {
			panic(err.Error())
		}
		user.Roles = append(user.Roles, role)
	}
	return &user
}

func Login(loginName, pwd string) (string, error) {
	con, err := repository.Open(repository.NewSupperVerifier())
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	sqlBuilder := dialect.GetSQLBuilder()
	var password string
	err = con.Dbx.QueryRow(sqlBuilder.BuildLoginSQL(), loginName).Scan(&password)
	if err != nil {
		fmt.Println(err)
		return "", errors.New("Login failed!")
	}

	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(pwd)) //验证（对比）
	if err != nil {
		fmt.Println(err, pwd, password)
		return "", errors.New("Password error!")
	}

	token, err := jwt.GenerateToken(loginName)
	if err != nil {
		panic(err.Error())
	}

	user := loadUser(loginName)
	TokenCache[token] = user
	return token, err
}

func Logout(token string) {
	TokenCache[token] = nil
}

func GetUserByToken(token string) (*common.User, error) {
	authUrl := config.AuthUrl()
	if authUrl == "" {
		return TokenCache[token], nil
	} else {
		return meFromRemote(token)
	}
}
