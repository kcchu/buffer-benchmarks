.PHONY: generate
generate:
	flatc --rust -o src/flatbuffers ../schemas/message.fbs

.PHONY: bench
bench:
	cargo bench
