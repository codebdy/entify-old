package domain

type Class struct {
	Uuid         string
	Type         string
	Associations map[string]*Association
	Attributes   []*Attribute
	Methods      []*Method
	Parents      []*Class
	Children     []*Class
}
