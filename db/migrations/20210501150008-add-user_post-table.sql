-- Migration: add-user_post-table
-- Created at: 2021-05-01 15:00:08 UTC
-- ====  UP  ====

BEGIN;

  CREATE TABLE huddlet.user_post (
    id          uuid        PRIMARY KEY DEFAULT uuid_generate_v4(),
    title       text        NOT NULL,
    post        text        NOT NULL,
    user_id     uuid         REFERENCES huddlet.user_account (id)   ON DELETE CASCADE ON UPDATE CASCADE,

    created_at  timestamp         NOT NULL DEFAULT (now() AT TIME ZONE 'UTC')
  );

  CREATE INDEX user_post_created_at_index ON huddlet.user_post (created_at DESC);

  CREATE TABLE huddlet.user_post_reaction (
    id          uuid        PRIMARY KEY DEFAULT uuid_generate_v4(),
    post_id     uuid        REFERENCES huddlet.user_post (id)   ON DELETE CASCADE ON UPDATE CASCADE,
    user_id     uuid        REFERENCES huddlet.user_account (id)   ON DELETE CASCADE ON UPDATE CASCADE,
    "like"        boolean     NOT NULL
  );

COMMIT;

-- ==== DOWN ====

BEGIN;

  DROP TABLE huddlet.user_post_reaction;
  DROP TABLE huddlet.user_post;

COMMIT;
