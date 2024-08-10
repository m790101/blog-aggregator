-- +goose Up
CREATE TABLE feeds (
    id UUID PRIMARY KEY NOT NULL,
    name varchar(255) NOT NULL,
    url varchar(255) UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE
);


-- +goose Down
DROP TABLE feeds;