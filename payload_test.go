package eio_parser

import (
	"bytes"
	"reflect"
	"testing"
)

func TestPayload_Encode(t *testing.T) {
	payload := Payload{
		Packets: []Packet{
			{Type: PacketTypeOpen},
			{Type: PacketTypeClose},
			{Type: PacketTypePing, Data: "probe"},
			{Type: PacketTypePong, Data: "probe"},
			{Type: PacketTypeMessage, Data: "test"},
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
			{Type: PacketTypeOpen},
			{Type: PacketTypeClose},
			{Type: PacketTypePing, Data: "probe"},
			{Type: PacketTypePong, Data: "probe"},
			{Type: PacketTypeMessage, Data: "test"},
		},
	}

	if !reflect.DeepEqual(got.Packets, want.Packets) {
		t.Errorf("Expected '%v', but got '%v'", want.Packets, got.Packets)
	}
}

// [{string open <nil>} {string close <nil>} {string ping probe} {string pong probe} {string message test}]
// [{string open <nil>} {string close <nil>} {string ping probe} {string pong probe} {string message tes}]
