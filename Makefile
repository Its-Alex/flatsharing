DOCKER_ENV=$(shell test ! -f /.dockerenv; echo "$$?")
OS := $(shell uname -s)

SWAGGER_UI_VERSION ?= v3.22.1

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
	make dep protoc build
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

build: build-api-auth build-api-flatsharing build-support

build-api-auth:
ifeq ($(DOCKER_ENV),0)
	docker-compose exec -T workspace make -C . build-api-auth
else
	cd internal/api-auth && packr
	CGO_ENABLED=0 go build -v -mod=readonly -o assets/bin/api-auth-server -i github.com/Its-Alex/flatsharing/internal/api-auth/server/...
endif

build-api-flatsharing:
ifeq ($(DOCKER_ENV),0)
	docker-compose exec -T workspace make -C . build-api-flatsharing
else
	cd internal/api-flatsharing && packr
	CGO_ENABLED=0 go build -v -mod=readonly -o assets/bin/api-flatsharing-server -i github.com/Its-Alex/flatsharing/internal/api-flatsharing/server/...
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

protoc: protoc-api-auth protoc-api-flatsharing

protoc-api-auth:
ifeq ($(DOCKER_ENV),0)
	docker-compose exec -T workspace make -C . protoc-api-auth
else
	mkdir -p internal/api-auth/swagger
	protoc \
		auth.proto \
		-I internal/protobuf/ \
		-I /usr/local/src/protobuf/src \
		-I /usr/local/src/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:. \
		--grpc-gateway_out=logtostderr=true:. \
		--swagger_out=logtostderr=true:internal/api-auth/swagger
	cd internal/api-auth && packr
endif

protoc-api-flatsharing:
ifeq ($(DOCKER_ENV),0)
	docker-compose exec -T workspace make -C . protoc-api-flatsharing
else
	mkdir -p internal/api-flatsharing/swagger
	protoc \
		flatsharing.proto \
		-I internal/protobuf/ \
		-I /usr/local/src/protobuf/src \
		-I /usr/local/src/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:. \
		--grpc-gateway_out=logtostderr=true:. \
		--swagger_out=logtostderr=true:internal/api-flatsharing/swagger
endif

docker-build-workspace: assert_out_docker
	@docker build -f build/docker/workspace.Dockerfile -t workspace:latest .

docker-build-api-auth: assert_out_docker
	@docker build -f build/docker/api-auth.Dockerfile -t api-auth:latest .

docker-build-api-flatsharing: assert_out_docker
	@docker build -f build/docker/api-flatsharing.Dockerfile -t api-flatsharing:latest .

migrate: assert_out_docker
	@docker run --rm -v $$(pwd)/assets/postgres/migrations:/migrations --network host migrate/migrate:v4.2.5 \
		-path=/migrations/ -database postgres://flatsharing:password@127.0.0.1:5432/flatsharing?sslmode=disable up

down: assert_out_docker
	docker-compose down --remove-orphans -v
	rm -rf assets/postgres/data

clean:
	rm -rf assets/postgres/data/ assets/bin/

import-swagger-ui: assert_out_docker
	if ! [[ -d "$$(pwd)/assets/swagger-ui" ]]; then \
		mkdir -p $$(pwd)/assets/swagger-ui; \
		docker run -it \
			-v `pwd`/assets/swagger-ui:/mnt/ \
			--rm swaggerapi/swagger-ui:$(SWAGGER_UI_VERSION) \
			sh -c "rm -rf /mnt/*; cp -R /usr/share/nginx/html/* /mnt/"; \
		docker-compose exec -T workspace bash -c "sed -i 's@https://petstore.swagger.io/v2/swagger.json@/swagger.json@g' assets/swagger-ui/index.html"; \
	fi

enter: assert_out_docker
	docker-compose exec workspace bash

enter-postgresql: assert_out_docker
	docker-compose exec --user postgres postgres bash -c "psql -U flatsharing"

.PHONY: assert_out_docker up dep build build-api-auth build-api-flatsharing build-support test lint coverage protoc protoc-api-auth protoc-api-flatsharing migrate down clean import-swagger-ui enter enter-postgresql
