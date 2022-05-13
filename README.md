# engine.io-parser

![workflow](https://github.com/funcards/engine.io-parser/actions/workflows/workflow.yml/badge.svg)
[![Coverage Status](https://coveralls.io/repos/funcards/engine.io-parser/badge.svg?branch=main&service=github)](https://coveralls.io/github/funcards/engine.io-parser?branch=main)
[![GoDoc](https://godoc.org/github.com/funcards/engine.io-parser?status.svg)](https://pkg.go.dev/github.com/funcards/engine.io-parser/v4)
![License](https://img.shields.io/dub/l/vibe-d.svg)

This is the GO parser version **4** for the engine.io protocol encoding.

## Installation

Use go get.

```bash
go get github.com/funcards/engine.io-parser/v4
```

Then import the parser package into your own code.

```go
import "github.com/funcards/engine.io-parser/v4"
```

## How to use

The parser can encode/decode packets, payloads and payloads as binary.

Example:

```go
payload := eiop.Payload{
    {Type: eiop.Open},
    {Type: eiop.Close},
    {Type: eiop.Ping, Data: "probe"},
    {Type: eiop.Pong, Data: "probe"},
    {Type: eiop.Message, Data: "test"},
}
encoded := payload.Encode(&buf)
fmt.Println(encoded.(string))

payload1 = eiop.DecodePayload(encoded) // payload == payload1
```

## License

Distributed under MIT License, please see license file within the code for more details.
