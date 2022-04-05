package graph

type Node interface {
	Uuid() string
	Name() string
	TableName() string
	Description() string
	isInterface() bool
	Interface() *Interface
	Entity() *Entity
	AddAssociation(a *Association)
	Attributes() []*Attribute
	Associations() []*Association
	AllAttributeNames() []string
	GetAttributeByName(name string) *Attribute
	IsEmperty() bool
}
