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
	// Open is sent from the server when new transport is opened (recheck)
	Open PacketType = iota
	// Close requests the close of this transport but does not shut down the connection itself.
	Close
	// Ping is sent by the client. Server should answer with a pong packet containing the same data
	Ping
	// Pong is sent by the server to respond to ping packets.
	Pong
	// Message denotes actual message, client and server should call their callbacks with the data.
	Message
	// Upgrade is sent by the client requesting the server to flush its cache on the old transport and switch to the new transport.
	Upgrade
	// Noop denotes a noop packet. Used primarily to force a poll cycle when an incoming websocket connection is received.
	Noop
	// Error is for internal use
	Error
)

const (
	StrOpen    = "open"
	StrClose   = "close"
	StrPing    = "ping"
	StrPong    = "pong"
	StrMessage = "message"
	StrUpgrade = "upgrade"
	StrNoop    = "noop"
	StrError   = "error"
)

const (
	// MessageTypeString indicates Message encoded as string
	MessageTypeString MessageType = iota
	// MessageTypeBinary indicates Message encoded as binary
	MessageTypeBinary
)

var (
	mapTypeToStr = map[PacketType]string{
		Open:    StrOpen,
		Close:   StrClose,
		Ping:    StrPing,
		Pong:    StrPong,
		Message: StrMessage,
		Upgrade: StrUpgrade,
		Noop:    StrNoop,
	}
	mapStrToType = map[string]PacketType{
		StrOpen:    Open,
		StrClose:   Close,
		StrPing:    Ping,
		StrPong:    Pong,
		StrMessage: Message,
		StrUpgrade: Upgrade,
		StrNoop:    Noop,
	}
)

func NewPacketType(str string) (PacketType, error) {
	if len(str) > 1 {
		if p, ok := mapStrToType[str]; ok {
			return p, nil
		}
		return Error, fmt.Errorf("%s error: %w", str, ErrInvalidPacketType)
	}

	n, err := strconv.Atoi(str)
	if err != nil || n < Open.Int() || n > Noop.Int() {
		return Error, fmt.Errorf("%s error: %w", str, ErrInvalidPacketType)
	}

	return PacketType(n), nil
}

// String returns string representation of a PacketType
func (p PacketType) String() string {
	if str, ok := mapTypeToStr[p]; ok {
		return str
	}
	return StrError
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
