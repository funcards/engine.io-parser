package eio_parser

import (
	"bytes"
	"fmt"
	"io"
)

const (
	Protocol  = 4
	Sep       = '\u001E' // rune(30)
	Separator = string(Sep)
)

type Payload struct {
	Packets        []Packet `json:"packets"`
	SupportsBinary bool     `json:"-"`
}

func (p *Payload) Encode(w io.Writer) error {
	for i, pkt := range p.Packets {
		if err := pkt.Encode(w); err != nil {
			return p.eerr(err)
		}

		if (i + 1) == len(p.Packets) {
			return nil
		}

		if _, err := w.Write([]byte(Separator)); err != nil {
			return p.eerr(err)
		}
	}

	return nil
}

func (p *Payload) Decode(r io.Reader) error {
	p.Packets = make([]Packet, 0)
	pktBuf := bytes.NewBuffer([]byte{})
	chunk := make([]byte, 128)

	write := func(i int) error {
		_, err := pktBuf.Write(chunk[:i])
		if err != nil {
			return p.derr(err)
		}
		return nil
	}

	for {
		n, err := r.Read(chunk)
		if err != nil && err != io.EOF {
			return p.derr(err)
		}

	start:
		for i, j := 0, 1; i < n; i, j = i+1, j+1 {
			if chunk[i] != byte(Sep) && j != n {
				continue
			}

			if j == n {
				if err1 := write(j); err1 != nil {
					return err1
				}
			} else {
				if err1 := write(i); err1 != nil {
					return err1
				}
			}

			var pkt Packet
			err1 := pkt.Decode(pktBuf)
			if err1 != nil {
				return p.derr(err1)
			}

			p.Packets = append(p.Packets, pkt)
			pktBuf.Reset()
			chunk = chunk[i+1:]
			n -= j

			if j != n {
				goto start
			}
		}

		if err == io.EOF {
			break
		}
	}

	return nil
}

func (p *Payload) eerr(err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("encode payload error: %w", err)
}

func (p *Payload) derr(err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("decode payload error: %w", err)
}
