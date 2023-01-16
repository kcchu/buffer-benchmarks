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

### Go (16 Jan 2023)

  * Benchmarks are CPU bound (no disk operations)
  * OS: macOS 13.0.1
  * CPU: Intel(R) Core(TM) i7-8559U CPU @ 2.70GHz
  * Cap'n Proto: [go-capnproto2](https://github.com/capnproto/go-capnproto2) (using packed encoding)

| Test                              | Protobuf | FlatBuffers | Capnp |
| --------------------------------- | -------- | ----------- | ----- |
| Encode (ns/op)                    | 882.7    | 859.0       | 2541  |
| Decode (ns/op)                    | 1165     | 18.36       | 1726  |
| Wire format size (bytes)          | 299      | 432         | 344   |
| Wire format size, gzipped (bytes) | 323      | 406         | 368   |

> **Note**
> Smaller is better