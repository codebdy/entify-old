package db

import (
	"database/sql"
	"database/sql/driver"
)

var (
	_ sql.Scanner  = (*NullUint64)(nil)
	_ driver.Value = (*NullUint64)(nil)
)

// NullUint64 represents an uint64 that may be null.
// NullUint64 implements the Scanner interface so
// it can be used as a scan destination, similar to NullString.
type NullUint64 struct {
	Uint64 uint64
	Valid  bool // Valid is true if Uint64 is not NULL
	i      struct {
		sql.NullInt64
	}
}

// Scan implements the Scanner interface.
func (n *NullUint64) Scan(value interface{}) error {
	if err := n.i.NullInt64.Scan(value); err != nil {
		return err
	}
	n.Uint64 = uint64(n.i.NullInt64.Int64)
	n.Valid = n.i.NullInt64.Valid
	return nil
}

// Value implements the driver Valuer interface.
func (n NullUint64) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return int64(n.Uint64), nil
}
