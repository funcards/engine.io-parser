package eio_parser

import (
	"bytes"
	"reflect"
	"testing"
)

func TestPayload_Encode(t *testing.T) {
	payload := Payload{
		Packets: []Packet{
			{Type: Open},
			{Type: Close},
			{Type: Ping, Data: "probe"},
			{Type: Pong, Data: "probe"},
			{Type: Message, Data: "test"},
		},
	}

	got := bytes.NewBuffer(make([]byte, 0, 16))
	if err := payload.Encode(got); err != nil {
		t.Errorf("Expected encode, but got: '%v'", err)
	}

	want := []byte{48, 30, 49, 30, 50, 112, 114, 111, 98, 101, 30, 51, 112, 114, 111, 98, 101, 30, 52, 116, 101, 115, 116}
	if bytes.Compare(got.Bytes(), want) != 0 {
		t.Errorf("Expected '%s', but got '%s'", want, got)
	}
}

func TestPayload_Decode(t *testing.T) {
	r := bytes.NewBuffer([]byte{48, 30, 49, 30, 50, 112, 114, 111, 98, 101, 30, 51, 112, 114, 111, 98, 101, 30, 52, 116, 101, 115, 116})

	var got Payload
	if err := got.Decode(r); err != nil {
		t.Errorf("Expected decode, but got: '%v'", err)
	}

	want := Payload{
		Packets: []Packet{
			{Type: Open},
			{Type: Close},
			{Type: Ping, Data: "probe"},
			{Type: Pong, Data: "probe"},
			{Type: Message, Data: "test"},
		},
	}

	if !reflect.DeepEqual(got.Packets, want.Packets) {
		t.Errorf("Expected '%v', but got '%v'", want.Packets, got.Packets)
	}
}
