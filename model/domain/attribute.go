package domain

import "rxdrag.com/entity-engine/model/meta"

type Attribute struct {
	meta.AttributeMeta
	Class *Class
}

func NewAttribute(a *meta.AttributeMeta, c *Class) *Attribute {
	return &Attribute{
		AttributeMeta: *a,
		Class:         c,
	}
}
