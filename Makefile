DOCKER_ENV=$(shell test ! -f /.dockerenv; echo "$$?")
OS := $(shell uname -s)

SWAGGER_UI_VERSION ?= v3.19.3

assert_out_docker:
ifeq ($(DOCKER_ENV),1)
	printf '\e[1;31m%-6s\e[m\n' "You must be outside of a container to use this command!"
	exit
endif

up:
	docker-compose up -d postgres
	docker-compose up wait_postgres
	make migrate
	docker-compose up -d workspace
	make dep protoc
	docker-compose up -d

dep:
ifeq ($(DOCKER_ENV),0)
	docker-compose exec -T workspace make -C . dep
else
	go install -v github.com/golang/protobuf/...
	go install -v github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway/...
	go install -v github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/...
	go install -v github.com/gobuffalo/packr/...
	go get -v golang.org/x/lint/golint
	go mod download
endif

build: build-auth build-flatsharing build-support

build-auth:
ifeq ($(DOCKER_ENV),0)
	docker-compose exec -T workspace make -C . build-auth
else
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -a -ldflags '-extldflags "-static"' -tags netgo -o bin/auth-server -i github.com/Its-Alex/flatsharing/src/auth/server/...
endif

build-flatsharing:
ifeq ($(DOCKER_ENV),0)
	docker-compose exec -T workspace make -C . build-flatsharing
else
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -a -ldflags '-extldflags "-static"' -tags netgo -o bin/flatsharing-server -i github.com/Its-Alex/flatsharing/src/flatsharing/server/...
endif

build-support:
ifeq ($(DOCKER_ENV),0)
	docker-compose exec -T workspace make -C . build-support
else
	go build -v -o bin/support-server -i github.com/Its-Alex/flatsharing/src/support/...
endif

test:
ifeq ($(DOCKER_ENV),0)
	docker-compose exec -T workspace make -C . test
else
	go test -v github.com/Its-Alex/flatsharing/src/...
endif

lint:
ifeq ($(DOCKER_ENV),0)
	docker-compose exec -T workspace make -C . lint
else
	golint src/...
endif

coverage:
ifeq ($(DOCKER_ENV),0)
	docker-compose exec -T workspace make -C . coverage
else
	go test -v github.com/Its-Alex/flatsharing/src/... -race -coverprofile=coverage.txt -covermode=atomic
endif

protoc: protoc-auth protoc-flatsharing

protoc-auth:
ifeq ($(DOCKER_ENV),0)
	docker-compose exec -T workspace make -C . protoc-auth
else
	mkdir -p src/auth/swagger
	protoc \
		auth.proto \
		-I src/protobuf/ \
		-I /usr/local/src/protobuf/src \
		-I /usr/local/src/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:. \
		--grpc-gateway_out=logtostderr=true:. \
		--swagger_out=logtostderr=true:src/auth/swagger
	cd src/auth && packr
endif

protoc-flatsharing:
ifeq ($(DOCKER_ENV),0)
	docker-compose exec -T workspace make -C . protoc-flatsharing
else
	mkdir -p src/flatsharing/swagger
	protoc \
		flatsharing.proto \
		-I src/protobuf/ \
		-I /usr/local/src/protobuf/src \
		-I /usr/local/src/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:. \
		--grpc-gateway_out=logtostderr=true:. \
		--swagger_out=logtostderr=true:src/flatsharing/swagger
	cd src/flatsharing && packr
endif

migrate: assert_out_docker
	@docker run --rm -v $$(pwd)/migrations:/migrations --network host itsalex/migrate-docker \
		-path=/migrations/ -database postgres://flatsharing:password@localhost:5432/flatsharing?sslmode=disable up

down: assert_out_docker
	docker-compose down

clean:
	rm -rf data/ bin/

import-swagger-ui: assert_out_docker
	if ! [[ -d "$$(pwd)/src/auth/swagger" ]]; then mkdir -p $$(pwd)/src/auth/swagger; fi
	if ! [[ -d "$$(pwd)/src/flatsharing/swagger" ]]; then mkdir -p $$(pwd)/src/flatsharing/swagger; fi
	docker run -it \
		-v `pwd`/swagger-ui:/mnt/ \
		--rm swaggerapi/swagger-ui:$(SWAGGER_UI_VERSION) \
		sh -c "rm -rf /mnt/*; cp -R /usr/share/nginx/html/* /mnt/"
	cp -r swagger-ui/* src/auth/swagger
	cp -r swagger-ui/* src/flatsharing/swagger
	docker-compose exec -T workspace bash -c "sed -i 's@https://petstore.swagger.io/v2/swagger.json@/auth.swagger.json@g' src/auth/swagger/index.html"
	docker-compose exec -T workspace bash -c "sed -i 's@https://petstore.swagger.io/v2/swagger.json@/flatsharing.swagger.json@g' src/flatsharing/swagger/index.html"
	rm -rf swagger-ui

enter: assert_out_docker
	docker-compose exec workspace bash -c "export COLUMNS=`tput cols`; export LINES=`tput lines`; exec bash"

enter-postgresql: assert_out_docker
	docker-compose exec --user postgres postgres bash -c "export COLUMNS=`tput cols`; export LINES=`tput lines`; exec psql -U flatsharing"

.PHONY: assert_out_docker up dep build build-auth build-flatsharing build-support test lint coverage protoc protoc-auth protoc-flatsharing migrate down clean import-swagger-ui enter enter-postgresql
