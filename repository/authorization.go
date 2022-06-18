package repository

import (
	"encoding/json"

	"github.com/graphql-go/graphql"
	"github.com/mitchellh/mapstructure"
	"rxdrag.com/entify/authcontext"
	"rxdrag.com/entify/common"
	"rxdrag.com/entify/consts"
	"rxdrag.com/entify/model"
	"rxdrag.com/entify/model/graph"
	"rxdrag.com/entify/model/meta"
)

type AbilityVerifier struct {
	me        *common.User
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
	v.me = me
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

	if v.me != nil {
		return v.me.IsSupper
	}

	return false
}

func (v *AbilityVerifier) IsDemo() bool {
	if v.me != nil {
		return v.me.IsDemo
	}

	return false
}

func (v *AbilityVerifier) WeaveAuthInArgs(entityUuid string, args interface{}) interface{} {
	if v.IsSupper() || v.IsDemo() {
		return args
	}

	var rootAnd []map[string]interface{}

	if args == nil {
		rootAnd = []map[string]interface{}{}
	} else {
		argsMap := args.(map[string]interface{})
		if argsMap[consts.ARG_AND] == nil {
			rootAnd = []map[string]interface{}{}
		} else {
			rootAnd = argsMap[consts.ARG_AND].([]map[string]interface{})
		}
	}

	// if len(v.Abilities) == 0 && !v.IsSupper() && !v.IsDemo() {
	// 	rootAnd = append(rootAnd, map[string]interface{}{
	// 		consts.ID: map[string]interface{}{
	// 			consts.ARG_EQ: 0,
	// 		},
	// 	})

	// 	return map[string]interface{}{
	// 		consts.ARG_AND: rootAnd,
	// 	}
	// }

	expArg := v.queryEntityArgsMap(entityUuid)
	if len(expArg) > 0 {
		rootAnd = append(rootAnd, expArg)
	}

	if args == nil {
		return map[string]interface{}{
			consts.ARG_AND: rootAnd,
		}
	} else {
		argsMap := args.(map[string]interface{})
		argsMap[consts.ARG_AND] = rootAnd
		return argsMap
	}
}

func (v *AbilityVerifier) CanReadEntity(entityUuid string) bool {
	if v.IsSupper() || v.IsDemo() {
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
	abilities := QueryEntity(model.GlobalModel.Graph.GetEntityByUuid(consts.ABILITY_UUID), graph.QueryArg{
		consts.ARG_WHERE: graph.QueryArg{
			"roleId": graph.QueryArg{
				consts.ARG_IN: v.RoleIds,
			},
			// "abilityType": QueryArg{
			// 	consts.ARG_EQ: v.AbilityType,
			// },
			"entityUuid": graph.QueryArg{
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
