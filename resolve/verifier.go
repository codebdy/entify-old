package resolve

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entify/model/graph"
	"rxdrag.com/entify/repository"
)

func makeEntityAbilityVerifier(p graphql.ResolveParams, entityUuid string) *repository.AbilityVerifier {
	verifier := repository.NewVerifier()

	verifier.Init(p, []string{entityUuid})

	// if !verifier.CanReadEntity() && !node.IsInterface() {
	// 	panic("No permission to read: " + node.Name())
	// }

	// args := verifier.WeaveAuthInArgs(inputArgs)
	return verifier
}

func makeInterfaceAbilityVerifier(p graphql.ResolveParams, intf *graph.Interface) *repository.AbilityVerifier {
	verifier := repository.NewVerifier()
	var uuids []string
	for i := range intf.Children {
		uuids = append(uuids, intf.Children[i].Uuid())
	}
	verifier.Init(p, uuids)
	return verifier
}

func makeAssociAbilityVerifier(p graphql.ResolveParams, association *graph.Association) *repository.AbilityVerifier {
	verifier := repository.NewVerifier()
	if association.TypeClass().IsInterface() {
		var uuids []string
		intf := association.TypeClass().Interface()
		for i := range intf.Children {
			uuids = append(uuids, intf.Children[i].Uuid())
		}
		verifier.Init(p, uuids)
	} else {
		verifier.Init(p, []string{association.TypeClass().Entity().Uuid()})
	}
	return verifier
}
