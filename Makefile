GO=export GOPATH=$$(pwd) &&

.PHONY: dep
dep:
	$(GO) go get -u github.com/golang/dep/cmd/dep
	$(GO) go get -u github.com/alecthomas/gometalinter
	$(GO) gometalinter --install
	$(GO) cd src/flatsharing && dep ensure -v -vendor-only

.PHONY:build
build:
	$(GO) go install -v flatsharing/auth
	$(GO) go install -v flatsharing/support

.PHONY: test
test:
	$(GO) go test -v flatsharing/...

.PHONY: lint
lint:
	$(GO) gometalinter --config=.gometalinter.json src/flatsharing/...

.PHONY: coverage
coverage:
	$(GO) go test flatsharing/... -race -coverprofile=coverage.txt -covermode=atomic

.PHONY: watch
watch:
	$(GO) reflex -R '^bin' -R '^pkg' -R '^src/(github.com|flatsharing/vendor){1}' \
		-r '^src/flatsharing/(auth|core){1}/[A-Za-z0-9/]+\.go$$' \
		-s -- go run src/flatsharing/auth/*.go

.PHONY: migrate
migrate:
	docker run -v $$(pwd)/migrations:/migrations --network host itsalex/migrate-docker \
		-path=/migrations/ -database postgres://flatsharing:611bukBNpbA3@localhost:5432/flatsharing?sslmode=disable up
