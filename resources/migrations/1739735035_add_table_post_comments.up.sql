CREATE TABLE post_comments (
    id integer NOT NULL,
    created_at datetime,
    updated_at datetime,
    posted_at datetime,
    post_id integer NOT NULL,
    user_id integer NOT NULL,
    comment VARCHAR(1000) NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY(post_id) REFERENCES posts(id),
    FOREIGN KEY(user_id) REFERENCES users(id)
);