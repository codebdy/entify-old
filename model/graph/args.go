package graph

import (
	"fmt"

	"rxdrag.com/entify/consts"
)

const PREFIX_T string = "t"

type QueryArg = map[string]interface{}

type Ider interface {
	CreateId() int
}

type ArgAssociation struct {
	Association *Association
	ArgEntities []*ArgEntity
}

type ArgEntity struct {
	Id             int
	Entity         *Entity
	Associations   []*ArgAssociation
	ExpressionArgs map[string]interface{}
}

func argEntitiesFromAssociation(associ *Association, ider Ider) []*ArgEntity {
	var entities []*ArgEntity
	noder := associ.TypeClass()
	if noder.IsInterface() {
		children := noder.Interface().Children
		for i := range children {
			entities = append(entities, &ArgEntity{
				Id:     ider.CreateId(),
				Entity: children[i],
			})
		}
	} else {
		entities = append(entities, &ArgEntity{
			Id:     ider.CreateId(),
			Entity: noder.Entity(),
		})
	}
	return entities
}

func (a *ArgEntity) GetAssociation(name string) *ArgAssociation {
	for i := range a.Associations {
		if a.Associations[i].Association.Name() == name {
			return a.Associations[i]
		}
	}
	panic("Can not find entity association:" + a.Entity.NameWithPartial() + "." + name)
}

func (a *ArgEntity) GetWithMakeAssociation(name string, ider Ider) *ArgAssociation {
	for i := range a.Associations {
		if a.Associations[i].Association.Name() == name {
			return a.Associations[i]
		}
	}
	allAssociations := a.Entity.AllAssociations()
	for i := range allAssociations {
		if allAssociations[i].Name() == name {
			asso := &ArgAssociation{
				Association: allAssociations[i],
				ArgEntities: argEntitiesFromAssociation(allAssociations[i], ider),
			}

			a.Associations = append(a.Associations, asso)

			return asso
		}
	}
	panic("Can not find entity association:" + a.Entity.NameWithPartial() + "." + name)
}

func (e *ArgEntity) Alise() string {
	return fmt.Sprintf("%s%d", PREFIX_T, e.Id)
}

func (a *ArgAssociation) GetTypeEntity(uuid string) *ArgEntity {
	entities := a.ArgEntities
	for i := range entities {
		if entities[i].Entity.Uuid() == uuid {
			return entities[i]
		}
	}

	panic("Can not find association entity by uuid")
}

func BuildArgEntity(entity *Entity, where interface{}, ider Ider) *ArgEntity {
	rootEntity := &ArgEntity{
		Id:     ider.CreateId(),
		Entity: entity,
	}
	if where != nil {
		if whereMap, ok := where.(QueryArg); ok {
			buildWhereEntity(rootEntity, whereMap, ider)
		}
	}
	return rootEntity
}

func buildWhereEntity(argEntity *ArgEntity, where QueryArg, ider Ider) {
	for key, value := range where {
		switch key {
		case consts.ARG_AND, consts.ARG_NOT, consts.ARG_OR:
			if subWhere, ok := value.(QueryArg); ok {
				buildWhereEntity(argEntity, subWhere, ider)
			}
			break
		default:
			association := argEntity.Entity.GetAssociationByName(key)
			if association != nil {
				argAssociation := argEntity.GetWithMakeAssociation(key, ider)
				if subWhere, ok := value.(QueryArg); ok {
					for i := range argAssociation.ArgEntities {
						buildWhereEntity(argAssociation.ArgEntities[i], subWhere, ider)
					}
				}
			}
			break
		}
	}
}
