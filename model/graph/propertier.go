package graph

type Propertier interface {
	GetName() string
	GetType() string
	GetEumnType() *Enum
	GetEnityType() *Entity
}
