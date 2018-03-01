[![codecov](https://codecov.io/gh/Its-Alex/flatsharing/branch/master/graph/badge.svg?token=3VmgQm5LGw)](https://codecov.io/gh/Its-Alex/flatsharing)

# FlatSharing
FlatSharing configuration repository at http://flatsharing.io

## Requirements
- `direnv`
- `docker`

## How-to launch
```
$ docker-compose up -d
```
Use twice if mongo-express not working

## Postgres configuration

To migrate database at first launch
```
$ make migrate-db
```

## "Models"

### Users

#### Database

| id                         | mail          | login       | username    | password | firstname   | lastname    | role |              date             |
|----------------------------|---------------|-------------|-------------|----------|-------------|-------------|------|:-----------------------------:|
| char(26)                   | varchar(128)  | varchar(36) | varchar(36) | bytea    | varchar(36) | varchar(36) | int  |          timestamptz          |
| 01C77BZC14YS004JDSP008MARS | test@test.com | Alex        | ItsAlex     | \x2...1  | Alexandre   | MARRE       | 0    | 2018-02-25 20:55:24.513498+00 |

#### Features

- Signup
- Signin
- List
- Update/Edit
- Remove

### Tokens

#### Database

|id           | token           | date        |
| ----------- | --------------- | ----------- |
| token1      | sadad879787978a | 2131213     |
| char(26)    | char(128)       | timestamptz |

#### Features

- Add
- List
- Update/Edit
- Remove

### Groups

#### Database

| id          | name          | date        |
| ----------- | ------------- | ----------- |
| group1      | Home at Paris | 2131213     |
| char(26)    | char(36)      | timestamptz |

#### Features

- Add group
- Remove group
- Update/Edit group
- Add users
- Remove users
- Update/Edit group
- ACL (Number/ Smaller has more permissions)

### Purchases (link to group)

#### Database

| id          | group-id | user-id   | buyer-id | amount        | shop        | desc    | date        |
| ----------- | -------- | --------- | -------- | ------------- | ----------- | ------- | ----------- |
| purchase1   | group1   | user1     | user2    | 53.75         | Super U     | Courses | 1517149821  |
| char(26)    | char(26) | char(26)  | char(26) | numeric (p=4) | varchar(36) | text    | timestamptz |

#### Features

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
