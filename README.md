[![codecov](https://codecov.io/gh/Its-Alex/flatsharing/branch/master/graph/badge.svg?token=3VmgQm5LGw)](https://codecov.io/gh/Its-Alex/flatsharing)

# FlatSharing

FlatSharing configuration repository at http://flatsharing.io

## Requirements

- `direnv`
- `docker`
- `Go 1.11`

If you don't want to get `direnv` you can source .envrc each time you open a new shell session

```
$ source .envrc
```

This repo use new go module feature, please read how it works before start

## Environments variables

All services have some environments variables to change some dynamics values.
All variables for all services all listed in [`docker-compose.yml`](/docker-compose.yml)

## Get started

```
$ make up
```

All services is up with a database

## Features

- Add purchases
- Remove purchases
- Edit purchases
- Sort by month/period

# Stack

## Ask

Use docker ?

Orchestrator ? (Rancher, ansible, kubernetes, puppet, chef)

Scale ? ([traefik](https://docs.traefik.io/))

# Refs

- [ulid](https://github.com/oklog/ulid) 128bytes/1024bits

---
[MIT](LICENSE)
