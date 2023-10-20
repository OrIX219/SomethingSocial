-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE TABLE users
(
    id                int primary key,
    name              varchar(255) not null default 'Unnamed',
    registration_date timestamptz not null default current_timestamp,
    last_login        timestamptz not null default current_timestamp,
    karma             int not null default 0,
    posts_count       int not null default 0
);

CREATE TABLE following
(
    follower_id int,
    follow_id   int,
    foreign key (follower_id) references users(id) on delete cascade,
    foreign key (follow_id) references users(id) on delete cascade
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DROP TABLE following;

DROP TABLE users;

-- +goose StatementEnd
