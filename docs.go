/*
Package jawn - Zero allocation Nullable structures in one library with handy conversion functions,
marshallers and unmarshallers.

Features:
- short name "jawn"
- handy conversion functions
- select which marshalers you want and limit dependencies to only those you actually need

Supported types:
- bool
- float32
- float64
- int
- int8
- int16
- int32
- int64
- string
- time.Time
- more types will be added at necessary

Supported marshallers:
- Standart JSON
- jsoniter
- easyjson
- Scylla and Cassandra. Compatible with gocql
- SQL

Usage

Simply create struct field or variable with one of the exported types and use it without any changes to external API.

JSON input/output will be converted to null or non null values. Scylla and Cassandra will use wire format compatible
with gocql.

	var data struct {
		Code jawn.NullString `json:"code"`
	}

	b, err := jsoniter.Marshal(data)
	if err != nil {
		panic(err)
	}

	// {"code":null}
	fmt.Println(string(b))

	data.Code = jawn.String("1")
	// Or
	// data.Code = jawn.NullString{String: "1", Valid: true}

	b, err = jsoniter.Marshal(data)
	if err != nil {
		panic(err)
	}

	// {"code":"1"}
	fmt.Println(string(b))
*/
package jawn
