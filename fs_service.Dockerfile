FROM golang:1.11.1-alpine3.8 as builder

ENV GO111MODULE=on

COPY src/auth /flatsharing/src/flatsharing
COPY src/core /flatsharing/src/core
COPY go.mod /flatsharing/go.mod
COPY go.sum /flatsharing/go.sum

WORKDIR /flatsharing

# Install gox
RUN apk --no-cache add git ca-certificates gcc musl-dev && \
    go get -v github.com/Its-Alex/flatsharing/src/... && \
    go get -v github.com/mitchellh/gox && \
    gox -output="flatsharing_{{.OS}}_{{.Arch}}" -osarch="linux/amd64" github.com/Its-Alex/flatsharing/src/auth/...

FROM amd64/alpine:3.8

# Copy executalle from builder
COPY --from=builder /flatsharing/flatsharing_linux_amd64 /usr/local/bin/auth

EXPOSE 8080
EXPOSE 8081

CMD ["/usr/local/bin/auth"]