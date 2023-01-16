# Serialization Formats Benchmarks

This repository contains benchmarking code for different serialization formats across programing languages.

Sample data came from [Farcaster](https://github.com/farcasterxyz/hub).

## Tested

### Serialization Formats

  * [Protocol Buffers](https://protobuf.dev)
  * [FlatBuffers](https://google.github.io/flatbuffers/)

### Programming Languages

  * Go

## Results

### Go (16 Jan 2023)

  * Benchmarks are CPU bound (no disk operations).
  * OS: macOS 13.0.1
  * CPU: Intel(R) Core(TM) i7-8559U CPU @ 2.70GHz

| Test                              | Protobuf | FlatBuffers |
| --------------------------------- | -------- | ----------- |
| Encode (ns/op)                    | 670.6    | 722.7       |
| Decode (ns/op)                    | 890.2    | 19.43       |
| Wire format size (bytes)          | 268      | 372         |
| Wire format size, gzipped (bytes) | 179      | 219         |

(Smaller is better)