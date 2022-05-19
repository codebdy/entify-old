package resolve

import (
	"time"

	"github.com/graphql-go/graphql"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/crypto/bcrypt"
	"rxdrag.com/entify/config"
	"rxdrag.com/entify/consts"
	"rxdrag.com/entify/model"
	"rxdrag.com/entify/model/data"
	"rxdrag.com/entify/repository"
	"rxdrag.com/entify/utils"
)

type InstallArg struct {
	ID            int    `json:"id"`
	Admin         string `json:"admin"`
	AdminPassword string `json:"adminPassword"`
	WithDemo      bool   `json:"withDemo"`
}

const INPUT = "input"

func bcryptEncode(value string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	if err != nil {
		panic(err.Error())
	}
	encodeValue := string(hash)
	return encodeValue
}

// // 正确密码验证
// err = bcrypt.CompareHashAndPassword([]byte(encodePW), []byte(passwordOK))
// if err != nil {
//     fmt.Println("pw wrong")
// } else {
//     fmt.Println("pw ok")
// }

// // 错误密码验证
// err = bcrypt.CompareHashAndPassword([]byte(encodePW), []byte(passwordERR))
// if err != nil {
//     fmt.Println("pw wrong")
// } else {
//     fmt.Println("pw ok")
// }

func userEntity() map[string]interface{} {
	return map[string]interface{}{
		"content": map[string]interface{}{
			"classes": []map[string]interface{}{
				{
					consts.NAME: "User",
					consts.UUID: "META_USER_UUID",
					"innerId":   2,
					"attributes": []map[string]interface{}{
						{
							consts.NAME: "id",
							consts.TYPE: "ID",
							consts.UUID: "RX_USER_ID_UUID",
							"primary":   true,
							"typeLabel": "ID",
						},
						{
							consts.NAME: "name",
							consts.TYPE: "String",
							consts.UUID: "RX_USER_NAME_UUID",
							"typeLabel": "String",
							"nullable":  true,
						},
						{
							consts.NAME: "loginName",
							consts.TYPE: "String",
							consts.UUID: "RX_USER_LOGINNAME_UUID",
							"typeLabel": "String",
						},
						{
							consts.NAME: "password",
							consts.TYPE: "String",
							consts.UUID: "RX_USER_PASSWORD_UUID",
							"typeLabel": "String",
						},
						{
							consts.NAME: "isSupper",
							consts.TYPE: "Boolean",
							consts.UUID: "RX_USER_ISSUPPER_UUID",
							"typeLabel": "Boolean",
							"nullable":  true,
						},
						{
							consts.NAME: "isDemo",
							consts.TYPE: "Boolean",
							consts.UUID: "RX_USER_ISDEMO_UUID",
							"typeLabel": "Boolean",
							"nullable":  true,
						},
						{
							consts.NAME: consts.META_CREATEDAT,
							consts.TYPE: "Date",
							consts.UUID: "RX_USER_CREATEDAT_UUID",
							"typeLabel": "Date",
						},
						{
							consts.NAME: consts.META_UPDATEDAT,
							consts.TYPE: "Date",
							consts.UUID: "RX_USER_META_UPDATEDAT_UUID",
							"typeLabel": "Date",
						},
					},
					"stereoType": "Entity",
				},
			},
		},
		consts.META_CREATEDAT: time.Now(),
		consts.META_UPDATEDAT: time.Now(),
	}
}

func adminInstance(name string, password string) map[string]interface{} {
	return map[string]interface{}{
		consts.NAME:           "Admin",
		consts.LOGIN_NAME:     name,
		consts.PASSWORD:       bcryptEncode(password),
		consts.META_CREATEDAT: time.Now(),
		consts.META_UPDATEDAT: time.Now(),
	}
}

func demoInstance() map[string]interface{} {
	return map[string]interface{}{
		consts.NAME:           "Demo",
		consts.LOGIN_NAME:     "demo",
		consts.PASSWORD:       bcryptEncode("demo"),
		consts.META_CREATEDAT: time.Now(),
		consts.META_UPDATEDAT: time.Now(),
	}
}

func InstallResolve(p graphql.ResolveParams) (interface{}, error) {
	defer utils.PrintErrorStack()
	dbConfig := config.DbConfig{}
	input := InstallArg{}
	mapstructure.Decode(p.Args[INPUT], &dbConfig)
	mapstructure.Decode(p.Args[INPUT], &input)
	//config.SetDbConfig(dbConfig)

	//创建通过 Install 创建Meta表
	repository.Install(dbConfig)

	//创建User实体
	instance := data.NewInstance(userEntity(), model.GlobalModel.Graph.GetMetaEntity())
	_, err := repository.SaveOne(instance)

	if err != nil {
		return nil, err
	}
	err = doPublish()
	if err != nil {
		return nil, err
	}
	if input.Admin != "" {
		instance = data.NewInstance(
			adminInstance(input.Admin, input.AdminPassword),
			model.GlobalModel.Graph.GetEntityByName("User"),
		)
		_, err = repository.SaveOne(instance)
		if err != nil {
			return nil, err
		}
		if input.WithDemo {
			instance = data.NewInstance(
				demoInstance(),
				model.GlobalModel.Graph.GetEntityByName("User"),
			)
			_, err = repository.SaveOne(instance)
			if err != nil {
				return nil, err
			}
		}
	}
	//config.SetBool(consts.INSTALLED, true)
	//config.SetInt(consts.SERVICE_ID, input.ID)
	//config.WriteConfig()
	return repository.Installed(), nil
}
