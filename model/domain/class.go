package domain

type Class struct {
	Type         string
	Associations map[string]*Association
	Attributes   []*Attribute
	Methods      []*Method
	Parents      []*Class
	Children     []*Class
}
