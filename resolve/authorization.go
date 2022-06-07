package resolve

import (
	"github.com/graphql-go/graphql"
	"github.com/mitchellh/mapstructure"
	"rxdrag.com/entify/authorization"
	"rxdrag.com/entify/common"
	"rxdrag.com/entify/consts"
	"rxdrag.com/entify/model"
	"rxdrag.com/entify/repository"
)

func InitAbilityVerifier(v *authorization.AbilityVerifier, p graphql.ResolveParams, entityUuids []string, abilityType string) {
	me := authorization.ParseContextValues(p).Me
	v.Me = me
	if me != nil {
		for i := range me.Roles {
			v.RoleIds = append(v.RoleIds, me.Roles[i].Id)
		}
	} else {
		v.RoleIds = append(v.RoleIds, consts.GUEST_ROLE_ID)
	}

	v.AbilityType = abilityType

	queryRolesAbilities(v, entityUuids)
}

func queryRolesAbilities(v *authorization.AbilityVerifier, entityUuids []string) {
	abilities := repository.QueryEntity(model.GlobalModel.Graph.GetEntityByUuid(consts.ABILITY_UUID), repository.QueryArg{
		consts.ARG_WHERE: repository.QueryArg{
			"roleId": repository.QueryArg{
				consts.ARG_IN: v.RoleIds,
			},
			"abilityType": repository.QueryArg{
				consts.ARG_EQ: v.AbilityType,
			},
			"entityUuid": repository.QueryArg{
				consts.ARG_IN: entityUuids,
			},
		},
	})

	for _, abilityMap := range abilities {
		var ability common.Ability
		err := mapstructure.Decode(abilityMap, &ability)
		if err != nil {
			panic(err.Error())
		}
		v.Abilities = append(v.Abilities, &ability)
	}
}
