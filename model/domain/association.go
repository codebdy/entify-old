package domain

type Association struct {
	Name        string
	Relation    *Relation
	OwnerClass  *Class
	TypeClass   *Class
	Description string
}
