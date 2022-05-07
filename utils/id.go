package utils

import "rxdrag.com/entify/config"

const (
	SERVICE_BITS   = 52
	ENTITY_ID_BITS = 32
)

func EncodeBaseId(entityInnerId uint64) uint64 {
	return config.SERVICE_ID<<SERVICE_BITS + entityInnerId<<ENTITY_ID_BITS
}

func DecodeEntityInnerId(id uint64) uint64 {
	return (id - config.SERVICE_ID<<SERVICE_BITS) >> ENTITY_ID_BITS
}
