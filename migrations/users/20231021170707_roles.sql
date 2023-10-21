-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE TYPE user_role AS ENUM ('user', 'moderator', 'admin');

ALTER TABLE users
    ADD COLUMN role user_role NOT NULL DEFAULT 'user';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

ALTER TABLE users
    DROP COLUMN role;

DROP TYPE user_role;

-- +goose StatementEnd
