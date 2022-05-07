package eio_parser

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io"
)

type Encoder struct {
	w io.Writer
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w}
}

func (e *Encoder) Encode(pkt Packet) error {
	if data, ok := pkt.Data.([]byte); ok {
		// only 'message' packets can contain binary, so the type prefix is not needed
		return e.err(e.write(data, pkt.MsgType))
	}

	if err := e.write(pkt.Type.Bytes(), MessageTypeBinary); err != nil {
		return e.err(err)
	}

	if data, ok := pkt.Data.(string); ok {
		return e.err(e.write([]byte(data), MessageTypeBinary))
	}

	return nil
}

func (e *Encoder) err(err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("encode packet error: %w", err)
}

func (e *Encoder) write(data []byte, msgType MessageType) error {
	if msgType == MessageTypeBinary {
		_, err := e.w.Write(data)
		return err
	}

	_, err := e.w.Write([]byte("b"))
	if err != nil {
		return err
	}

	enc := base64.NewEncoder(base64.StdEncoding, e.w)
	_, err = enc.Write(data)
	if err != nil {
		return err
	}
	return enc.Close()
}

type Decoder struct {
	r io.Reader
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r}
}

func (d *Decoder) Decode(pkt *Packet) error {
	if pkt == nil {
		return d.err(ErrNilPacket)
	}

	if pkt.MsgType == MessageTypeBinary {
		data, err := io.ReadAll(d.r)
		if err != nil {
			return d.err(err)
		}

		pkt.Type = Message
		if len(data) > 0 {
			pkt.Data = data
		}
		return nil
	}

	r := bufio.NewReader(d.r)
	char, _, err := r.ReadRune()
	if err != nil {
		return d.err(err)
	}

	if char == 'b' {
		d.r = base64.NewDecoder(base64.StdEncoding, r)
		pkt.Type = Message
		return d.readAsStr(pkt)
	}

	pktType, err := NewPacketType(string(char))
	if err != nil {
		return d.err(err)
	}

	d.r = r
	pkt.Type = pktType
	return d.readAsStr(pkt)
}

func (d *Decoder) err(err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("decode packet error: %w", err)
}

func (d *Decoder) readAsStr(pkt *Packet) error {
	data, err := io.ReadAll(d.r)
	if err != nil {
		return d.err(err)
	}
	pkt.MsgType = MessageTypeString
	if len(data) > 0 {
		pkt.Data = string(data)
	}
	return nil
}
