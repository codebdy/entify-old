package graph

type Noder interface {
	Uuid() string
	InnerId() uint64
	Name() string
	TableName() string
	Description() string
	isInterface() bool
	Interface() *Interface
	Entity() *Entity
	AddAssociation(a *Association)
	//Attributes() []*Attribute
	AllAssociations() []*Association
	AllAttributes() []*Attribute
	AllMethods() []*Method
	AllAttributeNames() []string
	GetAttributeByName(name string) *Attribute
	IsEmperty() bool
}
