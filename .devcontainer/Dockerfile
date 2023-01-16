FROM mcr.microsoft.com/devcontainers/go:1-bullseye

RUN apt-get update && apt-get install -y \
    cmake \
    protobuf-compiler \
    && rm -rf /var/lib/apt/lists/*
RUN cd /tmp \
    && git clone https://github.com/google/flatbuffers.git \
    && cd flatbuffers \
    && cmake -G "Unix Makefiles" -DCMAKE_BUILD_TYPE=Release \
    && make install \
    && rm -fr /tmp/flatbuffers
RUN su vscode -c "go install google.golang.org/protobuf/cmd/protoc-gen-go@latest"