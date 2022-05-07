package eio_parser

import (
	"bytes"
	"fmt"
	"testing"
)

func TestPktDecoder_Decode(t *testing.T) {
	cases := []string{"", "a123"}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("should fail decode \"%s\"", tc), func(t *testing.T) {
			dec := NewPktDecoder(bytes.NewReader([]byte(tc)))
			got := dec.Decode(new(Packet))
			if got == nil {
				t.Fail()
			}
		})
	}

	wantType := PacketTypeMessage
	wantData := "test"
	arg := wantType.Str() + wantData
	t.Run(fmt.Sprintf("%s=%s%s", arg, wantType.Str(), wantData), func(t *testing.T) {
		var pkt Packet
		dec := NewPktDecoder(bytes.NewReader([]byte(arg)))
		if err := dec.Decode(&pkt); err != nil {
			t.Errorf("Expected decode, but got: '%v'", err)
		}
		if pkt.Type != wantType || pkt.Data.(string) != wantData {
			t.Errorf("Expected '%s%s', but got '%s%s'", wantType.String(), wantData, pkt.MsgType.String(), pkt.Data)
		}
	})
}

func TestPktEncoder_Encode(t *testing.T) {
	type testCase struct {
		arg  Packet
		want string
	}
	cases := []testCase{
		{Packet{MsgType: MessageTypeBinary, Type: PacketTypeMessage, Data: "test"}, "4test"},
		{Packet{MsgType: MessageTypeBinary, Type: PacketTypeMessage, Data: []byte("test")}, "test"},
	}
	for _, tc := range cases {
		t.Run(tc.want, func(t *testing.T) {
			var buf bytes.Buffer
			if err := NewPktEncoder(&buf).Encode(tc.arg); err != nil {
				t.Errorf("Expected encode, but got: '%v'", err)
			}
			got := buf.String()
			if tc.want != got {
				t.Errorf("Expected '%s', but got '%s'", tc.want, got)
			}
		})
	}
}
