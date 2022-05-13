package utils

import "rxdrag.com/entify/config"

const (
	SERVICE_BITS   = 52
	ENTITY_ID_BITS = 32
)

func EncodeBaseId(entityInnerId uint64) uint64 {
	return uint64(config.ServiceId())<<SERVICE_BITS + entityInnerId<<ENTITY_ID_BITS
}

func DecodeEntityInnerId(id uint64) uint64 {
	return (id - uint64(config.ServiceId())<<SERVICE_BITS) >> ENTITY_ID_BITS
}
