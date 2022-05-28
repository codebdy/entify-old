package resolve

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entify/consts"
)

func weaveAuthInArgs(p graphql.ResolveParams) {
	var roleIds []string
	me := ContextValues(p).Me
	if me != nil {
		for i := range me.Roles {
			roleIds = append(roleIds, me.Roles[i].Id)
		}
	} else {
		roleIds = append(roleIds, consts.GUEST_ROLE_ID)
	}
}
