package jawn

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/bsontype"
)

//go:generate pkger -o cmd/jawn

// Validator is implemented by all jawn types and returns Valid field
type Validator interface {
	IsValid() bool
}

type Zeroer interface {
	IsZero() bool
}

type Mashaler interface {
	MarshalBSONValue() (bsontype.Type, []byte, error)
	UnmarshalBSONValue(bType bsontype.Type, bBytes []byte) error
}

// NullInt64 - nullable int64
//easyjson:skip
type NullInt64 struct {
	Int64 int64
	Valid bool // Valid is true if Int64 is not NULL
}

func (n NullInt64) IsValid() bool {
	return n.Valid
}

func (n NullInt64) IsZero() bool {
	return !n.Valid
}

func (n NullInt64) IsDefined() bool {
	return n.Valid
}

// NullInt32 - nullable int32
//easyjson:skip
type NullInt32 struct {
	Int32 int32
	Valid bool // Valid is true if Int32 is not NULL
}

func (n NullInt32) IsValid() bool {
	return n.Valid
}

func (n NullInt32) IsZero() bool {
	return !n.Valid
}

func (n NullInt32) IsDefined() bool {
	return n.Valid
}

// NullInt16 - nullable int16
//easyjson:skip
type NullInt16 struct {
	Int16 int16
	Valid bool // Valid is true if Int16 is not NULL
}

func (n NullInt16) IsValid() bool {
	return n.Valid
}

func (n NullInt16) IsZero() bool {
	return !n.Valid
}

func (n NullInt16) IsDefined() bool {
	return n.Valid
}

// NullInt8 - nullable int8
//easyjson:skip
type NullInt8 struct {
	Int8  int8
	Valid bool // Valid is true if Int8 is not NULL
}

func (n NullInt8) IsValid() bool {
	return n.Valid
}

func (n NullInt8) IsZero() bool {
	return !n.Valid
}

func (n NullInt8) IsDefined() bool {
	return n.Valid
}

// NullInt - nullable int
//easyjson:skip
type NullInt struct {
	Int   int
	Valid bool // Valid is true if Int8 is not NULL
}

func (n NullInt) IsValid() bool {
	return n.Valid
}

func (n NullInt) IsZero() bool {
	return !n.Valid
}

func (n NullInt) IsDefined() bool {
	return n.Valid
}

// NullTime - nullable time.Time
//easyjson:skip
type NullTime struct {
	Time  time.Time
	Valid bool // Valid is true if Time is not NULL
}

func (n NullTime) IsValid() bool {
	return n.Valid
}

func (n NullTime) IsZero() bool {
	return !n.Valid
}

func (n NullTime) IsDefined() bool {
	return n.Valid
}

// NullDuration - nullable time.Duration
//easyjson:skip
type NullDuration struct {
	Duration time.Duration
	Valid    bool // Valid is true if Duration is not NULL
}

func (n NullDuration) IsValid() bool {
	return n.Valid
}

func (n NullDuration) IsZero() bool {
	return !n.Valid
}

func (n NullDuration) IsDefined() bool {
	return n.Valid
}

// NullString - nullable string
//easyjson:skip
type NullString struct {
	String string
	Valid  bool // Valid is true if String is not NULL
}

func (n NullString) IsValid() bool {
	return n.Valid
}

func (n NullString) IsZero() bool {
	return !n.Valid
}

func (n NullString) IsDefined() bool {
	return n.Valid
}

// NullBool - nullable bool
//easyjson:skip
type NullBool struct {
	Bool  bool
	Valid bool // Valid is true if Bool is not NULL
}

func (n NullBool) IsValid() bool {
	return n.Valid
}

func (n NullBool) IsZero() bool {
	return !n.Valid
}

func (n NullBool) IsDefined() bool {
	return n.Valid
}

// NullFloat64 - nullable float64
//easyjson:skip
type NullFloat64 struct {
	Float64 float64
	Valid   bool // Valid is true if Float64 is not NULL
}

func (n NullFloat64) IsValid() bool {
	return n.Valid
}

func (n NullFloat64) IsZero() bool {
	return !n.Valid
}

func (n NullFloat64) IsDefined() bool {
	return n.Valid
}

// NullFloat32 - nullable float32
//easyjson:skip
type NullFloat32 struct {
	Float32 float32
	Valid   bool // Valid is true if Float32 is not NULL
}

func (n NullFloat32) IsValid() bool {
	return n.Valid
}

func (n NullFloat32) IsZero() bool {
	return !n.Valid
}

func (n NullFloat32) IsDefined() bool {
	return n.Valid
}

// type needed for template for custom types generator
//easyjson:skip
type initialTemplateType string
