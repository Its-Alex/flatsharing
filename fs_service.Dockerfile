FROM golang:1.11-alpine3.8 as builder

ENV GO111MODULE=on

COPY src/flatsharing /flatsharing/src/flatsharing
COPY src/auth /flatsharing/src/auth
COPY src/core /flatsharing/src/core
COPY go.mod /flatsharing/go.mod
COPY go.sum /flatsharing/go.sum

WORKDIR /flatsharing

# Install gox
RUN apk --no-cache add git ca-certificates gcc musl-dev && \
    go mod download && \
    go get -v github.com/mitchellh/gox && \
    gox -output="fs_service_{{.OS}}_{{.Arch}}" -osarch="linux/amd64" github.com/Its-Alex/flatsharing/src/flatsharing/...

FROM amd64/alpine:3.8

# Copy executalle from builder
COPY --from=builder /flatsharing/fs_service_linux_amd64 /usr/local/bin/fs_service

EXPOSE 8080
EXPOSE 8081

CMD ["/usr/local/bin/fs_service"]