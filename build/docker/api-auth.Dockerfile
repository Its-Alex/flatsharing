FROM golang:1.12-alpine3.9 as builder

ENV GO111MODULE=on

COPY internal/api-auth /api-auth/internal/api-auth
COPY internal/core /api-auth/internal/core
COPY go.mod /api-auth/go.mod
COPY go.sum /api-auth/go.sum

WORKDIR /api-auth

# Install gox
RUN apk --no-cache add git ca-certificates gcc musl-dev && \
    go get -v github.com/Its-Alex/flatsharing/internal/... && \
    go get -v github.com/mitchellh/gox && \
    gox -output="api-auth_{{.OS}}_{{.Arch}}" -osarch="linux/amd64" github.com/Its-Alex/flatsharing/internal/api-auth/...

FROM amd64/alpine:3.8

# Copy executalle from builder
COPY --from=builder /flatsharing/api-auth_linux_amd64 /usr/local/bin/api-auth

EXPOSE 8080
EXPOSE 8081

CMD ["/usr/local/bin/api-auth"]