-- Migration: init
-- Created at: 2021-05-01 14:35:37 UTC
-- ====  UP  ====

BEGIN;

  CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
  CREATE EXTENSION IF NOT EXISTS citext;

  CREATE SCHEMA huddlet;

  CREATE DOMAIN huddlet."identifier" AS citext
    CHECK (
      length(VALUE) BETWEEN 3 AND 100
    );

  CREATE DOMAIN huddlet."name" AS citext
    CHECK (
      length(VALUE) BETWEEN 1 AND 100
    );


  CREATE DOMAIN huddlet."email" AS citext
    CHECK (
      VALUE ~ '^[a-zA-Z0-9.!#$%&''*+\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$'
    );

  CREATE DOMAIN huddlet."phone" AS varchar
    CHECK (
      VALUE ~ '^\+[1-9]\d{9,14}$'
    );

COMMIT;

-- ==== DOWN ====

BEGIN;

  DROP DOMAIN huddlet."phone";
  DROP DOMAIN huddlet."email";
  DROP DOMAIN huddlet."name";
  DROP DOMAIN huddlet."identifier";
  DROP SCHEMA huddlet;

COMMIT;
