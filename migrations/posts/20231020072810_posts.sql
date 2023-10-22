-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE TABLE posts
(
    id        uuid primary key,
    content   text not null,
    post_date timestamptz not null default current_timestamp,
    karma     int not null default 0,
    author    int not null
);

CREATE TABLE upvotes
(
    post_id uuid,
    user_id int,
    foreign key (post_id) references posts(id) on delete cascade
);

CREATE TABLE downvotes
(
    post_id uuid,
    user_id int,
    foreign key (post_id) references posts(id) on delete cascade
);

CREATE OR REPLACE FUNCTION upvote_insert() RETURNS TRIGGER AS $upvote_insert$
    BEGIN
        UPDATE posts SET karma=karma+1 WHERE id=NEW.post_id;
        RETURN NULL;
    END;
$upvote_insert$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION upvote_delete() RETURNS TRIGGER AS $upvote_delete$
    BEGIN
        UPDATE posts SET karma=karma-1 WHERE id=OLD.post_id;
        RETURN NULL;
    END;
$upvote_delete$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION downvote_insert() RETURNS TRIGGER AS $downvote_insert$
    BEGIN
        UPDATE posts SET karma=karma-1 WHERE id=NEW.post_id;
        RETURN NULL;
    END;
$downvote_insert$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION downvote_delete() RETURNS TRIGGER AS $downvote_delete$
    BEGIN
        UPDATE posts SET karma=karma+1 WHERE id=OLD.post_id;
        RETURN NULL;
    END;
$downvote_delete$ LANGUAGE plpgsql;

CREATE TRIGGER tr_upvote_insert AFTER INSERT ON upvotes FOR EACH ROW EXECUTE PROCEDURE upvote_insert();
CREATE TRIGGER tr_upvote_delete AFTER DELETE ON upvotes FOR EACH ROW EXECUTE PROCEDURE upvote_delete();
CREATE TRIGGER tr_downvote_insert AFTER INSERT ON downvotes FOR EACH ROW EXECUTE PROCEDURE downvote_insert();
CREATE TRIGGER tr_downvote_delete AFTER DELETE ON downvotes FOR EACH ROW EXECUTE PROCEDURE downvote_delete();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DROP TRIGGER tr_upvote_insert ON upvotes;
DROP TRIGGER tr_upvote_delete ON upvotes;
DROP TRIGGER tr_downvote_insert ON downvotes;
DROP TRIGGER tr_downvote_delete ON downvotes;

DROP TABLE posts;

DROP TABLE upvotes;

DROP TABLE downvotes;

DROP FUNCTION upvote_insert();

DROP FUNCTION upvote_delete();

DROP FUNCTION downvote_insert();

DROP FUNCTION downvote_delete();

-- +goose StatementEnd
