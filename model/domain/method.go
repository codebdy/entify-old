package domain

import "rxdrag.com/entity-engine/model/meta"

type Method struct {
	meta.MethodMeta
	Class *Class
}

func NewMethod(m *meta.MethodMeta, c *Class) *Method {
	return &Method{
		MethodMeta: *m,
		Class:      c,
	}
}
