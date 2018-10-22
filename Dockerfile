FROM golang:1.11.1

RUN apt-get update -y && \
    apt-get upgrade -y && \
    apt-get install -y bsdtar bash-completion jq uuid-runtime ca-certificates && \
    cd /usr/local/ && \
    curl -s -L "https://github.com/google/protobuf/releases/download/v3.6.1/protoc-3.6.1-linux-x86_64.zip" | bsdtar -xf- bin/protoc && \
	chmod u+x bin/protoc && \
    cd /usr/local/src/ && \
    curl -s -L "https://github.com/google/protobuf/archive/v3.6.1.tar.gz" | tar xvz && \
	mv protobuf-3.6.1/ /usr/local/src/protobuf/ && \
	curl -s -L "https://github.com/grpc-ecosystem/grpc-gateway/archive/v1.5.1.tar.gz" | tar xvz && \
	mv grpc-gateway-1.5.1/ /usr/local/src/grpc-gateway/ && \
    # Makefile completion
    apt-get install -y bash-completion && \
    echo ". /etc/bash_completion" >> /root/.bashrc

WORKDIR /code

ENV GO111MODULE=on
ENV GOBIN=/code/bin/
ENV PATH=/code/bin/:$PATH
