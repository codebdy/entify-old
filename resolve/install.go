package resolve

import (
	"time"

	"github.com/graphql-go/graphql"
	"github.com/mitchellh/mapstructure"
	"rxdrag.com/entify/config"
	"rxdrag.com/entify/consts"
	"rxdrag.com/entify/model"
	"rxdrag.com/entify/model/data"
	"rxdrag.com/entify/model/meta"
	"rxdrag.com/entify/repository"
)

type InstallArg struct {
	DbConfig      config.DbConfig
	Admin         string `json:"admin"`
	AdminPassword string `json:"adminPassword"`
	WithDemo      string `json:"withDemo"`
}

const INPUT = "input"

func InstallResolve(p graphql.ResolveParams) (interface{}, error) {
	input := InstallArg{}
	mapstructure.Decode(p.Args[INPUT], &input)

	//创建通过 Install 创建Meta表
	repository.Install(input.DbConfig)

	//创建User实体
	object := map[string]interface{}{
		"content": map[string]interface{}{
			"classes": []map[string]interface{}{
				{
					"name":    "User",
					"uuid":    "META_USER_UUID",
					"innerId": 2,
					"attributes": []map[string]interface{}{
						{
							"name":      "id",
							"type":      "ID",
							"uuid":      "RX_USER_ID_UUID",
							"primary":   true,
							"typeLabel": "ID",
						},
						{
							"name":      "name",
							"type":      "String",
							"uuid":      "RX_USER_NAME_UUID",
							"typeLabel": "String",
						},
						{
							"name":      "loginName",
							"type":      "String",
							"uuid":      "RX_USER_LOGINNAME_UUID",
							"typeLabel": "String",
						},
						{
							"name":      "password",
							"type":      "String",
							"uuid":      "RX_USER_PASSWORD_UUID",
							"typeLabel": "String",
						},
						{
							"name":      "isSupper",
							"type":      "Boolean",
							"uuid":      "RX_USER_ISSUPPER_UUID",
							"typeLabel": "Boolean",
						},
						{
							"name":      "isDemo",
							"type":      "Boolean",
							"uuid":      "RX_USER_ISDEMO_UUID",
							"typeLabel": "Boolean",
						},
					},
					"stereoType": "Entity",
				},
			},
			consts.META_STATUS:      meta.META_STATUS_PUBLISHED,
			consts.META_PUBLISHEDAT: time.Now(),
			consts.META_CREATEDAT:   time.Now(),
			consts.META_UPDATEDAT:   time.Now(),
		},
	}
	instance := data.NewInstance(object, model.GlobalModel.Graph.GetMetaEntity())
	repository.SaveOne(instance)

	config.SetDbConfig(input.DbConfig)
	config.SetBool(consts.INSTALLED, true)
	config.WriteConfig()
	return config.GetBool(consts.INSTALLED), nil
}
