# engine.io-parser

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
