# FlatSharing
FlatSharing configuration repository

# Requirements
`docker`

# How-to launch
```
$ docker-compose up -d
```
Use twice if mongo-express not working

# Mongo configuration

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

# License
[MIT](LICENSE)