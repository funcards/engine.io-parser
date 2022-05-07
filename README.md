# engine.io-parser
![workflow](https://github.com/funcards/engine.io-parser/actions/workflows/workflow.yml/badge.svg)
[![Coverage Status](https://coveralls.io/repos/funcards/engine.io-parser/badge.svg?branch=main&service=github)](https://coveralls.io/github/funcards/engine.io-parser?branch=main)
[![GoDoc](https://godoc.org/github.com/funcards/engine.io-parser?status.svg)](https://pkg.go.dev/github.com/funcards/engine.io-parser/v4)
![License](https://img.shields.io/dub/l/vibe-d.svg)

This is the GO parser **version 4** for the engine.io protocol encoding.

## How to use

The parser can encode/decode packets, payloads and payloads as binary.

Example:

```go
var buf bytes.Buffer

payload := eio_parser.Payload{
    Packets: []eio_parser.Packet{
        {Type: eio_parser.PacketTypeOpen},
        {Type: eio_parser.PacketTypeClose},
        {Type: eio_parser.PacketTypePing, Data: "probe"},
        {Type: eio_parser.PacketTypePong, Data: "probe"},
        {Type: eio_parser.PacketTypeMessage, Data: "test"},
    },
}
payload.Encode(&buf)
fmt.Println(buf.Bytes())

payload1 = eio_parser.Payload{}
payload1.Decode(&buf) // payload == payload1
```

## License

Distributed under MIT License, please see license file within the code for more details.
