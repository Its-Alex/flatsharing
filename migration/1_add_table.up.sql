CREATE TABLE users(
   id char(26) NOT NULL unique,
   mail varchar(128) NOT NULL unique,
   login varchar(36) NOT NULL unique,
   username varchar(36),
   firstname varchar(36),
   lastname varchar(36),
   password bytea,
   role int NOT NULL,
   date timestamptz NOT NULL DEFAULT NOW(),
   PRIMARY KEY(id)
);

CREATE TABLE tokens(
    id char(26) NOT NULL unique,
    token char(128) NOT NULL unique,
    date timestamptz NOT NULL DEFAULT NOW(),
    PRIMARY KEY(id)
);
