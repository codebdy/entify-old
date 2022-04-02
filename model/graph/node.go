package graph

type Node interface {
	Uuid() string
	Name() string
	Description() string
	isInterface() bool
	Interface() *Interface
	Entity() *Entity
}
