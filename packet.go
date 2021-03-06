package eiop

import "strconv"

var NoopPacket = Packet{Type: Noop}

type (
	// PacketType indicates type of engine.io Packet
	PacketType byte
	Payload    []Packet

	Packet struct {
		Type PacketType
		Data any
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

func (p PacketType) Encode() string {
	return strconv.Itoa(p.Int())
}

func (p PacketType) Bytes() []byte {
	return []byte(p.Encode())
}

func TextPacket(t PacketType, data ...string) Packet {
	return Packet{Type: t, Data: optstr(data...)}
}

func BinaryPacket(t PacketType, data []byte) Packet {
	return Packet{Type: t, Data: data}
}

func MessagePacket(data any) Packet {
	return Packet{Type: Message, Data: data}
}

func OpenPacket(data ...string) Packet {
	return TextPacket(Open, data...)
}

func ClosePacket(reason ...string) Packet {
	return TextPacket(Close, reason...)
}

func PingPacket(data ...string) Packet {
	return TextPacket(Ping, data...)
}

func PongPacket(data ...string) Packet {
	return TextPacket(Pong, data...)
}

func ErrorPacket(err error) Packet {
	return Packet{Type: Error, Data: err}
}

func optstr(str ...string) any {
	if len(str) == 0 || len(str[0]) == 0 {
		return nil
	}
	return str[0]
}

func (p Payload) Encode() any {
	return EncodePayload(p)
}

func (p Packet) Encode(supportsBinary bool) any {
	return EncodePacket(p, supportsBinary)
}
