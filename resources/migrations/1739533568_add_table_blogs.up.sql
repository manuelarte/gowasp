CREATE TABLE blogs (
    id integer NOT NULL,
    created_at datetime,
    updated_at datetime,
    posted_at datetime,
    user_id integer NOT NULL,
    title VARCHAR(200) NOT NULL,
    content VARCHAR(2000) NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY(user_id) REFERENCES users(id)
);