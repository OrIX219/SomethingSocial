-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

ALTER TABLE users
    ADD COLUMN followers int NOT NULL DEFAULT 0,
    ADD COLUMN following int NOT NULL DEFAULT 0;

CREATE OR REPLACE FUNCTION follow() RETURNS TRIGGER AS $follow$
    BEGIN
        UPDATE users SET following=following+1 WHERE id=NEW.follower_id;
        UPDATE users SET followers=followers+1 WHERE id=NEW.follow_id;
        RETURN NULL;
    END;
$follow$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION unfollow() RETURNS TRIGGER AS $unfollow$
    BEGIN
        UPDATE users SET following=following-1 WHERE id=OLD.follower_id;
        UPDATE users SET followers=followers-1 WHERE id=OLD.follow_id;
        RETURN NULL;
    END;
$unfollow$ LANGUAGE plpgsql;

CREATE TRIGGER tr_follow AFTER INSERT ON following FOR EACH ROW EXECUTE PROCEDURE follow();
CREATE TRIGGER tr_unfollow AFTER DELETE ON following FOR EACH ROW EXECUTE PROCEDURE unfollow();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DROP TRIGGER tr_unfollow ON following;
DROP TRIGGER tr_follow ON following;

DROP FUNCTION unfollow();

DROP FUNCTION follow();

ALTER TABLE users
    DROP COLUMN following,
    DROP COLUMN followers;

-- +goose StatementEnd
