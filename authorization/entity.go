package authorization

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entify/common"
	"rxdrag.com/entify/consts"
)

func WeaveAuthInArgs(p graphql.ResolveParams, classUuid string) {
	var roleIds []string
	me := common.ParseContextValues(p).Me
	if me != nil {
		for i := range me.Roles {
			roleIds = append(roleIds, me.Roles[i].Id)
		}
	} else {
		roleIds = append(roleIds, consts.GUEST_ROLE_ID)
	}

}
