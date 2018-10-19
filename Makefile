DOCKER_ENV=$(shell test ! -f /.dockerenv; echo "$$?")
OS := $(shell uname -s)

up: dep
	docker-compose up -d postgres
	docker-compose up wait_postgres
	make migrate
	docker-compose up -d
	make build

dep:
	# Install protoc
ifeq ($(OS),Darwin)
	curl -s -L "https://github.com/protocolbuffers/protobuf/releases/download/v3.6.1/protoc-3.6.1-osx-x86_64.zip" | bsdtar -xf- bin/protoc
	chmod u+x bin/protoc
else
	curl -s -L "https://github.com/google/protobuf/releases/download/v3.6.1/protoc-3.6.1-linux-x86_64.zip" | bsdtar -xf- bin/protoc
	chmod u+x bin/protoc
endif
	# Install protobuf
	mkdir -p deps/
	curl -s -L "https://github.com/google/protobuf/archive/v3.6.1.tar.gz" | tar xvz
	mv protobuf-3.6.1/ deps/protobuf/
	curl -s -L "https://github.com/grpc-ecosystem/grpc-gateway/archive/v1.5.1.tar.gz" | tar xvz
	mv grpc-gateway-1.5.1/ deps/grpc-gateway/
	go get -v golang.org/x/lint/golint
	go get -v github.com/Its-Alex/flatsharing/src/...
	GOBIN=$(shell pwd)/bin go install -v github.com/golang/protobuf/...
	GOBIN=$(shell pwd)/bin go install -v github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	GOBIN=$(shell pwd)/bin go install -v github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

build: build-auth

build-auth:
	rm -rf bin/auth-server*
	go build -v -o bin/auth-server -i github.com/Its-Alex/flatsharing/src/auth/server/...
	GOOS=linux GOARCH=amd64 go build -v -o bin/auth-server-linux-amd64 -i github.com/Its-Alex/flatsharing/src/auth/server/

build-support:
	go install -v github.com/Its-Alex/flatsharing/src/support/...

test:
	go test -v github.com/Its-Alex/flatsharing/src/...

lint:
	golint src/...

coverage:
	go test -v github.com/Its-Alex/flatsharing/src/... -race -coverprofile=coverage.txt -covermode=atomic

protoc-auth:
	protoc \
		auth.proto \
		-I src/protobuf/ \
		-I deps/protobuf/src \
		-I deps/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:. \
		--grpc-gateway_out=logtostderr=true:. \
		--swagger_out=logtostderr=true:src/auth/v1

migrate:
	@docker run --rm -v $$(pwd)/migrations:/migrations --network host itsalex/migrate-docker \
		-path=/migrations/ -database postgres://flatsharing:password@localhost:5432/flatsharing?sslmode=disable up

down:
	docker-compose down
	rm -rf bin/ deps/ data/

enter-postgresql:
	docker exec -it --user postgres `docker-compose ps -q postgres` bash -c "export COLUMNS=`tput cols`; export LINES=`tput lines`; exec psql -U flatsharing"

.PHONY: up dep build build-auth build-support test lint coverage protoc-auth migrate down enter-postgresql
