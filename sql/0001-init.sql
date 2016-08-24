-- cat sql/0001-init.sql | PGPASSWORD=pass psql -h localhost -Ucc cc_users

CREATE TABLE "users" (
    "user_id"          SERIAL PRIMARY KEY,
    "email"            VARCHAR(80) UNIQUE,
    "first_name"        VARCHAR(80),
    "last_name"        VARCHAR(80),
    "hash"             BYTEA,
    "salt"             BYTEA,
    "valid_jti"        BYTEA,
    "answers"          JSONB,
    "reset_hash"       BYTEA,
    "reset_expiration" TIMESTAMP
);
