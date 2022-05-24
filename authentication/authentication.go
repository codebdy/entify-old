package authentication

import (
	"errors"
	"fmt"
	"strings"

	"github.com/mitchellh/mapstructure"
	"golang.org/x/crypto/bcrypt"
	"rxdrag.com/entify/authentication/jwt"
	"rxdrag.com/entify/config"
	"rxdrag.com/entify/consts"
	"rxdrag.com/entify/db/dialect"
	"rxdrag.com/entify/entity"
	"rxdrag.com/entify/model"
	"rxdrag.com/entify/repository"
)

var TokenCache = map[string]*entity.User{}

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

	token, err := jwt.GenerateToken(loginName)
	if err != nil {
		panic(err.Error())
	}

	userMap := repository.QueryOne(model.GlobalModel.Graph.GetEntityByName(consts.META_USER), repository.QueryArg{
		consts.ARG_WHERE: repository.QueryArg{
			consts.LOGIN_NAME: repository.QueryArg{
				consts.ARG_EQ: loginName,
			},
		},
	})

	var user entity.User

	err = mapstructure.Decode(userMap, user)

	if err != nil {
		panic(err.Error())
	}

	TokenCache[token] = &user

	return token, err
}

func GetUserByToken(token string) *entity.User {
	authUrl := config.AuthUrl()
	if authUrl == "" {
		return TokenCache[token]
	} else {
		return meFromRemote(token)
	}
}
