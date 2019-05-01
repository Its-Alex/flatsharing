FROM golang:1.12-alpine3.9 as builder

ENV GO111MODULE=on

COPY internal/flatsharing /flatsharing/internal/flatsharing
COPY internal/auth /flatsharing/internal/auth
COPY internal/core /flatsharing/internal/core
COPY go.mod /flatsharing/go.mod
COPY go.sum /flatsharing/go.sum

WORKDIR /flatsharing

# Install gox
RUN apk --no-cache add git ca-certificates gcc musl-dev && \
    go mod download && \
    go get -v github.com/mitchellh/gox && \
    gox -output="fs_service_{{.OS}}_{{.Arch}}" -osarch="linux/amd64" github.com/Its-Alex/flatsharing/internal/flatsharing/...

FROM amd64/alpine:3.8

# Copy executalle from builder
COPY --from=builder /flatsharing/fs_service_linux_amd64 /usr/local/bin/fs_service

EXPOSE 8080
EXPOSE 8081

CMD ["/usr/local/bin/fs_service"]