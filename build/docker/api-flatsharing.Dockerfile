FROM golang:1.12-alpine3.9 as builder

ENV GO111MODULE=on

COPY internal/api-flatsharing /api-flatsharing/internal/api-flatsharing
COPY internal/api-auth /api-flatsharing/internal/api-auth
COPY internal/core /api-flatsharing/internal/core
COPY go.mod /api-flatsharing/go.mod
COPY go.sum /api-flatsharing/go.sum

WORKDIR /api-flatsharing

# Install gox
RUN apk --no-cache add git ca-certificates gcc musl-dev && \
    go mod download && \
    go get -v github.com/mitchellh/gox && \
    gox -output="api-flatsharing_{{.OS}}_{{.Arch}}" -osarch="linux/amd64" github.com/Its-Alex/api-flatsharing/internal/api-flatsharing/...

FROM amd64/alpine:3.8

# Copy executalle from builder
COPY --from=builder /api-flatsharing/api-flatsharing_linux_amd64 /usr/local/bin/api-flatsharing

EXPOSE 8080
EXPOSE 8081

CMD ["/usr/local/bin/api-flatsharing"]