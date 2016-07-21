--' init.sql

-- cat sql/0000-init.sql | PGPASSWORD=pass psql -h localhost -Ucc cc_users

DROP TABLE IF EXISTS "users";

CREATE TABLE "users" (
    "user_id"          SERIAL PRIMARY KEY,
    "email"            VARCHAR(80) UNIQUE,
    "first_name"       VARCHAR(80),
    "last_name"        VARCHAR(80),
    "hash"             BYTEA,
    "salt"             BYTEA,
    "valid_jti"        BYTEA,
    "answers"          JSONB,
    "public"           BOOLEAN,
    "location"         JSONB,
    "total_footprint"  JSONB,
    "reset_hash"       BYTEA,
    "reset_expiration" TIMESTAMP
);

CREATE INDEX leaders_public_footprint_index ON users(first_name, last_name, total_footprint, location) WHERE public IS TRUE AND total_footprint IS NOT NULL;
