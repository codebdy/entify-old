package repository

import (
	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/model/graph"
)

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

func (a *ArgClass) GetWithMakeAssociation(name string) *ArgAssociation {
	for i := range a.associations {
		if a.associations[i].association.Name() == name {
			return a.associations[i]
		}
	}
	allAssociations := a.noder.AllAssociations()
	for i := range allAssociations {
		if allAssociations[i].Name() == name {
			asso := &ArgAssociation{
				association: allAssociations[i],
				argClass:    a.con.NewArgClass(allAssociations[i].TypeClass()),
			}

			a.associations = append(a.associations, asso)

			return asso
		}
	}
	panic("Can not find entity association:" + a.noder.Name() + "." + name)
}

func (con *Connection) buildWhereNodes(noder graph.Noder, where QueryArg) *ArgClass {
	rootClass := con.NewArgClass(noder)
	if where == nil {
		buildWhereClass(rootClass, where)
	}
	return rootClass
}

func buildWhereClass(cls *ArgClass, where QueryArg) {
	for key, value := range where {
		switch key {
		case consts.ARG_AND, consts.ARG_NOT, consts.ARG_OR:
			if subWhere, ok := value.(QueryArg); ok {
				buildWhereClass(cls, subWhere)
			}
			break
		default:
			association := cls.noder.GetAssociationByName(key)
			if association != nil {
				argAssociation := cls.GetWithMakeAssociation(key)
				if subWhere, ok := value.(QueryArg); ok {
					buildWhereClass(argAssociation.argClass, subWhere)
				}
			}
			break
		}
	}
}
