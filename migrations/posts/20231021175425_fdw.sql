-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE EXTENSION postgres_fdw;

CREATE SERVER users_server FOREIGN DATA WRAPPER postgres_fdw OPTIONS (
    host 'users-postgres', dbname 'postgres', port '5432', sslmode 'disable'
);

CREATE USER MAPPING FOR postgres SERVER users_server OPTIONS (
    user 'postgres', password '228'  
);

CREATE TYPE user_role AS ENUM ('user', 'moderator', 'admin');

CREATE FOREIGN TABLE users (
    id                int not null,
    name              varchar(255) not null default 'Unnamed',
    registration_date timestamptz not null default current_timestamp,
    last_login        timestamptz not null default current_timestamp,
    karma             int not null default 0,
    posts_count       int not null default 0,
    role user_role    not null default 'user'
)
SERVER users_server;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DROP FOREIGN TABLE users;

DROP TYPE user_role;

DROP USER MAPPING FOR postgres SERVER users_server;

DROP SERVER users_server;

DROP EXTENSION postgres_fdw;

-- +goose StatementEnd
