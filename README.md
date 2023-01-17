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
  * Protocol Buffer: [golang/protobuf](https://github.com/golang/protobuf)
  * FlatBuffers: [google/flatbuffers/go](https://github.com/google/flatbuffers/tree/master/go) v0.0.0-20230110200425-62e4d2e5b215
  * Cap'n Proto: [capnproto/go-capnproto2](https://github.com/capnproto/go-capnproto2) v3.0.0-alpha.23

| Test                              | Protobuf | FlatBuffers | Capnp | Capnp (packed) |
| --------------------------------- | -------- | ----------- | ----- | -------------- |
| Encode (ns/op)                    | 898.0    | 868.5       | 1740  | 2588           |
| Decode (ns/op)                    | 1205     | 18.56       | 837.3 | 1732           |
| Wire format size (bytes)          | 299      | 432         | 440   | 344            |
| Wire format size, gzipped (bytes) | 323      | 406         | 392   | 368            |

> **Note**
> Smaller is better