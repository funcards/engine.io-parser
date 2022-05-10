package eio_parser

import (
	"fmt"
	"testing"
)

func TestParseType(t *testing.T) {
	type testCase struct {
		arg  string
		want PacketType
	}

	cases := []testCase{
		{"message", Message},
		{"noop", Noop},
		{"4", Message},
		{"2", Ping},
		{"", Error},
		{"fail", Error},
		{"-2", Error},
		{"7", Error},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("%s=%d", tc.arg, tc.want), func(t *testing.T) {
			got, _ := ParseType(tc.arg)
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
		{Upgrade, StrUpgrade},
		{Open, StrOpen},
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
