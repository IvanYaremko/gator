-- +goose up
CREATE TABLE users (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    name VARCHAR UNIQUE NOT NULL
);

-- +goose down
DROP TABLE users;