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

func InstallResolve(p graphql.ResolveParams) (interface{}, error) {
	defer utils.PrintErrorStack()
	input := InstallArg{}
	mapstructure.Decode(p.Args[INPUT], &input)
	verifier := repository.NewSupperVerifier()
	instance, err := addAndPublishMeta(authClasses, authRelations)
	if err != nil {
		return nil, err
	}

	repository.LoadModel()

	if input.Admin != "" {
		instance = data.NewInstance(
			adminInstance(input.Admin, input.Password),
			model.GlobalModel.Graph.GetEntityByName(consts.META_USER),
		)
		_, err = repository.SaveOne(instance, verifier)
		if err != nil {
			return nil, err
		}
		if input.WithDemo {
			instance = data.NewInstance(
				demoInstance(),
				model.GlobalModel.Graph.GetEntityByName(consts.META_USER),
			)
			_, err = repository.SaveOne(instance, verifier)
			if err != nil {
				return nil, err
			}
		}
	}
	return repository.IsEntityExists(consts.META_USER), nil
}

func InstallMedia() {
	_, err := addAndPublishMeta(mediaClasses, []map[string]interface{}{})

	if err != nil {
		panic(err.Error())
	}
}

func addAndPublishMeta(classes []map[string]interface{}, relations []map[string]interface{}) (*data.Instance, error) {
	verifier := repository.NewSupperVerifier()
	nextMeta := repository.QueryNextMeta()
	if nextMeta != nil {
		panic("Please pushish meta first then install new function ")
	}

	predefined := map[string]interface{}{
		"content": map[string]interface{}{
			"classes":   classes,
			"relations": relations,
		},
		consts.META_CREATEDAT: time.Now(),
		consts.META_UPDATEDAT: time.Now(),
	}
	publishedMeta := repository.QueryPublishedMeta()

	if publishedMeta != nil {
		metaContent := *(publishedMeta.(map[string]interface{})[consts.META_CONTENT].(*utils.JSON))
		predefined[consts.META_CREATEDAT] = publishedMeta.(map[string]interface{})[consts.META_CREATEDAT]
		clses := metaContent[consts.META_CLASSES].([]interface{})
		for i := range classes {
			metaContent[consts.META_CLASSES] = append(clses, classes[i])
		}

		relas := metaContent[consts.META_RELATIONS].([]interface{})
		for i := range relas {
			metaContent[consts.META_RELATIONS] = append(relas, relas[i])
		}

		predefined[consts.META_CONTENT] = metaContent
	}
	//创建实体
	instance := data.NewInstance(predefined, model.GlobalModel.Graph.GetMetaEntity())
	_, err := repository.SaveOne(instance, verifier)

	if err != nil {
		return nil, err
	}
	err = doPublish(verifier)
	if err != nil {
		return nil, err
	}

	return instance, nil
}

var mediaClasses = []map[string]interface{}{
	{
		consts.NAME:    consts.MEDIA_ENTITY_NAME,
		consts.UUID:    consts.MEDIA_UUID,
		consts.INNERID: consts.MEDIA_INNER_ID,
		consts.ROOT:    true,
		consts.SYSTEM:  true,
		"attributes": []map[string]interface{}{
			{
				consts.NAME:   "id",
				consts.TYPE:   "ID",
				consts.UUID:   "RX_MEDIA_ID_UUID",
				"primary":     true,
				"typeLabel":   "ID",
				consts.SYSTEM: true,
			},
			{
				consts.NAME:   "name",
				consts.TYPE:   "String",
				consts.UUID:   "RX_MEDIA_NAME_UUID",
				"typeLabel":   "String",
				"nullable":    true,
				consts.SYSTEM: true,
			},
			{
				consts.NAME:   "mimetype",
				consts.TYPE:   "String",
				consts.UUID:   "RX_MEDIA_MIMETYPE_UUID",
				"typeLabel":   "String",
				consts.SYSTEM: true,
			},
			{
				consts.NAME:   "fileName",
				consts.TYPE:   "String",
				consts.UUID:   "RX_MEDIA_FILENAME_UUID",
				"typeLabel":   "String",
				"length":      128,
				consts.SYSTEM: true,
			},
			{
				consts.NAME:   "path",
				consts.TYPE:   "String",
				consts.UUID:   "RX_MEDIA_PATH_UUID",
				"typeLabel":   "String",
				"length":      256,
				consts.SYSTEM: true,
			},
			{
				consts.NAME:   "size",
				consts.TYPE:   "Int",
				consts.UUID:   "RX_MEDIA_SIZE_UUID",
				"typeLabel":   "Int",
				consts.SYSTEM: true,
			},
			{
				consts.NAME:   "mediaType",
				consts.TYPE:   "String",
				consts.UUID:   "RX_MEDIA_MEDIATYPE_UUID",
				"typeLabel":   "String",
				consts.SYSTEM: true,
			},
			{
				consts.NAME:       consts.META_CREATEDAT,
				consts.TYPE:       "Date",
				consts.UUID:       "RX_MEDIA_CREATEDAT_UUID",
				"typeLabel":       "Date",
				consts.CREATEDATE: true,
				consts.SYSTEM:     true,
			},
			{
				consts.NAME:       consts.META_UPDATEDAT,
				consts.TYPE:       "Date",
				consts.UUID:       "RX_MEDIA_UPDATEDAT_UUID",
				"typeLabel":       "Date",
				consts.UPDATEDATE: true,
				consts.SYSTEM:     true,
			},
		},
		"stereoType": "Entity",
	},
}

var authClasses = []map[string]interface{}{
	{
		consts.NAME:    consts.META_USER,
		consts.UUID:    consts.USER_UUID,
		consts.INNERID: consts.USER_INNER_ID,
		consts.ROOT:    true,
		consts.SYSTEM:  true,
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
				consts.NAME:   consts.IS_SUPPER,
				consts.TYPE:   "Boolean",
				consts.UUID:   "RX_USER_ISSUPPER_UUID",
				"typeLabel":   "Boolean",
				"nullable":    true,
				consts.SYSTEM: true,
			},
			{
				consts.NAME:   consts.IS_DEMO,
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
				consts.UUID:       "RX_USER_UPDATEDAT_UUID",
				"typeLabel":       "Date",
				consts.UPDATEDATE: true,
				consts.SYSTEM:     true,
			},
		},
		"stereoType": "Entity",
	},
	{
		consts.NAME:    consts.META_ROLE,
		consts.UUID:    consts.ROLE_UUID,
		consts.INNERID: consts.ROLE_INNER_ID,
		consts.ROOT:    true,
		consts.SYSTEM:  true,
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
}

var authRelations = []map[string]interface{}{
	{
		consts.UUID:          "META_RELATION_USER_ROLE_UUID",
		consts.INNERID:       consts.ROLE_USER_RELATION_INNER_ID,
		"sourceId":           consts.ROLE_UUID,
		"targetId":           consts.USER_UUID,
		"relationType":       "twoWayAssociation",
		"roleOfSource":       "roles",
		"roleOfTarget":       "users",
		"sourceMutiplicity":  "0..*",
		"targetMultiplicity": "0..*",
	},
}

func bcryptEncode(value string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	if err != nil {
		panic(err.Error())
	}
	encodeValue := string(hash)
	return encodeValue
}

func adminInstance(name string, password string) map[string]interface{} {
	return map[string]interface{}{
		consts.NAME:           "Admin",
		consts.LOGIN_NAME:     name,
		consts.PASSWORD:       bcryptEncode(password),
		consts.IS_SUPPER:      true,
		consts.META_CREATEDAT: time.Now(),
		consts.META_UPDATEDAT: time.Now(),
	}
}

func demoInstance() map[string]interface{} {
	return map[string]interface{}{
		consts.NAME:           "Demo",
		consts.LOGIN_NAME:     "demo",
		consts.PASSWORD:       bcryptEncode("demo"),
		consts.IS_DEMO:        true,
		consts.META_CREATEDAT: time.Now(),
		consts.META_UPDATEDAT: time.Now(),
	}
}
