package eio_parser

import (
	"errors"
	"fmt"
	"io"
	"strconv"
)

var (
	ErrNilPacket         = errors.New("packet is nil")
	ErrInvalidPacketType = errors.New("invalid packet type")
)

type (
	// PacketType indicates type of engine.io Packet
	PacketType byte
	// MessageType indicates type of engine.io Message
	MessageType byte

	Packet struct {
		MsgType MessageType `json:"-"`
		Type    PacketType  `json:"type"`
		Data    any         `json:"data,omitempty"`
	}
)

const (
	// PacketTypeOpen is sent from the server when new transport is opened (recheck)
	PacketTypeOpen PacketType = iota
	// PacketTypeClose requests the close of this transport but does not shut down the connection itself.
	PacketTypeClose
	// PacketTypePing is sent by the client. Server should answer with a pong packet containing the same data
	PacketTypePing
	// PacketTypePong is sent by the server to respond to ping packets.
	PacketTypePong
	// PacketTypeMessage denotes actual message, client and server should call their callbacks with the data.
	PacketTypeMessage
	// PacketTypeUpgrade is sent by the client requesting the server to flush its cache on the old transport and switch to the new transport.
	PacketTypeUpgrade
	// PacketTypeNoop denotes a noop packet. Used primarily to force a poll cycle when an incoming websocket connection is received.
	PacketTypeNoop
	// PacketTypeError is for internal use
	PacketTypeError
)

const (
	StrPacketTypeOpen    = "open"
	StrPacketTypeClose   = "close"
	StrPacketTypePing    = "ping"
	StrPacketTypePong    = "pong"
	StrPacketTypeMessage = "message"
	StrPacketTypeUpgrade = "upgrade"
	StrPacketTypeNoop    = "noop"
	StrPacketTypeError   = "error"
)

const (
	// MessageTypeString indicates Message encoded as string
	MessageTypeString MessageType = iota
	// MessageTypeBinary indicates Message encoded as binary
	MessageTypeBinary
)

var (
	mapTypeToStr = map[PacketType]string{
		PacketTypeOpen:    StrPacketTypeOpen,
		PacketTypeClose:   StrPacketTypeClose,
		PacketTypePing:    StrPacketTypePing,
		PacketTypePong:    StrPacketTypePong,
		PacketTypeMessage: StrPacketTypeMessage,
		PacketTypeUpgrade: StrPacketTypeUpgrade,
		PacketTypeNoop:    StrPacketTypeNoop,
	}
	mapStrToType = map[string]PacketType{
		StrPacketTypeOpen:    PacketTypeOpen,
		StrPacketTypeClose:   PacketTypeClose,
		StrPacketTypePing:    PacketTypePing,
		StrPacketTypePong:    PacketTypePong,
		StrPacketTypeMessage: PacketTypeMessage,
		StrPacketTypeUpgrade: PacketTypeUpgrade,
		StrPacketTypeNoop:    PacketTypeNoop,
	}
)

func NewPacketType(str string) (PacketType, error) {
	if len(str) > 1 {
		if p, ok := mapStrToType[str]; ok {
			return p, nil
		}
		return PacketTypeError, fmt.Errorf("%s error: %w", str, ErrInvalidPacketType)
	}

	n, err := strconv.Atoi(str)
	if err != nil || n < PacketTypeOpen.Int() || n > PacketTypeNoop.Int() {
		return PacketTypeError, fmt.Errorf("%s error: %w", str, ErrInvalidPacketType)
	}

	return PacketType(n), nil
}

// String returns string representation of a PacketType
func (p PacketType) String() string {
	if str, ok := mapTypeToStr[p]; ok {
		return str
	}
	return StrPacketTypeError
}

func (p PacketType) Int() int {
	return int(p)
}

func (p PacketType) Str() string {
	return strconv.Itoa(p.Int())
}

func (p PacketType) Bytes() []byte {
	return []byte(p.Str())
}

// String returns string representation of a MessageType
func (m MessageType) String() string {
	switch m {
	case MessageTypeString:
		return "string"
	case MessageTypeBinary:
		return "binary"
	}
	return "invalid"
}

func (p *Packet) Encode(w io.Writer) error {
	return NewPktEncoder(w).Encode(*p)
}

func (p *Packet) Decode(r io.Reader) error {
	return NewPktDecoder(r).Decode(p)
}
