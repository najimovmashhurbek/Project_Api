 CREATE TABLE posts(
    id uuid,
    name varchar(30),
    description varchar(30),
    user_id uuid,
    PRIMARY KEY(id)
    );

CREATE TABLE medias(
    id uuid,
    link varchar(60),
    posts_id uuid,
    PRIMARY KEY(id),
    CONSTRAINT fk_users
    FOREIGN KEY(posts_id) 
    REFERENCES posts(id)
    ON DELETE CASCADE 
);