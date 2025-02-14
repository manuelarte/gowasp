CREATE TABLE users (
    id integer NOT NULL,
    created_at datetime,
    updated_at datetime,
    username VARCHAR(18) NOT NULL UNIQUE,
    password VARCHAR(18) NOT NULL,
    PRIMARY KEY (id),
    UNIQUE (username)
);