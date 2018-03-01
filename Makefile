GO=export GOPATH=$$(pwd) &&

.PHONY: dep
dep:
	$(GO) go get -u github.com/golang/dep/cmd/dep
	$(GO) go get -u github.com/alecthomas/gometalinter
	$(GO) gometalinter --install
	$(GO) cd src/flatsharing && dep ensure -v -vendor-only

.PHONY:build
build:
	$(GO) go install flatsharing/...

.PHONY: test
test:
	$(GO) go test flatsharing/...

.PHONY: lint
lint:
	$(GO) gometalinter --config=.gometalinter.json src/flatsharing/...

.PHONY: migrate-db
migrate-db:
	$(GO) go get -u -d github.com/mattes/migrate/cli github.com/lib/pq
	$(GO) go build -tags 'postgres' -o $$GOPATH/bin/migrate github.com/mattes/migrate/cli
	migrate -database postgres://flatsharing:611bukBNpbA3@localhost:5432/flatsharing?sslmode=disable -path ./migration up 1
