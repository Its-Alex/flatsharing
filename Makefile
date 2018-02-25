.PHONY: dep
dep:
	go get -u github.com/golang/dep/cmd/dep
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install
	cd src/flatsharing && dep ensure -v -vendor-only

.PHONY:build
build:
	go install flatsharing/...

.PHONY: test
test:
	go test flatsharing/...

.PHONY: lint
lint:
	gometalinter --config=.gometalinter.json src/flatsharing/...

.PHONY: migrate-db
migrate-db:
	go get -u -d github.com/mattes/migrate/cli github.com/lib/pq
	go build -tags 'postgres' -o $$GOPATH/bin/migrate github.com/mattes/migrate/cli
	migrate -database postgres://flatsharing:611bukBNpbA3@localhost:5432/flatsharing?sslmode=disable -path ./migration up 1
