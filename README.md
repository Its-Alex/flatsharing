[![codecov](https://codecov.io/gh/Its-Alex/flatsharing/branch/master/graph/badge.svg?token=3VmgQm5LGw)](https://codecov.io/gh/Its-Alex/flatsharing)

FlatSharing backend repository at http://flatsharing.io

- [Requirements](#requirements)
- [Environments variables](#environments-variables)
- [Get started](#get-started)
- [How to dev](#how-to-dev)
- [Migrations](#migrations)
- [License](#license)

## Requirements

* `make`
- `docker`
- `direnv` (optional)

If you don't want to get `direnv` you can source .envrc each time you open a new shell session

```
$ source .envrc
```

This repo use new go module feature, please read how it works before start

You need nothing special in your host, this project use docker to dev/build

## Environments variables

All services have some environments variables to change some dynamics values
(usefull to setup and scale).

All variables for all services all listed in [`docker-compose.yml`](/docker-compose.yml)

## Get started

To up all services

```
$ make up
```

Then all services run in a container

You can fetch service with shared port in [`docker-compose.yml`](/docker-compose.yml)

When you build a service it's auto-refreshed in the container

There are global rule that works for all services

```
$ make build
```

This command will build all current services that can be build

You can stop project (containers) with this command:

```
$ make down
```

## How to dev

There are a container [`workspace`](/docker-compose.yml#L4) that is used to work inside it

Some makefile commands must be executed inside or outside this container, it will display
a message if your wrong about where you execute a command.

There is a rule to install dependencies inside container

```
$ make dep
```

So if you need to add a binary for the project see [`here`](/Makefile#L20)
or you can add it when container builds in [`Dockerfile`](/Dockerfile)

To enter inside workspace please use:

```
$ make enter
```

This command patch some annoying bugs. Same to entrer inside database

```
$ make enter-postgresql
```

This project use [`gRPC`](https://github.com/grpc/grpc-go) you can find code in 
[`src`](/src) folder and proto files in [`src/protobuf`](/src/protobuf). There are commands
to build proto

```
$ make protoc-${service}
```

Same to build go code

```
$ make build-${service}
```

You can check if tests failed:

```
$ make test
```

You can see if the linter pass:

```
$ make lint
```

Same for coverage:

```
$ make coverage
```

## Migrations

All databases migrations are defined in [`migrations`](/migrations) folder

You will find a up and down, they are used by [`migrate`](https://github.com/golang-migrate/migrate)

To apply all migrations:

```
$ make migrate
```

## License

[MIT](LICENSE)
