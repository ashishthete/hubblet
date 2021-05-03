-- Migration: add-user_account-table
-- Created at: 2021-05-01 14:38:50 UTC
-- ====  UP  ====

BEGIN;

  CREATE TABLE huddlet.user_account (
    id          uuid              PRIMARY KEY DEFAULT uuid_generate_v4(),
    name        huddlet.name       NOT NULL,
    username    huddlet.identifier NOT NULL,
    password    text              NOT NULL,
    email       huddlet.email,

    created_at  timestamp         NOT NULL DEFAULT (now() AT TIME ZONE 'UTC')
  );

  CREATE UNIQUE INDEX user_account_username_index   ON huddlet.user_account (username);
  CREATE INDEX user_account_created_at_index ON huddlet.user_account (created_at DESC);

COMMIT;

-- ==== DOWN ====

BEGIN;

  DROP TABLE huddlet.user_account;

COMMIT;
