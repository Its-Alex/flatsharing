DOCKER_ENV=$(shell test ! -f /.dockerenv; echo "$$?")
OS := $(shell uname -s)

.PHONY: dep
dep:
	# Install protoc
ifeq ($(),Darwin)
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
	go install github.com/golang/protobuf/...
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

.PHONY: build-auth
build-auth:
	go build -v -o bin/auth-server -i github.com/Its-Alex/flatsharing/src/auth/server/...
	GOOS=linux GOARCH=amd64 go build -v -o bin/auth-server-linux-amd64 -i github.com/Its-Alex/flatsharing/src/auth/server/

.PHONY: build-support
build-support:
	go install -v github.com/Its-Alex/flatsharing/src/support/...

.PHONY: test
test:
	go test -v github.com/Its-Alex/flatsharing/src/...

.PHONY: lint
lint:
	golint src/...

.PHONY: coverage
coverage:
	go test -v github.com/Its-Alex/flatsharing/src/... -race -coverprofile=coverage.txt -covermode=atomic

.PHONY: protoc-auth
protoc-auth:
	protoc \
		auth.proto \
		-I src/protobuf/ \
		-I deps/protobuf/src \
		-I deps/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:src \
		--grpc-gateway_out=logtostderr=true:src \
		--swagger_out=logtostderr=true:src/auth/v1

.PHONY: migrate
migrate:
	@docker run -v $$(pwd)/migrations:/migrations --network host itsalex/migrate-docker \
		-path=/migrations/ -database postgres://flatsharing:611bukBNpbA3@localhost:5432/flatsharing?sslmode=disable up

.PHONY: clean
clean:
	rm -rf bin/ deps/ pkg/

.PHONY: enter-postgresql
enter-postgresql:
	docker exec -it --user postgres `docker-compose ps -q postgres` bash -c "export COLUMNS=`tput cols`; export LINES=`tput lines`; exec psql gometal"
