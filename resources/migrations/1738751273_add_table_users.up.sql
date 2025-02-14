CREATE TABLE users (
    id integer NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    username VARCHAR(18) NOT NULL UNIQUE,
    password VARCHAR(18) NOT NULL,
    PRIMARY KEY (id),
    UNIQUE (username)
);