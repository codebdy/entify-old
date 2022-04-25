package graph

import (
	"fmt"

	"rxdrag.com/entity-engine/consts"
)

const PREFIX_T string = "t"

type QueryArg = map[string]interface{}

type Ider interface {
	CreateId() int
}

type ArgAssociation struct {
	Association *Association
	ArgClass    *ArgClass
}

type ArgEntity struct {
	FromClass *ArgClass
	Id        int
	Entity    *Entity
}

type ArgClass struct {
	Noder        Noder
	Associations []*ArgAssociation
	Ider         Ider
	Children     []*ArgEntity
}

func NewArgClass(noder Noder, ider Ider) *ArgClass {
	var entities []*ArgEntity
	argClass := &ArgClass{
		Noder: noder,
		Ider:  ider,
	}
	if noder.IsInterface() {
		children := noder.Interface().Children
		for i := range children {
			entities = append(entities, &ArgEntity{
				Id:        ider.CreateId(),
				Entity:    children[i],
				FromClass: argClass,
			})
		}
	} else {
		entities = append(entities, &ArgEntity{
			Id:     ider.CreateId(),
			Entity: noder.Entity(),
		})
	}
	argClass.Children = entities
	return argClass
}

func (a *ArgClass) GetWithMakeAssociation(name string) *ArgAssociation {
	for i := range a.Associations {
		if a.Associations[i].Association.Name() == name {
			return a.Associations[i]
		}
	}
	allAssociations := a.Noder.AllAssociations()
	for i := range allAssociations {
		if allAssociations[i].Name() == name {
			asso := &ArgAssociation{
				Association: allAssociations[i],
				ArgClass:    NewArgClass(allAssociations[i].TypeClass(), a.Ider),
			}

			a.Associations = append(a.Associations, asso)

			return asso
		}
	}
	panic("Can not find entity association:" + a.Noder.Name() + "." + name)
}

func (e *ArgEntity) Alise() string {
	return fmt.Sprintf("%s%d", PREFIX_T, e.Id)
}

func (a *ArgAssociation) GetTypeEntity(uuid string) *ArgEntity {
	entities := a.ArgClass.Children
	for i := range entities {
		if entities[i].Entity.Uuid() == uuid {
			return entities[i]
		}
	}

	panic("Can not find association entity by uuid")
}

func BuldArgClass(noder Noder, where interface{}, ider Ider) *ArgClass {
	rootClass := NewArgClass(noder, ider)
	if where != nil {
		if whereMap, ok := where.(QueryArg); ok {
			buildWhereClass(rootClass, whereMap)
		}
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
			association := cls.Noder.GetAssociationByName(key)
			if association != nil {
				argAssociation := cls.GetWithMakeAssociation(key)
				if subWhere, ok := value.(QueryArg); ok {
					buildWhereClass(argAssociation.ArgClass, subWhere)
				}
			}
			break
		}
	}
}
