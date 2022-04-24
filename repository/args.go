package repository

import "rxdrag.com/entity-engine/model/graph"

//path Entity Map

type ArgAssociation struct {
	association *graph.Association
	argClass    *ArgClass
}

type ArgEntity struct {
	id     int
	entity *graph.Entity
}

type ArgClass struct {
	noder        graph.Noder
	associations []*ArgAssociation
	con          *Connection
	children     []*ArgEntity
}

func (con *Connection) NewArgClass(noder graph.Noder) *ArgClass {
	var entities []*ArgEntity
	if noder.IsInterface() {
		children := noder.Interface().Children
		for i := range children {
			entities = append(entities, &ArgEntity{
				id:     con.createId(),
				entity: children[i],
			})
		}
	} else {
		entities = append(entities, &ArgEntity{
			id:     con.createId(),
			entity: noder.Entity(),
		})
	}
	return &ArgClass{
		noder:    noder,
		con:      con,
		children: entities,
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
