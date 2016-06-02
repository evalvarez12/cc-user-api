--' init.sql

-- cat sql/0000-init.sql | PGPASSWORD=pass psql -h localhost -Umazing financy

DROP TABLE IF EXISTS "users";

CREATE TABLE "users" (
    "user_id"         SERIAL PRIMARY KEY,
    "email"           VARCHAR(80) UNIQUE,
    "fist_name"       VARCHAR(80),
    "last_name"       VARCHAR(80),
    "hash"            BYTEA,
    "salt"            BYTEA,
    "valid_jti"       VARCHAR(80),
    "answers"         JSONB
);
