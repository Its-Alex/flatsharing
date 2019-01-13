# Database

## Models

### Users

| id                         | mail          | login       | username    | password | firstname   | lastname    | role |              date             |
|----------------------------|---------------|-------------|-------------|----------|-------------|-------------|------|-------------------------------|
| char(26)                   | varchar(128)  | varchar(36) | varchar(36) | bytea    | varchar(36) | varchar(36) | int  |          timestamptz          |
| 01C77BZC14YS004JDSP008MARS | test@test.com | Alex        | ItsAlex     | \x2...1  | Alexandre   | MARRE       | 0    | 2018-02-25 20:55:24.513498+00 |

### Tokens

|id           | token           | date        |
| ----------- | --------------- | ----------- |
| token1      | sadad879787978a | 2131213     |
| char(26)    | char(128)       | timestamptz |

### Flats

| id          | name          | date        |
| ----------- | ------------- | ----------- |
| home1       | Home at Paris | 2131213     |
| char(26)    | char(36)      | timestamptz |

### Users - Flats junction

| fk_user_id                 | fk_flat_id                 |
| -------------------------- | -------------------------- |
| 01C77BZC14YS004JDSP008MARS | 01C77BZC14YS004JDSP008MARS |
| char(26)                   | char(26)                   |

### Purchases

| id          | home-id  | user-id   | buyer-id | price         | shop        | desc    | date        |
| ----------- | -------- | --------- | -------- | ------------- | ----------- | ------- | ----------- |
| purchase1   | group1   | user1     | user2    | 53.75         | Super U     | Courses | 1517149821  |
| char(26)    | char(26) | char(26)  | char(26) | numeric (p=4) | varchar(36) | text    | timestamptz |

## Relations

![Sql graph](assets/bdd_schema.png)
