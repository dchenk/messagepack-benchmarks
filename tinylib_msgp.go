// This is how you would do it using github.com/tinylib/msgp
package mb

//go:generate msgp_tinylib

type Struct1B struct {
	Str          string            `msg:"str"`
	Slice        []byte            `msg:"slice"`
	MapStrUint16 map[string]uint16 `msg:"map_str_uint16"`
	InnerStruct  InnerStructB      `msg:"inner_struct"`
	Bt           uint8             `msg:"bt"`
	Int          int32             `msg:"int"`
	Int64        int64             `msg:"i64"`
}

type InnerStructB struct {
	Bool    bool    `msg:"bool"`
	Float64 float64 `msg:"f64"`
	Int8    int8    `msg:"i8"`
	Uint64  uint64  `msg:"u64"`
}

type Struct2B struct {
	Str          string            `msg:"str"`
	Slice        []byte            `msg:"slice"`
	MapStrUint64 map[string]uint64 `msg:"map_str_uint64"`
	InnerStruct  InnerStructB      `msg:"inner_struct"`
	Float64      float64           `msg:"float_64"`
	MapStrSlice  map[string][]int8 `msg:"map_str_slice"`
	InnerStruct2 InnerStruct2B     `msg:"inner_struct_2"`
}

type InnerStruct2B struct {
	SliceUint64 []uint64               `msg:"slice_uint64"`
	MapStrIntf  map[string]interface{} `msg:"map_str_intf"`
	Str2        string                 `msg:"string"`
	F32         float32                `msg:"f32"`
}

var Struct1B_inst = Struct1B{
	Str:   "hello",
	Slice: []byte{'a', 'b', 'c'},
	MapStrUint16: map[string]uint16{
		"abc": 23,
		"lasdhflahsdlfkhlasdf": 4,
		"c": 1356,
	},
	InnerStruct: InnerStructB{true, 213.513, 23, 7869},
	Int:         42,
	Int64:       -27823325,
}

var Struct2B_inst = Struct2B{
	Str:   "lk\n\nsdj flasjfdlkaj\nsdlfka\tjs\nldfk\"jal kas_ ad\"fsfka\tjs\nldfk\"jal kas_ ad\"fska\tjs\nldfk\"jal kas_ adk",
	Slice: []byte{0x01, 'a', 'b', 'c', 'Z', 0x03, 0x05},
	MapStrUint64: map[string]uint64{
		"abc":          23532525,
		"defghijkl":    1,
		"ghijkl":       98,
		"ghijklghijkl": 8000111333,
	},
	InnerStruct: InnerStructB{false, -132.88889, 104, 24},
	Float64:     3.14159265359,
	MapStrSlice: map[string][]int8{
		"": nil, // We must give a nil, not an allocated empty, slice to make deep equal comparison work.
		"uiouoiyioutuuiouoiyioutuuiouoiyioutu": {0x32, 0x12, 0x78, 0x00, 0x76},
		"RTQYWYEUIIUQIE": {0x32, 0x12, 0x78, 0x00, 0x76, 0x32, 0x12, 0x78, 0x00, 0x76, 0x32, 0x12, 0x78, 0x76,
			0x32, 0x12, 0x78, 0x00, 0x76, 0x32, 0x12, 0x78, 0x00, 0x76, 0x32, 0x12, 0x78, 0x76, 0x32, 0x12, 0x78, 0x00, 0x76, 0x32,
			0x12, 0x78, 0x00, 0x76, 0x32, 0x12, 0x78, 0x76, 0x32, 0x12, 0x78, 0x00, 0x76, 0x32, 0x12, 0x78, 0x00, 0x76, 0x32, 0x12},
	},
	InnerStruct2: InnerStruct2B{
		SliceUint64: []uint64{0, 1, 3, 100, 18446744073709551615, 1844674407370955, 4294967295},
		MapStrIntf: map[string]interface{}{
			"asdfasdf":     1,
			"asd":          "string",
			"oeiuwroewiro": float64(98.676),
			"kl":           float32(-234.0),
			"lkasdjf":      uint32(4294967290),
			"ss":           []string{"h", "ksjldf"},
		},
		Str2: `hello
there, this is a funky string. Tab: 	.`,
	},
}
