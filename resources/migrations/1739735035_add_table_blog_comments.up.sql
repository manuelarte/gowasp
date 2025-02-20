CREATE TABLE blog_comments (
    id integer NOT NULL,
    created_at datetime,
    updated_at datetime,
    posted_at datetime,
    blog_id integer NOT NULL,
    user_id integer NOT NULL,
    comment VARCHAR(1000) NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY(blog_id) REFERENCES blogs(id),
    FOREIGN KEY(user_id) REFERENCES users(id)
);