DOCKER_ENV=$(shell test ! -f /.dockerenv; echo "$$?")
OS := $(shell uname -s)

SWAGGER_UI_VERSION ?= v3.21.0

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
	go mod download
	go install -v github.com/golang/protobuf/...
	go install -v github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway/...
	go install -v github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/...
	go install -v github.com/gobuffalo/packr/...
	go install -v golang.org/x/lint/golint
endif

build: build-auth build-flatsharing build-support

build-auth:
ifeq ($(DOCKER_ENV),0)
	docker-compose exec -T workspace make -C . build-auth
else
	CGO_ENABLED=0 go build -v -mod=readonly -o assets/bin/auth-server -i github.com/Its-Alex/flatsharing/internal/auth/server/...
endif

build-flatsharing:
ifeq ($(DOCKER_ENV),0)
	docker-compose exec -T workspace make -C . build-flatsharing
else
	CGO_ENABLED=0 go build -v -mod=readonly -o assets/bin/flatsharing-server -i github.com/Its-Alex/flatsharing/internal/flatsharing/server/...
endif

build-support:
ifeq ($(DOCKER_ENV),0)
	docker-compose exec -T workspace make -C . build-support
else
	CGO_ENABLED=0 go build -v -mod=readonly -o bin/support-server -i github.com/Its-Alex/flatsharing/internal/support/...
endif

test:
ifeq ($(DOCKER_ENV),0)
	docker-compose exec -T workspace make -C . test
else
	go test -v -mod=readonly github.com/Its-Alex/flatsharing/internal/...
endif

lint:
ifeq ($(DOCKER_ENV),0)
	docker-compose exec -T workspace make -C . lint
else
	golint internal/...
endif

coverage:
ifeq ($(DOCKER_ENV),0)
	docker-compose exec -T workspace make -C . coverage
else
	go test -v -mod=readonly github.com/Its-Alex/flatsharing/internal/... -race -coverprofile=coverage.txt -covermode=atomic
endif

protoc: protoc-auth protoc-flatsharing

protoc-auth:
ifeq ($(DOCKER_ENV),0)
	docker-compose exec -T workspace make -C . protoc-auth
else
	mkdir -p internal/auth/swagger
	protoc \
		auth.proto \
		-I internal/protobuf/ \
		-I /usr/local/src/protobuf/src \
		-I /usr/local/src/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:. \
		--grpc-gateway_out=logtostderr=true:. \
		--swagger_out=logtostderr=true:internal/auth/swagger
	cd internal/auth && packr
endif

protoc-flatsharing:
ifeq ($(DOCKER_ENV),0)
	docker-compose exec -T workspace make -C . protoc-flatsharing
else
	mkdir -p internal/flatsharing/swagger
	protoc \
		flatsharing.proto \
		-I internal/protobuf/ \
		-I /usr/local/src/protobuf/src \
		-I /usr/local/src/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:. \
		--grpc-gateway_out=logtostderr=true:. \
		--swagger_out=logtostderr=true:internal/flatsharing/swagger
	cd internal/flatsharing && packr
endif

migrate: assert_out_docker
	@docker run --rm -v $$(pwd)/assets/postgres/migrations:/migrations --network host migrate/migrate:v4.2.5 \
		-path=/migrations/ -database postgres://flatsharing:password@127.0.0.1:5432/flatsharing?sslmode=disable up

down: assert_out_docker
	docker-compose down

clean:
	rm -rf assets/postgres/data/ assets/bin/

import-swagger-ui: assert_out_docker
	if ! [[ -d "$$(pwd)/internal/auth/swagger" ]]; then mkdir -p $$(pwd)/internal/auth/swagger; fi
	if ! [[ -d "$$(pwd)/internal/flatsharing/swagger" ]]; then mkdir -p $$(pwd)/internal/flatsharing/swagger; fi
	docker run -it \
		-v `pwd`/swagger-ui:/mnt/ \
		--rm swaggerapi/swagger-ui:$(SWAGGER_UI_VERSION) \
		sh -c "rm -rf /mnt/*; cp -R /usr/share/nginx/html/* /mnt/"
	cp -r swagger-ui/* internal/auth/swagger
	cp -r swagger-ui/* internal/flatsharing/swagger
	docker-compose exec -T workspace bash -c "sed -i 's@https://petstore.swagger.io/v2/swagger.json@/auth.swagger.json@g' internal/auth/swagger/index.html"
	docker-compose exec -T workspace bash -c "sed -i 's@https://petstore.swagger.io/v2/swagger.json@/flatsharing.swagger.json@g' internal/flatsharing/swagger/index.html"
	rm -rf swagger-ui

enter: assert_out_docker
	docker-compose exec workspace bash

enter-postgresql: assert_out_docker
	docker-compose exec --user postgres postgres bash -c "psql -U flatsharing"

.PHONY: assert_out_docker up dep build build-auth build-flatsharing build-support test lint coverage protoc protoc-auth protoc-flatsharing migrate down clean import-swagger-ui enter enter-postgresql
