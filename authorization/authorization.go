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

func New(p graphql.ResolveParams, entityUuid string, abilityType string) *AbilityVerifier {
	verifier := AbilityVerifier{}
	me := common.ParseContextValues(p).Me
	if me != nil {
		for i := range me.Roles {
			verifier.roleIds = append(verifier.roleIds, me.Roles[i].Id)
		}
	} else {
		verifier.roleIds = append(verifier.roleIds, consts.GUEST_ROLE_ID)
	}

	verifier.abilityType = abilityType

	verifier.queryRolesAbilities()
	verifier.parseQueryUserMap()

	return &verifier
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
