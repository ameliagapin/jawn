package jawn

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

func (n NullFloat64) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.MarshalValue(n.Float64)
}

func (n *NullFloat64) UnmarshalBSONValue(bType bsontype.Type, bBytes []byte) error {
	rawValue := bson.RawValue{Type: bType, Value: bBytes}

	if err := rawValue.Unmarshal(&n.Float64); err == nil {
		n.Valid = true
		return nil
	}

	return fmt.Errorf("Cannot unmarshal %s into jawn.Float64", bType)
}

func (n NullFloat32) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.MarshalValue(n.Float32)
}

func (n *NullFloat32) UnmarshalBSONValue(bType bsontype.Type, bBytes []byte) error {
	rawValue := bson.RawValue{Type: bType, Value: bBytes}
	if err := rawValue.Unmarshal(&n.Float32); err == nil {
		n.Valid = true
		return nil
	}

	return fmt.Errorf("Cannot unmarshal %s into jawn.Float32", bType)
}

func (n NullInt) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.MarshalValue(n.Int)
}

func (n *NullInt) UnmarshalBSONValue(bType bsontype.Type, bBytes []byte) error {
	rawValue := bson.RawValue{Type: bType, Value: bBytes}

	if err := rawValue.Unmarshal(&n.Int); err == nil {
		n.Valid = true
		return nil
	}

	return fmt.Errorf("Cannot unmarshal %s into jawn.Int", bType)
}

func (n NullInt64) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.MarshalValue(n.Int64)
}

func (n *NullInt64) UnmarshalBSONValue(bType bsontype.Type, bBytes []byte) error {
	rawValue := bson.RawValue{Type: bType, Value: bBytes}
	if err := rawValue.Unmarshal(&n.Int64); err == nil {
		n.Valid = true
		return nil
	}

	return fmt.Errorf("Cannot unmarshal %s into jawn.Int64", bType)
}

func (n NullInt32) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.MarshalValue(n.Int32)
}

func (n *NullInt32) UnmarshalBSONValue(bType bsontype.Type, bBytes []byte) error {
	rawValue := bson.RawValue{Type: bType, Value: bBytes}

	if err := rawValue.Unmarshal(&n.Int32); err == nil {
		n.Valid = true
		return nil
	}

	return fmt.Errorf("Cannot unmarshal %s into jawn.Int32", bType)
}

func (n NullInt8) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.MarshalValue(n.Int8)
}

func (n *NullInt8) UnmarshalBSONValue(bType bsontype.Type, bBytes []byte) error {
	rawValue := bson.RawValue{Type: bType, Value: bBytes}
	if err := rawValue.Unmarshal(&n.Int8); err == nil {
		n.Valid = true
		return nil
	}

	return fmt.Errorf("Cannot unmarshal %s into jawn.Int8", bType)
}

func (n NullInt16) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.MarshalValue(n.Int16)
}

func (n *NullInt16) UnmarshalBSONValue(bType bsontype.Type, bBytes []byte) error {
	rawValue := bson.RawValue{Type: bType, Value: bBytes}
	if err := rawValue.Unmarshal(&n.Int16); err == nil {
		n.Valid = true
		return nil
	}

	return fmt.Errorf("Cannot unmarshal %s into jawn.Int16", bType)
}

func (n NullBool) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.MarshalValue(n.Bool)
}

func (n *NullBool) UnmarshalBSONValue(bType bsontype.Type, bBytes []byte) error {
	rawValue := bson.RawValue{Type: bType, Value: bBytes}
	if err := rawValue.Unmarshal(&n.Bool); err == nil {
		n.Valid = true
		return nil
	}

	return fmt.Errorf("Cannot unmarshal %s into jawn.Bool", bType)
}

func (n NullString) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.MarshalValue(n.String)
}

func (n *NullString) UnmarshalBSONValue(bType bsontype.Type, bBytes []byte) error {
	rawValue := bson.RawValue{Type: bType, Value: bBytes}
	if err := rawValue.Unmarshal(&n.String); err == nil {
		n.Valid = true
		return nil
	}

	return fmt.Errorf("Cannot unmarshal %s into jawn.String", bType)
}

func (n NullTime) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.MarshalValue(n.Time)
}

func (n *NullTime) UnmarshalBSONValue(bType bsontype.Type, bBytes []byte) error {
	rawValue := bson.RawValue{Type: bType, Value: bBytes}
	if err := rawValue.Unmarshal(&n.Time); err == nil {
		n.Valid = true
		return nil
	}

	return fmt.Errorf("Cannot unmarshal %s into jawn.Time", bType)
}
