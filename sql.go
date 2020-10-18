package nan

// From https://github.com/golang/go/blob/go1.14.6/src/database/sql/convert.go
// sql.Null* types use these methods internally but they're not exposed.

import (
	"database/sql/driver"
	"time"
)

// Scan - scan value from sql driver
func (n *NullBool) Scan(value interface{}) error {
	if value == nil {
		n.Bool, n.Valid = false, false
		return nil
	}

	n.Valid = true

	return convertAssign(&n.Bool, value)
}

// Value implements the driver Valuer interface.
func (n NullBool) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}

	return n.Bool, nil
}

// Scan - scan value from sql driver
func (n *NullFloat32) Scan(value interface{}) error {
	if value == nil {
		n.Float32, n.Valid = 0, false
		return nil
	}

	n.Valid = true

	return convertAssign(&n.Float32, value)
}

// Value implements the driver Valuer interface.
func (n NullFloat32) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}

	return n.Float32, nil
}

// Scan - scan value from sql driver
func (n *NullFloat64) Scan(value interface{}) error {
	if value == nil {
		n.Float64, n.Valid = 0, false
		return nil
	}

	n.Valid = true

	return convertAssign(&n.Float64, value)
}

// Value implements the driver Valuer interface.
func (n NullFloat64) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}

	return n.Float64, nil
}

// Scan - scan value from sql driver
func (n *NullInt) Scan(value interface{}) error {
	if value == nil {
		n.Int, n.Valid = 0, false
		return nil
	}

	n.Valid = true

	return convertAssign(&n.Int, value)
}

// Value implements the driver Valuer interface.
func (n NullInt) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}

	return int64(n.Int), nil
}

// Scan - scan value from sql driver
func (n *NullInt8) Scan(value interface{}) error {
	if value == nil {
		n.Int8, n.Valid = 0, false
		return nil
	}

	n.Valid = true

	return convertAssign(&n.Int8, value)
}

// Value implements the driver Valuer interface.
func (n NullInt8) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}

	return int64(n.Int8), nil
}

// Scan - scan value from sql driver
func (n *NullInt16) Scan(value interface{}) error {
	if value == nil {
		n.Int16, n.Valid = 0, false
		return nil
	}

	n.Valid = true

	return convertAssign(&n.Int16, value)
}

// Value implements the driver Valuer interface.
func (n NullInt16) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}

	return int64(n.Int16), nil
}

// Scan - scan value from sql driver
func (n *NullInt32) Scan(value interface{}) error {
	if value == nil {
		n.Int32, n.Valid = 0, false
		return nil
	}

	n.Valid = true

	return convertAssign(&n.Int32, value)
}

// Value implements the driver Valuer interface.
func (n NullInt32) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}

	return int64(n.Int32), nil
}

// Scan - scan value from sql driver
func (n *NullInt64) Scan(value interface{}) error {
	if value == nil {
		n.Int64, n.Valid = 0, false
		return nil
	}

	n.Valid = true

	return convertAssign(&n.Int64, value)
}

// Value implements the driver Valuer interface.
func (n NullInt64) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}

	return n.Int64, nil
}

// Scan - scan value from sql driver
func (ns *NullString) Scan(value interface{}) error {
	if value == nil {
		ns.String, ns.Valid = "", false
		return nil
	}

	ns.Valid = true

	return convertAssign(&ns.String, value)
}

// Value implements the driver Valuer interface.
func (ns NullString) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}

	return ns.String, nil
}

// Scan - scan value from sql driver
func (n *NullTime) Scan(value interface{}) error {
	if value == nil {
		n.Time, n.Valid = time.Time{}, false
		return nil
	}

	n.Valid = true

	return convertAssign(&n.Time, value)
}

// Value implements the driver Valuer interface.
func (n NullTime) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}

	return n.Time, nil
}

func (n *initialTemplateType) Scan(value interface{}) error {
	// Function only for code validity
	return nil
}

func (n initialTemplateType) Value() (driver.Value, error) {
	// Function only for code validity
	return initialTemplateType(""), nil
}
