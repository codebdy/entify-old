package domain

import "rxdrag.com/entity-engine/model/meta"

type Attribute struct {
	meta.AttributeMeta
	Class *Class
}
