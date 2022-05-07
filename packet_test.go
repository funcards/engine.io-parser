package eio_parser

import (
	"fmt"
	"testing"
)

func TestNewPacketType(t *testing.T) {
	type testCase struct {
		arg  string
		want PacketType
	}

	cases := []testCase{
		{"message", PacketTypeMessage},
		{"noop", PacketTypeNoop},
		{"4", PacketTypeMessage},
		{"2", PacketTypePing},
		{"", PacketTypeError},
		{"fail", PacketTypeError},
		{"-2", PacketTypeError},
		{"7", PacketTypeError},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("%s=%d", tc.arg, tc.want), func(t *testing.T) {
			got, _ := NewPacketType(tc.arg)
			if tc.want != got {
				t.Errorf("Expected '%d', but got '%d'", tc.want, got)
			}
		})
	}
}

func TestPacketType_String(t *testing.T) {
	type testCase struct {
		arg  PacketType
		want string
	}
	cases := []testCase{
		{PacketTypeUpgrade, StrPacketTypeUpgrade},
		{PacketTypeOpen, StrPacketTypeOpen},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("%d=%s", tc.arg, tc.want), func(t *testing.T) {
			got := tc.arg.String()
			if tc.want != got {
				t.Errorf("Expected '%s', but got '%s'", tc.want, got)
			}
		})
	}
}
