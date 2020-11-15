package jawn

import (
	"bytes"
	"encoding/json"
	"time"
)

// MarshalJSON - marshaller for json
func (n NullBool) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(n.Bool)
}

// UnmarshalJSON - unmarshaller for json
func (n *NullBool) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		*n = NullBool{}
		return nil
	}

	var res bool

	err := json.Unmarshal(data, &res)
	if err == nil {
		*n = NullBool{Bool: res, Valid: true}
		return nil
	}

	var res2 struct {
		Bool  bool `json:"bool"`
		Valid bool `json:"valid"`
	}
	err = json.Unmarshal(data, &res2)
	if err == nil {
		*n = NullBool{Bool: res2.Bool, Valid: res2.Valid}
		return nil
	}
	return err
}

// MarshalJSON - marshaller for json
func (n NullFloat32) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(n.Float32)
}

// UnmarshalJSON - unmarshaller for json
func (n *NullFloat32) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		*n = NullFloat32{}
		return nil
	}

	var res float32

	err := json.Unmarshal(data, &res)
	if err == nil {
		*n = NullFloat32{Float32: res, Valid: true}
		return nil
	}

	var res2 struct {
		Float32 float32 `json:"float32"`
		Valid   bool    `json:"valid"`
	}
	err = json.Unmarshal(data, &res2)
	if err == nil {
		*n = NullFloat32{Float32: res2.Float32, Valid: res2.Valid}
		return nil
	}
	return err
}

// MarshalJSON - marshaller for json
func (n NullFloat64) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(n.Float64)
}

// UnmarshalJSON - unmarshaller for json
func (n *NullFloat64) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		*n = NullFloat64{}
		return nil
	}

	var res float64

	err := json.Unmarshal(data, &res)
	if err == nil {
		*n = NullFloat64{Float64: res, Valid: true}
		return nil
	}

	var res2 struct {
		Float64 float64 `json:"float64"`
		Valid   bool    `json:"valid"`
	}
	err = json.Unmarshal(data, &res2)
	if err == nil {
		*n = NullFloat64{Float64: res2.Float64, Valid: res2.Valid}
		return nil
	}
	return err
}

// MarshalJSON - marshaller for json
func (n NullInt) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(n.Int)
}

// UnmarshalJSON - unmarshaller for json
func (n *NullInt) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		*n = NullInt{}
		return nil
	}

	var res int

	err := json.Unmarshal(data, &res)
	if err == nil {
		*n = NullInt{Int: res, Valid: true}
		return nil
	}

	var res2 struct {
		Int   int  `json:"int"`
		Valid bool `json:"valid"`
	}
	err = json.Unmarshal(data, &res2)
	if err == nil {
		*n = NullInt{Int: res2.Int, Valid: res2.Valid}
		return nil
	}
	return err
}

// MarshalJSON - marshaller for json
func (n NullInt8) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(n.Int8)
}

// UnmarshalJSON - unmarshaller for json
func (n *NullInt8) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		*n = NullInt8{}
		return nil
	}

	var res int8

	err := json.Unmarshal(data, &res)
	if err == nil {
		*n = NullInt8{Int8: res, Valid: true}
		return nil
	}

	var res2 struct {
		Int8  int8 `json:"int8"`
		Valid bool `json:"valid"`
	}
	err = json.Unmarshal(data, &res2)
	if err == nil {
		*n = NullInt8{Int8: res2.Int8, Valid: res2.Valid}
		return nil
	}
	return err
}

// MarshalJSON - marshaller for json
func (n NullInt16) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(n.Int16)
}

// UnmarshalJSON - unmarshaller for json
func (n *NullInt16) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		*n = NullInt16{}
		return nil
	}

	var res int16

	err := json.Unmarshal(data, &res)
	if err == nil {
		*n = NullInt16{Int16: res, Valid: true}
		return nil
	}

	var res2 struct {
		Int16 int16 `json:"int16"`
		Valid bool  `json:"valid"`
	}
	err = json.Unmarshal(data, &res2)
	if err == nil {
		*n = NullInt16{Int16: res2.Int16, Valid: res2.Valid}
		return nil
	}
	return err
}

// MarshalJSON - marshaller for json
func (n NullInt32) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(n.Int32)
}

// UnmarshalJSON - unmarshaller for json
func (n *NullInt32) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		*n = NullInt32{}
		return nil
	}

	var res int32

	err := json.Unmarshal(data, &res)
	if err == nil {
		*n = NullInt32{Int32: res, Valid: true}
		return nil
	}

	var res2 struct {
		Int32 int32 `json:"int32"`
		Valid bool  `json:"valid"`
	}
	err = json.Unmarshal(data, &res2)
	if err == nil {
		*n = NullInt32{Int32: res2.Int32, Valid: res2.Valid}
		return nil
	}
	return err
}

// MarshalJSON - marshaller for json
func (n NullInt64) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(n.Int64)
}

// UnmarshalJSON - unmarshaller for json
func (n *NullInt64) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		*n = NullInt64{}
		return nil
	}

	var res int64

	err := json.Unmarshal(data, &res)
	if err == nil {
		*n = NullInt64{Int64: res, Valid: true}
		return nil
	}

	var res2 struct {
		Int64 int64 `json:"int64"`
		Valid bool  `json:"valid"`
	}
	err = json.Unmarshal(data, &res2)
	if err == nil {
		*n = NullInt64{Int64: res2.Int64, Valid: res2.Valid}
		return nil
	}
	return err
}

// MarshalJSON - marshaller for json
func (n NullString) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(n.String)
}

// UnmarshalJSON - unmarshaller for json
func (n *NullString) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		*n = NullString{}
		return nil
	}

	var res string

	err := json.Unmarshal(data, &res)
	if err == nil {
		*n = NullString{String: res, Valid: true}
		return nil
	}

	var res2 struct {
		String string `json:"string"`
		Valid  bool   `json:"valid"`
	}
	err = json.Unmarshal(data, &res2)
	if err == nil {
		*n = NullString{String: res2.String, Valid: res2.Valid}
		return nil
	}
	return err
}

// MarshalJSON - marshaller for json
func (n NullTime) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(n.Time.Format(time.RFC3339Nano))
}

// UnmarshalJSON - unmarshaller for json
func (n *NullTime) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		*n = NullTime{}
		return nil
	}

	var res time.Time

	err := json.Unmarshal(data, &res)
	if err == nil {
		*n = NullTime{Time: res, Valid: true}
		return nil
	}

	var res2 struct {
		Time  time.Time `json:"time"`
		Valid bool      `json:"valid"`
	}
	err = json.Unmarshal(data, &res2)
	if err == nil {
		*n = NullTime{Time: res2.Time, Valid: res2.Valid}
		return nil
	}
	return err
}

// MarshalJSON - marshaller for json
func (n NullDuration) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(n.Duration)
}

// UnmarshalJSON - unmarshaller for json
func (n *NullDuration) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		*n = NullDuration{}
		return nil
	}

	var res time.Duration

	err := json.Unmarshal(data, &res)
	if err == nil {
		*n = NullDuration{Duration: res, Valid: true}
		return nil
	}

	var res2 struct {
		Duration time.Duration `json:"duration"`
		Valid    bool          `json:"valid"`
	}
	err = json.Unmarshal(data, &res2)
	if err == nil {
		*n = NullDuration{Duration: res2.Duration, Valid: res2.Valid}
		return nil
	}
	return err
}
