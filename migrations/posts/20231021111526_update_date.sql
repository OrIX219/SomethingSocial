-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

ALTER TABLE posts
    ADD COLUMN update_date timestamptz;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

ALTER TABLE posts
    DROP COLUMN update_date;

-- +goose StatementEnd
