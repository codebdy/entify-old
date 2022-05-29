package authorization

import (
	"github.com/graphql-go/graphql"
	"github.com/mitchellh/mapstructure"
	"rxdrag.com/entify/common"
	"rxdrag.com/entify/consts"
	"rxdrag.com/entify/model"
	"rxdrag.com/entify/repository"
)

type AbilityVerifier struct {
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
	if me != nil {
		for i := range me.Roles {
			v.roleIds = append(v.roleIds, me.Roles[i].Id)
		}
	} else {
		v.roleIds = append(v.roleIds, consts.GUEST_ROLE_ID)
	}

	v.abilityType = abilityType

	v.queryRolesAbilities()
	v.parseQueryUserMap()
}

func (v *AbilityVerifier) WeaveAuthInArgs(args map[string]interface{}) {

}

func (v *AbilityVerifier) EntityMutationCan(entityData map[string]interface{}) bool {
	return false
}

func (v *AbilityVerifier) FieldCan(entityData map[string]interface{}) bool {
	return false
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

func (v *AbilityVerifier) parseQueryUserMap() {

}
