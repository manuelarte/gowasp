CREATE TABLE users (
    id integer NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    username text NOT NULL UNIQUE,
    password text NOT NULL,
    PRIMARY KEY (id),
    UNIQUE (username)
);