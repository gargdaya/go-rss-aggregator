-- +goose Up
CREATE TABLE feed_follows (
    id UUID PRIMARY KEY,
    feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (feed_id, user_id)
);

-- +goose Down
DROP TABLE feed_follows;