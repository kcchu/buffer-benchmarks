.PHONY: generate
generate:
	flatc --rust -o src/flatbuffers ../schemas/message.fbs

.PHONY: bench
bench:
	rustup default nightly
	cargo bench -- --nocapture
