package eiop

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

func TestDecodePacket(t *testing.T) {
	cases := []string{"", "a123"}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("should fail decode \"%s\"", tc), func(t *testing.T) {
			_, got := DecodePacket(tc)
			if got == nil {
				t.Fail()
			}
		})
	}

	wantType := Message
	wantData := "test"
	arg := wantType.Encode() + wantData
	t.Run(fmt.Sprintf("%s=%s%s", arg, wantType.Encode(), wantData), func(t *testing.T) {
		pkt, err := DecodePacket(arg)
		if err != nil {
			t.Errorf("Expected decode, but got: '%v'", err)
		}
		if pkt.Type != wantType || pkt.Data.(string) != wantData {
			t.Errorf("Expected '%s%s', but got '%s%s'", wantType.String(), wantData, pkt.Type.String(), pkt.Data)
		}
	})
}

func TestEncodePacket(t *testing.T) {
	type testCase struct {
		arg  Packet
		want any
	}
	cases := []testCase{
		{Packet{Type: Message, Data: "test"}, "4test"},
		{Packet{Type: Message, Data: []byte("test")}, []byte("test")},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("%v", tc.want), func(t *testing.T) {
			got := EncodePacket(tc.arg, true)
			if !reflect.DeepEqual(tc.want, got) {
				t.Errorf("Expected '%v', but got '%v'", tc.want, got)
			}
		})
	}
}

func TestEncodePayload(t *testing.T) {
	payload := Payload{
		{Type: Open},
		{Type: Close},
		{Type: Ping, Data: "probe"},
		{Type: Pong, Data: "probe"},
		{Type: Message, Data: "test"},
	}

	got := EncodePayload(payload)
	want := []byte{48, 30, 49, 30, 50, 112, 114, 111, 98, 101, 30, 51, 112, 114, 111, 98, 101, 30, 52, 116, 101, 115, 116}
	if bytes.Compare(want, []byte(got)) != 0 {
		t.Errorf("Expected '%s', but got '%s'", want, got)
	}
}

func TestDecodePayload(t *testing.T) {
	data := string([]byte{48, 30, 49, 30, 50, 112, 114, 111, 98, 101, 30, 51, 112, 114, 111, 98, 101, 30, 52, 116, 101, 115, 116})

	got, err := DecodePayload(data)
	if err != nil {
		t.Errorf("Expected decode, but got: '%v'", err)
	}

	want := Payload{
		{Type: Open},
		{Type: Close},
		{Type: Ping, Data: "probe"},
		{Type: Pong, Data: "probe"},
		{Type: Message, Data: "test"},
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Expected '%v', but got '%v'", want, got)
	}
}
