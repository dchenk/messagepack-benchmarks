package mb

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/dchenk/msgp/msgp"
)

func TestDchenkMsgp1(t *testing.T) {

	buf := new(bytes.Buffer)

	err := msgp.Encode(buf, &Struct1A_inst)
	if err != nil {
		t.Fatalf("encoding: %v", err)
	}

	var dec Struct1A
	err = msgp.Decode(buf, &dec)
	if err != nil {
		t.Fatalf("decoding: %v", err)
	}

	if !reflect.DeepEqual(Struct1A_inst, dec) {
		t.Errorf("unequal;\ngot: %+v\nwanted: %+v", dec, Struct1A_inst)
	}

}

// Must check manually because of dynamic interface{} types inside.
func TestDchenkMsgp2(t *testing.T) {

	buf := new(bytes.Buffer)

	err := msgp.Encode(buf, &Struct2A_inst)
	if err != nil {
		t.Fatalf("encoding: %v", err)
	}

	var dec Struct2A
	err = msgp.Decode(buf, &dec)
	if err != nil {
		t.Fatalf("decoding: %v", err)
	}

	if dec.Str != Struct2A_inst.Str {
		t.Errorf("unequal Str; got: %v", dec.Str)
	}

	if !bytes.Equal(dec.Slice, Struct2A_inst.Slice) {
		t.Errorf("unequal Slice; got: %v", dec.Slice)
	}

	if !reflect.DeepEqual(dec.MapStrUint64, Struct2A_inst.MapStrUint64) {
		t.Errorf("unequal MapStrUint64; got: %v", dec.MapStrUint64)
	}

	if !reflect.DeepEqual(dec.InnerStruct, Struct2A_inst.InnerStruct) {
		t.Errorf("unequal InnerStruct; got: %+v", dec.InnerStruct)
	}

	if dec.Float64 != Struct2A_inst.Float64 {
		t.Errorf("unequal Float64; got: %v", dec.Float64)
	}

	if !reflect.DeepEqual(dec.MapStrSlice, Struct2A_inst.MapStrSlice) {
		mpOriginal := Struct2A_inst.MapStrSlice
		if lDec := len(dec.MapStrSlice); lDec != len(mpOriginal) {
			t.Errorf("unequal legths for MapStrSlice; got len %d", lDec)
		}
		for k, v := range dec.MapStrSlice {
			if !reflect.DeepEqual(v, mpOriginal[k]) {
				t.Errorf("unequal MapStrSlice values for key %q; got: %v", k, v)
			}
		}
	}

	if !reflect.DeepEqual(dec.InnerStruct2, Struct2A_inst.InnerStruct2) {

		orig := Struct2A_inst.InnerStruct2

		if !reflect.DeepEqual(dec.InnerStruct2.SliceUint64, orig.SliceUint64) {
			t.Errorf("unequal MapStrSlice.SliceUint64; got %v", dec.InnerStruct2.SliceUint64)
		}

		mpDec, mpOriginal := dec.InnerStruct2.MapStrIntf, orig.MapStrIntf
		if lDec := len(mpDec); lDec != len(mpOriginal) {
			t.Fatalf("unequal legths for InnerStruct2.MapStrIntf; got len %d", lDec)
		}
		for k, v := range mpDec {
			if !reflect.DeepEqual(v, mpOriginal[k]) {
				// Some of the elements in the map (with type interface{}) have a dynamic type that is not
				// decoded into the same exact type that they were decoded from. For these, listed in the
				// switch cases, it's acceptable to check if they are basically the same type and value.
				// Check out msgp.ReadIntfBytes and msgp.*Reader.ReadIntf.
				switch k {
				case "lkasdjf":
					valInt, ok := v.(uint64)
					if ok {
						val := uint32(valInt)
						if mpOriginal[k] != val {
							t.Errorf("wrong value; got value %v", v)
						}
					} else {
						t.Logf("unexpected type for InnerStruct2.MapStrIntf[%v], got type %T", k, v)
					}
				case "ss":
					origSlice, ok := mpOriginal[k].([]string)
					if !ok {
						t.Fatalf("original InnerStruct2.MapStrIntf[%v] must be a []string{}; type is %T", k, mpOriginal[k])
					}
					valSlice, ok := v.([]interface{})
					if ok {
						for i, val := range valSlice {
							strVal, ok := val.(string)
							if ok {
								if strVal != origSlice[i] {
									t.Errorf("(index %d) bad value %q", i, strVal)
								}
							}
						}
					} else {
						t.Errorf("unexpected type for InnerStruct2.MapStrIntf[%v], got type %T", k, v)
					}
				case "asdfasdf":
					valInt, ok := v.(int64)
					if ok {
						val := int(valInt)
						if mpOriginal[k] != val {
							t.Errorf("wrong value; got value %v", v)
						}
					} else {
						t.Logf("unexpected type for InnerStruct2.MapStrIntf[%v], got type %T", k, v)
					}
				default:
					t.Errorf("unequal InnerStruct2.MapStrIntf values for key %q; got: %v", k, v)
				}
			}
		}

	}

}
