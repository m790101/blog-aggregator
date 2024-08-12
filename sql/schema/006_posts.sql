-- +goose Up
CREATE TABLE posts (
    id UUID PRIMARY KEY NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title varchar(255) NOT NULL,
    url varchar(255) UNIQUE NOT NULL,
    description text NOT NULL,
    published_at TIMESTAMP NOT NULL,
    feed_id UUID NOT NULL REFERENCES feeds(id) on DELETE CASCADE
);


-- +goose Down
DROP TABLE posts;
