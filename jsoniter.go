package nan

import (
	"fmt"
	"time"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

func init() {
	jsoniter.RegisterTypeDecoderFunc(
		"nan.NullBool",
		func(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
			if iter.ReadNil() {
				return
			}

			*((*NullBool)(ptr)) = NullBool{Bool: iter.ReadBool(), Valid: true}
		},
	)
	jsoniter.RegisterTypeEncoderFunc(
		"nan.NullBool",
		func(ptr unsafe.Pointer, stream *jsoniter.Stream) {
			t := *((*NullBool)(ptr))

			if !t.Valid {
				stream.WriteNil()
				return
			}

			stream.WriteBool(t.Bool)
		},
		nil,
	)

	jsoniter.RegisterTypeDecoderFunc(
		"nan.NullFloat32",
		func(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
			if iter.ReadNil() {
				return
			}

			*((*NullFloat32)(ptr)) = NullFloat32{Float32: iter.ReadFloat32(), Valid: true}
		},
	)
	jsoniter.RegisterTypeEncoderFunc(
		"nan.NullFloat32",
		func(ptr unsafe.Pointer, stream *jsoniter.Stream) {
			t := *((*NullFloat32)(ptr))

			if !t.Valid {
				stream.WriteNil()
				return
			}

			stream.WriteFloat32(t.Float32)
		},
		nil,
	)

	jsoniter.RegisterTypeDecoderFunc(
		"nan.NullFloat64",
		func(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
			if iter.ReadNil() {
				return
			}

			*((*NullFloat64)(ptr)) = NullFloat64{Float64: iter.ReadFloat64(), Valid: true}
		},
	)
	jsoniter.RegisterTypeEncoderFunc(
		"nan.NullFloat64",
		func(ptr unsafe.Pointer, stream *jsoniter.Stream) {
			t := *((*NullFloat64)(ptr))

			if !t.Valid {
				stream.WriteNil()
				return
			}

			stream.WriteFloat64(t.Float64)
		},
		nil,
	)

	jsoniter.RegisterTypeDecoderFunc(
		"nan.NullInt",
		func(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
			if iter.ReadNil() {
				return
			}

			*((*NullInt)(ptr)) = NullInt{Int: iter.ReadInt(), Valid: true}
		},
	)
	jsoniter.RegisterTypeEncoderFunc(
		"nan.NullInt",
		func(ptr unsafe.Pointer, stream *jsoniter.Stream) {
			t := *((*NullInt)(ptr))

			if !t.Valid {
				stream.WriteNil()
				return
			}

			stream.WriteInt(t.Int)
		},
		nil,
	)

	jsoniter.RegisterTypeDecoderFunc(
		"nan.NullInt8",
		func(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
			if iter.ReadNil() {
				return
			}

			*((*NullInt8)(ptr)) = NullInt8{Int8: iter.ReadInt8(), Valid: true}
		},
	)
	jsoniter.RegisterTypeEncoderFunc(
		"nan.NullInt8",
		func(ptr unsafe.Pointer, stream *jsoniter.Stream) {
			t := *((*NullInt8)(ptr))

			if !t.Valid {
				stream.WriteNil()
				return
			}

			stream.WriteInt8(t.Int8)
		},
		nil,
	)

	jsoniter.RegisterTypeDecoderFunc(
		"nan.NullInt16",
		func(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
			if iter.ReadNil() {
				return
			}

			*((*NullInt16)(ptr)) = NullInt16{Int16: iter.ReadInt16(), Valid: true}
		},
	)
	jsoniter.RegisterTypeEncoderFunc(
		"nan.NullInt16",
		func(ptr unsafe.Pointer, stream *jsoniter.Stream) {
			t := *((*NullInt16)(ptr))

			if !t.Valid {
				stream.WriteNil()
				return
			}

			stream.WriteInt16(t.Int16)
		},
		nil,
	)

	jsoniter.RegisterTypeDecoderFunc(
		"nan.NullInt32",
		func(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
			if iter.ReadNil() {
				return
			}

			*((*NullInt32)(ptr)) = NullInt32{Int32: iter.ReadInt32(), Valid: true}
		},
	)
	jsoniter.RegisterTypeEncoderFunc(
		"nan.NullInt32",
		func(ptr unsafe.Pointer, stream *jsoniter.Stream) {
			t := *((*NullInt32)(ptr))

			if !t.Valid {
				stream.WriteNil()
				return
			}

			stream.WriteInt32(t.Int32)
		},
		nil,
	)

	jsoniter.RegisterTypeDecoderFunc(
		"nan.NullInt64",
		func(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
			if iter.ReadNil() {
				return
			}

			*((*NullInt64)(ptr)) = NullInt64{Int64: iter.ReadInt64(), Valid: true}
		},
	)
	jsoniter.RegisterTypeEncoderFunc(
		"nan.NullInt64",
		func(ptr unsafe.Pointer, stream *jsoniter.Stream) {
			t := *((*NullInt64)(ptr))

			if !t.Valid {
				stream.WriteNil()
				return
			}

			stream.WriteInt64(t.Int64)
		},
		nil,
	)

	jsoniter.RegisterTypeDecoderFunc(
		"nan.NullString",
		func(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
			if iter.ReadNil() {
				return
			}

			*((*NullString)(ptr)) = NullString{String: iter.ReadString(), Valid: true}
		},
	)
	jsoniter.RegisterTypeEncoderFunc(
		"nan.NullString",
		func(ptr unsafe.Pointer, stream *jsoniter.Stream) {
			t := *((*NullString)(ptr))

			if !t.Valid {
				stream.WriteNil()
				return
			}

			stream.WriteString(t.String)
		},
		nil,
	)

	jsoniter.RegisterTypeDecoderFunc(
		"nan.NullTime",
		func(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
			if iter.ReadNil() {
				return
			}

			val := iter.ReadString()
			if val == "" {
				iter.ReportError("NullTime", "invalid value for time decoding \"\"")
				return
			}

			tm, err := time.Parse(time.RFC3339Nano, val)
			if err != nil {
				msg := fmt.Sprintf("%v: invalid value for time decoding \"%s\"", err, val)
				iter.ReportError("NullTime", msg)
				return
			}

			*((*NullTime)(ptr)) = NullTime{Time: tm, Valid: true}
		},
	)
	jsoniter.RegisterTypeEncoderFunc(
		"nan.NullTime",
		func(ptr unsafe.Pointer, stream *jsoniter.Stream) {
			t := *((*NullTime)(ptr))

			if !t.Valid {
				stream.WriteNil()
				return
			}

			stream.WriteString(t.Time.Format(time.RFC3339Nano))
		},
		nil,
	)
}
