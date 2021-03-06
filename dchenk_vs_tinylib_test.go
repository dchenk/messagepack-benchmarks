package mb

import (
	"bytes"
	"testing"

	dc "github.com/dchenk/msgp/msgp"
	ti "github.com/tinylib/msgp/msgp"
)

func TestDchenkMsgp_TinylibMsgp_1(t *testing.T) {

	bufD := new(bytes.Buffer)
	err := dc.Encode(bufD, &Struct1A_inst)
	if err != nil {
		t.Fatalf("could not encode dchenk; %v", err)
	}

	bufT := new(bytes.Buffer)
	err = ti.Encode(bufT, &Struct1B_inst)
	if err != nil {
		t.Fatalf("could not encode tinylib; %v", err)
	}

	// Because we have (unordered) maps, the serializations will not always be equal, so we
	// only compare lengths.
	if bufD.Len() != bufT.Len() {
		t.Errorf("lengths are not equal")
	}

}

func TestDchenkMsgp_TinylibMsgp_2(t *testing.T) {

	bufD := new(bytes.Buffer)
	err := dc.Encode(bufD, &Struct2A_inst)
	if err != nil {
		t.Fatalf("could not encode dchenk; %v", err)
	}

	bufT := new(bytes.Buffer)
	err = ti.Encode(bufT, &Struct2B_inst)
	if err != nil {
		t.Fatalf("could not encode tinylib; %v", err)
	}

	// Because we have (unordered) maps, the serializations will not always be equal, so we
	// only compare lengths.
	if bufD.Len() != bufT.Len() {
		t.Errorf("lengths are not equal")
	}

}

func BenchmarkCustomDchenk1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Result, _ = Struct1A_inst.MarshalMsg(nil)
	}
}

func BenchmarkCustomTinylib1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Result, _ = Struct1B_inst.MarshalMsg(nil)
	}
}

func BenchmarkCustomDchenk2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Result, _ = Struct2A_inst.MarshalMsg(nil)
	}
}

func BenchmarkCustomTinylib2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Result, _ = Struct2B_inst.MarshalMsg(nil)
	}
}

var Result []byte

func BenchmarkCustomUnmarshalDchenk1(b *testing.B) {
	msg, err := Struct1A_inst.MarshalMsg(nil)
	if err != nil {
		b.Fatalf("error marshalling; %v", err)
	}
	for i := 0; i < b.N; i++ {
		Result, _ = Struct1A_inst.UnmarshalMsg(msg)
	}
}

func BenchmarkCustomUnmarshalTinylib1(b *testing.B) {
	msg, err := Struct1B_inst.MarshalMsg(nil)
	if err != nil {
		b.Fatalf("error marshalling; %v", err)
	}
	for i := 0; i < b.N; i++ {
		Result, _ = Struct1B_inst.UnmarshalMsg(msg)
	}
}
