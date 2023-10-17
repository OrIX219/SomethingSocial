-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE TABLE users
(
    id            serial primary key,
    username      varchar(255) not null,
    password_hash varchar(255) not null
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DROP TABLE users;

-- +goose StatementEnd
