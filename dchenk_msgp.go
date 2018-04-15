// This is how you would do it using github.com/dchenk/msgp
package mb

//go:generate msgp_dchenk

type Struct1A struct {
	Str          string            `msgp:"str"`
	Slice        []byte            `msgp:"slice"`
	MapStrUint16 map[string]uint16 `msgp:"map_str_uint16"`
	InnerStruct  InnerStructA      `msgp:"inner_struct"`
}

type InnerStructA struct {
	Bool    bool    `msgp:"bool"`
	Float32 float32 `msgp:"float_32"`
}

type Struct2A struct {
	Str          string            `msgp:"str"`
	Slice        []byte            `msgp:"slice"`
	MapStrUint64 map[string]uint64 `msgp:"map_str_uint64"`
	InnerStruct  InnerStructA      `msgp:"inner_struct"`
	Float64      float64           `msgp:"float_64"`
	MapStrSlice  map[string][]int8 `msgp:"map_str_slice"`
	InnerStruct2 InnerStruct2A     `msgp:"inner_struct_2"`
}

type InnerStruct2A struct {
	SliceUint64 []uint64               `msgp:"slice_uint64"`
	MapStrIntf  map[string]interface{} `msgp:"map_str_intf"`
	Str2        string                 `msgp:"string"`
}

var Struct1A_inst = Struct1A{
	Str:   "hello",
	Slice: []byte{'a', 'b', 'c'},
	MapStrUint16: map[string]uint16{
		"abc": 23,
		"lasdhflahsdlfkhlasdf": 4,
		"c": 1356,
	},
	InnerStruct: InnerStructA{true, 213.513},
}

var Struct2A_inst = Struct2A{
	Str:   "lk\n\nsdj flasjfdlkaj\nsdlfka\tjs\nldfk\"jal kas_ ad\"fsfka\tjs\nldfk\"jal kas_ ad\"fska\tjs\nldfk\"jal kas_ adk",
	Slice: []byte{0x01, 'a', 'b', 'c', 'Z', 0x03, 0x05},
	MapStrUint64: map[string]uint64{
		"abc":          23532525,
		"defghijkl":    1,
		"ghijkl":       98,
		"ghijklghijkl": 8000111333,
	},
	InnerStruct: InnerStructA{false, -132.88889},
	Float64:     3.14159265359,
	MapStrSlice: map[string][]int8{
		"": {},
		"uiouoiyioutuuiouoiyioutuuiouoiyioutu": {0x32, 0x12, 0x78, 0x00, 0x76},
		"RTQYWYEUIIUQIE": {0x32, 0x12, 0x78, 0x00, 0x76, 0x32, 0x12, 0x78, 0x00, 0x76, 0x32, 0x12, 0x78, 0x76,
			0x32, 0x12, 0x78, 0x00, 0x76, 0x32, 0x12, 0x78, 0x00, 0x76, 0x32, 0x12, 0x78, 0x76, 0x32, 0x12, 0x78, 0x00, 0x76, 0x32,
			0x12, 0x78, 0x00, 0x76, 0x32, 0x12, 0x78, 0x76, 0x32, 0x12, 0x78, 0x00, 0x76, 0x32, 0x12, 0x78, 0x00, 0x76, 0x32, 0x12},
	},
	InnerStruct2: InnerStruct2A{
		SliceUint64: []uint64{0, 1, 3, 100, 18446744073709551615, 1844674407370955, 4294967295},
		MapStrIntf: map[string]interface{}{
			"asdfasdf":     1,
			"asd":          "string",
			"oeiuwroewiro": float64(98.676),
			"kl":           float32(-234.0),
			"lkasdjf":      uint32(4294967290),
		},
		Str2: `hello
there, this is a funky string. Tab: 	.`,
	},
}
