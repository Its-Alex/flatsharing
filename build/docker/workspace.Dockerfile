FROM golang:1.13

RUN apt-get update -y && \
    apt-get upgrade -y && \
    apt-get install -y bsdtar bash-completion jq uuid-runtime ca-certificates && \
    cd /usr/local/ && \
    curl -s -L "https://github.com/protocolbuffers/protobuf/releases/download/v3.10.0/protoc-3.10.0-linux-x86_64.zip" | bsdtar -xf- bin/protoc && \
	chmod u+x bin/protoc && \
    cd /usr/local/src/ && \
    curl -s -L "https://github.com/protocolbuffers/protobuf/archive/v3.10.0.tar.gz" | tar xvz && \
	mv protobuf-3.10.0/ /usr/local/src/protobuf/ && \
	curl -s -L "https://github.com/grpc-ecosystem/grpc-gateway/archive/v1.11.3.tar.gz" | tar xvz && \
	mv grpc-gateway-1.11.3/ /usr/local/src/grpc-gateway/ && \
    # Makefile completion
    apt-get install -y bash-completion && \
    echo ". /etc/bash_completion" >> /root/.bashrc && \
    # Golint
    go get -u golang.org/x/lint/golint && \
    go install golang.org/x/lint/golint

WORKDIR /code

ENV GO111MODULE=on
ENV GOCACHE="/go/cache"
ENV GOBIN="/go/bin"
ENV PATH="/go/bin:$PATH"
