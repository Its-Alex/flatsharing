# FlatSharing
FlatSharing configuration repository

## Requirements
`docker`

## How-to launch
```
$ docker-compose up -d
```
Use twice if mongo-express not working

## Mongo configuration

Connect to **mongo**
```bash
$ docker exec -it flatsharing_mongo_1 mongo -u admin -p 611bukBNpbA3 --authenticationDatabase admin
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
```json
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

##### User
|id           | mail            | password     | nickname | role | date      |
| ----------- | --------------- | ------------ | -------- | ---- | --------- |
| user1       | user1@gmail.com | passBcrypted | Superman | 0    | 1223213   |
| bytea       | string          | string       | string   | int  | int       |

#### Features

- Signup
- Signin
- List
- Update
- Remove

### Tokens

#### Database

|id           | token           | date      |
| ----------- | --------------- | --------- |
| token1      | sadad879787978a | 2131213   |
| bytea       | string          | int       |

#### Features

- Add
- xList
- Update
- Remove

### Groups

#### Database

| id          | name          | date      |
| ----------- | ------------- | --------- |
| group1      | Home at Paris | 2131213   |
| bytea(ulid) | string        | int       |

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

| id (ulid)   | group-id | user-id   | buyer-id | amount | shop    | desc    | date       |
| ----------- | -------- | --------- | -------- | ------ | ------- | ------- | ---------- |
| purchase1   | group1   | user1     | user2    | 53.75  | Super U | Courses | 1517149821 |
| bytea       | bytea    | bytea     | bytea    | float  | string  | string  | int        |

#### Features

- Add purchases
- Remove purchases
- Edit purchases
- Sort by month/period
- Shop list

# Refs

- ulid: https://github.com/oklog/ulid

# License
[MIT](LICENSE)