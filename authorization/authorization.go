package authorization

import (
	"encoding/json"

	"github.com/graphql-go/graphql"
	"github.com/mitchellh/mapstructure"
	"rxdrag.com/entify/common"
	"rxdrag.com/entify/consts"
	"rxdrag.com/entify/model"
	"rxdrag.com/entify/model/meta"
	"rxdrag.com/entify/repository"
)

type AbilityVerifier struct {
	Me          *common.User
	roleIds     []string
	abilityType string
	abilities   []*common.Ability
	// expression Key : 从Auth模块返回的结果
	queryUserCache map[string][]common.User
}

func New() *AbilityVerifier {
	verifier := AbilityVerifier{}

	return &verifier
}

func (v *AbilityVerifier) Init(p graphql.ResolveParams, entityUuid string, abilityType string) {
	me := ParseContextValues(p).Me
	v.Me = me
	if me != nil {
		for i := range me.Roles {
			v.roleIds = append(v.roleIds, me.Roles[i].Id)
		}
	} else {
		v.roleIds = append(v.roleIds, consts.GUEST_ROLE_ID)
	}

	v.abilityType = abilityType

	v.queryRolesAbilities()
}

func (v *AbilityVerifier) WeaveAuthInArgs(args map[string]interface{}) map[string]interface{} {
	var rootAnd []map[string]interface{}
	if args[consts.ARG_ADD] == nil {
		rootAnd = []map[string]interface{}{}
	} else {
		rootAnd = args[consts.ARG_ADD].([]map[string]interface{})
	}

	expArg := v.queryEntityArgsMap()
	if len(expArg) > 0 {
		rootAnd = append(rootAnd, expArg)
	}

	args[consts.ARG_ADD] = rootAnd
	return args
}

func (v *AbilityVerifier) CanReadEntity() bool {
	if v.Me != nil && (v.Me.IsDemo || v.Me.IsSupper) {
		return true
	}
	for _, ability := range v.abilities {
		if ability.ColumnUuid == "" &&
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

func (v *AbilityVerifier) queryEntityArgsMap() map[string]interface{} {
	expMap := map[string]interface{}{}
	queryEntityExpressions := []string{}
	for _, ability := range v.abilities {
		if ability.ColumnUuid == "" &&
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

func (v *AbilityVerifier) queryRolesAbilities() {
	abilities := repository.Query(model.GlobalModel.Graph.GetEntityByUuid(consts.ABILITY_UUID), repository.QueryArg{
		consts.ARG_WHERE: repository.QueryArg{
			"roleId": repository.QueryArg{
				consts.ARG_IN: v.roleIds,
			},
			"abilityType": repository.QueryArg{
				consts.ARG_EQ: v.abilityType,
			},
		},
	})

	for _, abilityMap := range abilities {
		var ability common.Ability
		err := mapstructure.Decode(abilityMap, &ability)
		if err != nil {
			panic(err.Error())
		}
		v.abilities = append(v.abilities, &ability)
	}
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
