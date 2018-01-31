# FlatSharing
FlatSharing configuration repository at http://flatsharing.io

## Requirements
`docker`

## How-to launch
```
$ docker swarm init
$ docker stack deploy fs -c psql.yml/mongo.yml
```
Use twice if mongo-express not working

## Mongo configuration

Connect to **mongo**
```bash
$ docker exec -it {name} mongo -u admin -p 611bukBNpbA3 --authenticationDatabase admin
```

Create **database**
```
$ use flatsharing;
```

Add **user** who can read/write to database
```
$ db.createUser({'user':'flatsharing', 'pwd':'70z7AE7ZS0uiQxDH4ERRF5AR', roles:['readWrite']});
```

Or via mongo-express

- Connect to [mongo-express](http://localhost:8081)
- Create new database
- Add user to admin database
```json5
{
    "_id": "flatsharing.flatsharing", // database.user
    "user": "flatsharing",
    "db": "flatsharing",
    "credentials": {
        "SCRAM-SHA-1": {
            "iterationCount": 10000,
            "salt": "j85mqlwM/J45pvUJQWkdrA==",
            "storedKey": "kVpYiwatNje2jZNjygghQvyn/+4=",
            "serverKey": "xHSjGg0rq7zwVFlY9GWLmMEc8VY="
        }
    },
    "roles": [
        {
            "role": "readWrite", // readWrite or read
            "db": "flatsharing" // database
        }
    ]
}
```

It's ready to use ! :)
You can backup the database with keeping `data` folder in project who was generated at container start

## "Models"

### Users

#### Database

|id           | mail            | password     | nickname | role | date      |
| ----------- | --------------- | ------------ | -------- | ---- | --------- |
| user1       | user1@gmail.com | passBcrypted | Superman | 0    | 1223213   |
| bit         | string          | string       | string   | int  | int8      |

#### Features

- Signup
- Signin
- List
- Update/Edit
- Remove

### Tokens

#### Database

|id           | token           | date      |
| ----------- | --------------- | --------- |
| token1      | sadad879787978a | 2131213   |
| bit         | string          | int8      |

#### Features

- Add
- List
- Update/Edit
- Remove

### Groups

#### Database

| id          | name          | date      |
| ----------- | ------------- | --------- |
| group1      | Home at Paris | 2131213   |
| bit         | string        | int8      |

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

| id (ulid)   | group-id | user-id   | buyer-id | amount        | shop    | desc    | date       |
| ----------- | -------- | --------- | -------- | ------------- | ------- | ------- | ---------- |
| purchase1   | group1   | user1     | user2    | 53.75         | Super U | Courses | 1517149821 |
| bit         | bit      | bit       | bit      | numeric (p=4) | string  | string  | int8       |

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

# License
[MIT](LICENSE)
