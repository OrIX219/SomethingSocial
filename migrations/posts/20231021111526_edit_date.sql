-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

ALTER TABLE posts
    ADD COLUMN edit_date timestamptz;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

ALTER TABLE posts
    DROP COLUMN edit_date;

-- +goose StatementEnd
