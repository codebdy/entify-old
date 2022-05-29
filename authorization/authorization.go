package authorization

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entify/common"
	"rxdrag.com/entify/consts"
)

type Expression = map[string]interface{}

type AbilityVerifier struct {
	roleIds []string
	// path: Expression
	queryUserMap map[string]Expression
}

func New(p graphql.ResolveParams, entityUuid string) *AbilityVerifier {
	verifier := AbilityVerifier{}
	me := common.ParseContextValues(p).Me
	if me != nil {
		for i := range me.Roles {
			verifier.roleIds = append(verifier.roleIds, me.Roles[i].Id)
		}
	} else {
		verifier.roleIds = append(verifier.roleIds, consts.GUEST_ROLE_ID)
	}

	return &verifier
}

func getUserAbilities(userId uint64, entityUuid string) {

}

func getGuestAbilities(entityUuid string) {

}

func isExpand(entityUuid string) {

}

func canReadEntity(entityUuid string, roleId uint64) (bool, string) {
	return false, ""
}
