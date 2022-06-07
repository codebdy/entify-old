package authorization

import (
	"encoding/json"

	"rxdrag.com/entify/common"
	"rxdrag.com/entify/consts"
	"rxdrag.com/entify/model/meta"
)

type AbilityVerifier struct {
	Me          *common.User
	RoleIds     []string
	AbilityType string
	Abilities   []*common.Ability
	// expression Key : 从Auth模块返回的结果
	QueryUserCache map[string][]common.User
}

func NewVerifier() *AbilityVerifier {
	verifier := AbilityVerifier{}

	return &verifier
}

func (v *AbilityVerifier) WeaveAuthInArgs(entityUuid string, args map[string]interface{}) map[string]interface{} {
	var rootAnd []map[string]interface{}
	if args[consts.ARG_ADD] == nil {
		rootAnd = []map[string]interface{}{}
	} else {
		rootAnd = args[consts.ARG_ADD].([]map[string]interface{})
	}

	expArg := v.queryEntityArgsMap(entityUuid)
	if len(expArg) > 0 {
		rootAnd = append(rootAnd, expArg)
	}

	args[consts.ARG_ADD] = rootAnd
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
