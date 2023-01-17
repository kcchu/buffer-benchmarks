# Serialization Formats Benchmarks

This repository contains benchmarking code for different serialization formats across programing languages.

Sample data came from [Farcaster](https://github.com/farcasterxyz/hub).

## Tested

### Serialization Formats

  * [Protocol Buffers](https://protobuf.dev)
  * [FlatBuffers](https://google.github.io/flatbuffers/)
  * [Cap'n Proto](https://capnproto.org/)

### Programming Languages

  * Go

## Results

### Go (17 Jan 2023)

  * Benchmarks are CPU bound (no disk operations)
  * OS: macOS 13.0.1
  * CPU: Intel(R) Core(TM) i7-8559U CPU @ 2.70GHz
  * Protobuf: [golang/protobuf](https://github.com/golang/protobuf)
  * Protobuf (gogofaster): [gogoprotobuf](https://github.com/gogo/protobuf)
  * FlatBuffers: [google/flatbuffers/go](https://github.com/google/flatbuffers/tree/master/go) v0.0.0-20230110200425-62e4d2e5b215
  * Capnp: [capnproto/go-capnproto2](https://github.com/capnproto/go-capnproto2) v3.0.0-alpha.23

| Test                              | Protobuf | Protobuf (gogofaster) | FlatBuffers | Capnp | Capnp (packed) |
| --------------------------------- | -------- | --------------------| ----------- | ----- | -------------- |
| Encode (ns/op)                    | 883.8    | 384.4               | 856.8       | 1709  | 2591           |
| Decode (ns/op)                    | 1179     | 496.2               | 18.89       | 830.8 | 1716           |
| Wire format size (bytes)          | 299      | 299                 | 432         | 440   | 344            |
| Wire format size, gzipped (bytes) | 323      | 323                 | 406         | 392   | 368            |

> **Note**
> Smaller is better