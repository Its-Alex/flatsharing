GO=export GOPATH=$$(pwd) &&

.PHONY: up
up:
	@docker-compose up -d postgres
	@docker-compose run --rm wait_postgres
	@make migrate

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
	docker run --rm -v $$(pwd)/migrations:/migrations --network host migrate/migrate \
		-path=/migrations/ -database postgres://flatsharing:611bukBNpbA3@localhost:5432/flatsharing?sslmode=disable up

.PHONY: enter-postgresql
enter-postgresql:
	docker exec -it --user postgres `docker-compose ps -q postgres` bash -c "export COLUMNS=`tput cols`; export LINES=`tput lines`; exec psql flatsharing"

.PHONY: clean
clean:
	docker-compose down -v
	rm -rf data
