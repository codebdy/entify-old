package resolve

import (
	"time"

	"github.com/graphql-go/graphql"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/crypto/bcrypt"
	"rxdrag.com/entify/consts"
	"rxdrag.com/entify/model"
	"rxdrag.com/entify/model/data"
	"rxdrag.com/entify/repository"
	"rxdrag.com/entify/utils"
)

type InstallArg struct {
	Admin    string `json:"admin"`
	Password string `json:"password"`
	WithDemo bool   `json:"withDemo"`
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

func predefinedEntities() map[string]interface{} {
	return map[string]interface{}{
		"content": map[string]interface{}{
			"classes": []map[string]interface{}{
				{
					consts.NAME:   consts.META_USER,
					consts.UUID:   "META_USER_UUID",
					"innerId":     consts.USER_INNER_ID,
					"root":        true,
					consts.SYSTEM: true,
					"attributes": []map[string]interface{}{
						{
							consts.NAME:   "id",
							consts.TYPE:   "ID",
							consts.UUID:   "RX_USER_ID_UUID",
							"primary":     true,
							"typeLabel":   "ID",
							consts.SYSTEM: true,
						},
						{
							consts.NAME:   "name",
							consts.TYPE:   "String",
							consts.UUID:   "RX_USER_NAME_UUID",
							"typeLabel":   "String",
							"length":      128,
							"nullable":    true,
							consts.SYSTEM: true,
						},
						{
							consts.NAME:   "loginName",
							consts.TYPE:   "String",
							consts.UUID:   "RX_USER_LOGINNAME_UUID",
							"typeLabel":   "String",
							"length":      128,
							consts.SYSTEM: true,
						},
						{
							consts.NAME:   "password",
							consts.TYPE:   "String",
							consts.UUID:   "RX_USER_PASSWORD_UUID",
							"typeLabel":   "String",
							"length":      256,
							consts.SYSTEM: true,
						},
						{
							consts.NAME:   "isSupper",
							consts.TYPE:   "Boolean",
							consts.UUID:   "RX_USER_ISSUPPER_UUID",
							"typeLabel":   "Boolean",
							"nullable":    true,
							consts.SYSTEM: true,
						},
						{
							consts.NAME:   "isDemo",
							consts.TYPE:   "Boolean",
							consts.UUID:   "RX_USER_ISDEMO_UUID",
							"typeLabel":   "Boolean",
							"nullable":    true,
							consts.SYSTEM: true,
						},
						{
							consts.NAME:       consts.META_CREATEDAT,
							consts.TYPE:       "Date",
							consts.UUID:       "RX_USER_CREATEDAT_UUID",
							"typeLabel":       "Date",
							consts.CREATEDATE: true,
							consts.SYSTEM:     true,
						},
						{
							consts.NAME:       consts.META_UPDATEDAT,
							consts.TYPE:       "Date",
							consts.UUID:       "RX_USER_META_UPDATEDAT_UUID",
							"typeLabel":       "Date",
							consts.UPDATEDATE: true,
							consts.SYSTEM:     true,
						},
					},
					"stereoType": "Entity",
				},
				{
					consts.NAME:   consts.META_ROLE,
					consts.UUID:   "META_ROLE_UUID",
					"innerId":     consts.ROLE_INNER_ID,
					"root":        true,
					consts.SYSTEM: true,
					"attributes": []map[string]interface{}{
						{
							consts.NAME:   "id",
							consts.TYPE:   "ID",
							consts.UUID:   "RX_ROLE_ID_UUID",
							"primary":     true,
							"typeLabel":   "ID",
							consts.SYSTEM: true,
						},
						{
							consts.NAME:   "name",
							consts.TYPE:   "String",
							consts.UUID:   "RX_ROLE_NAME_UUID",
							"typeLabel":   "String",
							consts.SYSTEM: true,
						},
						{
							consts.NAME:   "description",
							consts.TYPE:   "String",
							consts.UUID:   "RX_ROLE_DESCRIPTION_UUID",
							"typeLabel":   "String",
							"nullable":    true,
							consts.SYSTEM: true,
						},
						{
							consts.NAME:       consts.META_CREATEDAT,
							consts.TYPE:       "Date",
							consts.UUID:       "RX_ROLE_CREATEDAT_UUID",
							"typeLabel":       "Date",
							consts.CREATEDATE: true,
							consts.SYSTEM:     true,
						},
						{
							consts.NAME:       consts.META_UPDATEDAT,
							consts.TYPE:       "Date",
							consts.UUID:       "RX_ROLE_META_UPDATEDAT_UUID",
							"typeLabel":       "Date",
							consts.UPDATEDATE: true,
							consts.SYSTEM:     true,
						},
					},
					"stereoType": "Entity",
				},
			},
			"relations": []map[string]interface{}{
				{
					"uuid":               "META_RELATION_USER_ROLE_UUID",
					"innerId":            consts.ROLE_USER_RELATION_INNER_ID,
					"sourceId":           "META_ROLE_UUID",
					"targetId":           "META_USER_UUID",
					"relationType":       "twoWayAssociation",
					"roleOfSource":       "roles",
					"roleOfTarget":       "users",
					"sourceMutiplicity":  "0..*",
					"targetMultiplicity": "0..*",
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
	input := InstallArg{}
	mapstructure.Decode(p.Args[INPUT], &input)

	//创建实体
	instance := data.NewInstance(predefinedEntities(), model.GlobalModel.Graph.GetMetaEntity())
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
			adminInstance(input.Admin, input.Password),
			model.GlobalModel.Graph.GetEntityByName(consts.META_USER),
		)
		_, err = repository.SaveOne(instance)
		if err != nil {
			return nil, err
		}
		if input.WithDemo {
			instance = data.NewInstance(
				demoInstance(),
				model.GlobalModel.Graph.GetEntityByName(consts.META_USER),
			)
			_, err = repository.SaveOne(instance)
			if err != nil {
				return nil, err
			}
		}
	}
	return repository.IsEntityExists(consts.META_USER), nil
}
