package repository

import (
	"encoding/json"

	"github.com/graphql-go/graphql"
	"github.com/mitchellh/mapstructure"
	"rxdrag.com/entify/authcontext"
	"rxdrag.com/entify/common"
	"rxdrag.com/entify/consts"
	"rxdrag.com/entify/model"
	"rxdrag.com/entify/model/meta"
)

type AbilityVerifier struct {
	Me        *common.User
	RoleIds   []string
	Abilities []*common.Ability
	// expression Key : 从Auth模块返回的结果
	QueryUserCache map[string][]common.User
	isSupper       bool
}

func NewVerifier() *AbilityVerifier {
	verifier := AbilityVerifier{}

	return &verifier
}

func NewSupperVerifier() *AbilityVerifier {
	verifier := AbilityVerifier{isSupper: true}

	return &verifier
}

func (v *AbilityVerifier) Init(p graphql.ResolveParams, entityUuids []string) {
	me := authcontext.ParseContextValues(p).Me
	v.Me = me
	if me != nil {
		for i := range me.Roles {
			v.RoleIds = append(v.RoleIds, me.Roles[i].Id)
		}
	} else {
		v.RoleIds = append(v.RoleIds, consts.GUEST_ROLE_ID)
	}

	v.queryRolesAbilities(entityUuids)
}

func (v *AbilityVerifier) IsSupper() bool {
	if v.isSupper {
		return true
	}

	if v.Me != nil {
		return v.Me.IsSupper
	}

	return false
}

func (v *AbilityVerifier) WeaveAuthInArgs(entityUuid string, args map[string]interface{}) map[string]interface{} {
	var rootAnd []map[string]interface{}
	if args[consts.ARG_AND] == nil {
		rootAnd = []map[string]interface{}{}
	} else {
		rootAnd = args[consts.ARG_AND].([]map[string]interface{})
	}

	expArg := v.queryEntityArgsMap(entityUuid)
	if len(expArg) > 0 {
		rootAnd = append(rootAnd, expArg)
	}

	args[consts.ARG_AND] = rootAnd
	return args
}

func (v *AbilityVerifier) CanReadEntity(entityUuid string) bool {
	if v.Me != nil && (v.Me.IsDemo || v.Me.IsSupper) {
		return true
	}
	for _, ability := range v.Abilities {
		if ability.EntityUuid == entityUuid &&
			ability.ColumnUuid == "" &&
			ability.Can &&
			ability.AbilityType == meta.META_ABILITY_TYPE_READ {
			return true
		}
	}
	return false
}

func (v *AbilityVerifier) EntityMutationCan(entityData map[string]interface{}) bool {
	return false
}

func (v *AbilityVerifier) FieldCan(entityData map[string]interface{}) bool {
	return false
}

func (v *AbilityVerifier) queryEntityArgsMap(entityUuid string) map[string]interface{} {
	expMap := map[string]interface{}{}
	queryEntityExpressions := []string{}
	for _, ability := range v.Abilities {
		if ability.EntityUuid == entityUuid &&
			ability.ColumnUuid == "" &&
			ability.Can &&
			ability.AbilityType == meta.META_ABILITY_TYPE_READ &&
			ability.Expression != "" {
			queryEntityExpressions = append(queryEntityExpressions, ability.Expression)
		}
	}
	if len(queryEntityExpressions) > 0 {
		expMap[consts.ARG_OR] = expressionArrayToArgs(queryEntityExpressions)
	}
	return expMap
}

func expressionToKey(expression string) string {
	return ""
}

func expressionArrayToArgs(expressionArray []string) []map[string]interface{} {
	var args []map[string]interface{}
	for _, expression := range expressionArray {
		args = append(args, expressionToArg(expression))
	}
	return args
}

func expressionToArg(expression string) map[string]interface{} {
	arg := map[string]interface{}{}
	err := json.Unmarshal([]byte(expression), &arg)
	if err != nil {
		panic("Parse authorization expression error:" + err.Error())
	}
	return arg
}

func (v *AbilityVerifier) queryRolesAbilities(entityUuids []string) {
	abilities := QueryEntity(model.GlobalModel.Graph.GetEntityByUuid(consts.ABILITY_UUID), QueryArg{
		consts.ARG_WHERE: QueryArg{
			"roleId": QueryArg{
				consts.ARG_IN: v.RoleIds,
			},
			// "abilityType": QueryArg{
			// 	consts.ARG_EQ: v.AbilityType,
			// },
			"entityUuid": QueryArg{
				consts.ARG_IN: entityUuids,
			},
		},
	}, NewSupperVerifier())

	for _, abilityMap := range abilities {
		var ability common.Ability
		err := mapstructure.Decode(abilityMap, &ability)
		if err != nil {
			panic(err.Error())
		}
		v.Abilities = append(v.Abilities, &ability)
	}
}
