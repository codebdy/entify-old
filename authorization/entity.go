package authorization

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entify/consts"
)

func WeaveAuthInArgs(p graphql.ResolveParams, entityUuid string) {
	var roleIds []string
	me := ParseContextValues(p).Me
	if me != nil {
		for i := range me.Roles {
			roleIds = append(roleIds, me.Roles[i].Id)
		}
	} else {
		roleIds = append(roleIds, consts.GUEST_ROLE_ID)
	}

}
