package eio_parser

import (
	"encoding/base64"
	"strings"
)

// EncodePacket encode packet
// supportsBinary whether the transport supports binary encoding
func EncodePacket(packet Packet, supportsBinary bool) any {
	if data, ok := packet.Data.([]byte); ok {
		return encodeByteArray(data, supportsBinary)
	}

	encoded := packet.Type.Encode()
	if packet.Data != nil {
		encoded += packet.Data.(string)
	}
	return encoded
}

func encodeByteArray(data []byte, supportsBinary bool) any {
	if supportsBinary {
		return data
	}
	return "b" + base64.StdEncoding.EncodeToString(data)
}

func DecodePacket(payload any) (Packet, error) {
	if payload == nil {
		return ErrorPacket(ErrInvalidPayload), ErrInvalidPayload
	}

	switch raw := payload.(type) {
	case string:
		if len(raw) == 0 {
			return ErrorPacket(ErrPayloadEmpty), ErrPayloadEmpty
		}
		// 98 = 'b'
		if 98 == raw[0] {
			data, err := base64.StdEncoding.DecodeString(raw[1:])
			if err != nil {
				return ErrorPacket(ErrDecodeBase64), err
			}
			return MessagePacket(data), nil
		}
		t, err := ParseTypeASCII(raw[0])
		if err != nil {
			return ErrorPacket(err), err
		}
		return TextPacket(t, raw[1:]), nil
	case []byte:
		return MessagePacket(raw), nil
	}

	return ErrorPacket(ErrInvalidType), ErrInvalidType
}

func EncodePayload(payload Payload) string {
	data := make([]string, 0, len(payload))
	for _, pkt := range payload {
		data = append(data, EncodePacket(pkt, false).(string))
	}
	return strings.Join(data, Separator)
}

func DecodePayload(payload any) (Payload, error) {
	str, ok := payload.(string)
	if !ok {
		return nil, ErrInvalidPayload
	}
	data := strings.Split(str, Separator)
	packets := make(Payload, 0, len(data))
	for _, item := range data {
		pkt, err := DecodePacket(item)
		if err != nil {
			return nil, err
		}
		packets = append(packets, pkt)
	}
	return packets, nil
}
