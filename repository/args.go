package repository

import "rxdrag.com/entity-engine/model/graph"

//path Entity Map

type ArgAssociation struct {
	association *graph.Association
	argClass    *ArgClass
}

type ArgClass struct {
	id           int
	noder        graph.Noder
	associations []*ArgAssociation
	con          *Connection
}

func (con *Connection) NewArgClass(noder graph.Noder) *ArgClass {
	return &ArgClass{
		id:    con.createId(),
		noder: noder,
		con:   con,
	}
}

func (a *ArgClass) GetAssociationByName(name string) *graph.Association {
	allAssociations := a.noder.AllAssociations()
	for i := range allAssociations {
		if allAssociations[i].Name() == name {
			return allAssociations[i]
		}
	}
	panic("Can not find entity association:" + a.noder.Name() + "." + name)
}
