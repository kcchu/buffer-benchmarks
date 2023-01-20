use criterion::{criterion_group, criterion_main, Criterion};

use buffer_benchmarks::{flatbuffers, prost, protobuf};

fn bench_protobuf(c: &mut Criterion) {
    c.bench_function("encode_protobuf", |b| b.iter(|| protobuf::encode()));
    let buf = protobuf::encode();
    println!("Wire format size (bytes) = {}", buf.len());
    c.bench_function("decode_protobuf", |b| b.iter(|| protobuf::decode(&buf)));
}

fn bench_prost(c: &mut Criterion) {
    c.bench_function("encode_prost", |b| b.iter(|| prost::encode()));
    let buf = prost::encode();
    println!("Wire format size (bytes) = {}", buf.len());
    c.bench_function("decode_prost", |b| b.iter(|| prost::decode(&buf)));
}

fn bench_flatbuffers(c: &mut Criterion) {
    c.bench_function("encode_flatbuffers", |b| b.iter(|| flatbuffers::encode()));
    let buf = flatbuffers::encode();
    println!("Wire format size (bytes) = {}", buf.len());
    c.bench_function("decode_flatbuffers", |b| b.iter(|| flatbuffers::decode(&buf)));
}

criterion_group!(benches, bench_protobuf, bench_prost, bench_flatbuffers);
criterion_main!(benches);
