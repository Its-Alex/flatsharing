SET TIME ZONE 'UTC-02';

CREATE TABLE users(
   id char(26) NOT NULL UNIQUE,
   mail varchar(128) NOT NULL UNIQUE,
   login varchar(64) NOT NULL UNIQUE,
   username varchar(64),
   firstname varchar(64),
   lastname varchar(64),
   password bytea,
   role int NOT NULL,
   created_at timestamptz NOT NULL DEFAULT NOW(),
   PRIMARY KEY (id)
);

CREATE TABLE tokens(
    id char(26) NOT NULL,
    fk_user_id char(26) NOT NULL,
    token char(128) NOT NULL UNIQUE,
    created_at timestamptz NOT NULL DEFAULT NOW(),
    PRIMARY KEY (id),
    CONSTRAINT fk_user_id FOREIGN KEY (fk_user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE flats(
    id char(26) NOT NULL,
    name varchar(64) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE user_flat(
    fk_user_id char(26) NOT NULL,
    fk_flat_id char(26) NOT NULL,
    PRIMARY KEY (fk_user_id, fk_flat_id),
    CONSTRAINT fk_user_id FOREIGN KEY (fk_user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_flat_id FOREIGN KEY (fk_flat_id) REFERENCES flats(id) ON DELETE CASCADE
);

CREATE TABLE shops(
    id char(26) NOT NULL,
    name varchar(64) NOT NULL UNIQUE,
    PRIMARY KEY (id)
);

CREATE TABLE purchases(
    id char(26) NOT NULL,
    fk_flats_id char(26) NOT NULL,
    fk_user_id char(26) NOT NULL,
    fk_buyer_id char(26) NOT NULL,
    fk_shop_id char(26) NOT NULL,
    price numeric(8, 2) NOT NULL,
    description text NOT NULL,
    created_at timestamptz NOT NULL DEFAULT NOW(),
    PRIMARY KEY (id),
    CONSTRAINT fk_flats_id FOREIGN KEY (fk_flats_id) REFERENCES flats(id) ON DELETE CASCADE,
    CONSTRAINT fk_user_id FOREIGN KEY (fk_user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_buyer_id FOREIGN KEY (fk_buyer_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_shop_id FOREIGN KEY (fk_shop_id) REFERENCES shops(id),
    CONSTRAINT positive_price CHECK (price > 0)
);
