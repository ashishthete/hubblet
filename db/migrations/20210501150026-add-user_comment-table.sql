-- Migration: add-user_comment-table
-- Created at: 2021-05-01 15:00:26 UTC
-- ====  UP  ====

BEGIN;

  CREATE TABLE huddlet.user_comment (
    id         uuid         PRIMARY KEY DEFAULT uuid_generate_v4(),
    post_id    uuid         REFERENCES huddlet.user_post (id)   ON DELETE CASCADE ON UPDATE CASCADE,
    user_id    uuid         REFERENCES huddlet.user_account (id)   ON DELETE CASCADE ON UPDATE CASCADE,

    comment    text         NOT NULL,
    created_at  timestamp   NOT NULL DEFAULT (now() AT TIME ZONE 'UTC')
  );

  CREATE INDEX user_comment_created_at_index ON huddlet.user_comment (created_at DESC);

  CREATE TABLE huddlet.user_comment_reaction (
    id              uuid        PRIMARY KEY DEFAULT uuid_generate_v4(),
    comment_id      uuid        REFERENCES huddlet.user_comment (id)   ON DELETE CASCADE ON UPDATE CASCADE,
    user_id         uuid        REFERENCES huddlet.user_account (id)   ON DELETE CASCADE ON UPDATE CASCADE,
    "like"            boolean     NOT NULL
  );

  CREATE TABLE huddlet.parent_child_comment (
    parent_id   uuid        REFERENCES huddlet.user_comment (id)   ON DELETE CASCADE ON UPDATE CASCADE,
    child_id    uuid        REFERENCES huddlet.user_comment (id)   ON DELETE CASCADE ON UPDATE CASCADE,
    PRIMARY KEY (parent_id, child_id)
  );

COMMIT;

-- ==== DOWN ====

BEGIN;

  DROP TABLE huddlet.parent_child_comment;
  DROP TABLE huddlet.user_comment_reaction;
  DROP TABLE huddlet.user_comment;

COMMIT;
