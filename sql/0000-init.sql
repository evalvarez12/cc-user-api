--' init.sql

-- cat sql/0000-init.sql | PGPASSWORD=pass psql -h localhost -Umazing financy

DROP TABLE IF EXISTS "users";

CREATE TABLE "users" (
    "user_id"         SERIAL PRIMARY KEY,
    "email"           VARCHAR(80) UNIQUE,
    "name"            VARCHAR(80),
    "hash"            BYTEA,
    "salt"            BYTEA,
    "default_account" INTEGER,
    "created_at"      TIMESTAMP,
    "updated_at"      TIMESTAMP
);

DROP TABLE IF EXISTS "charges";

CREATE TABLE "charges" (
    "id"            SERIAL PRIMARY KEY,
    "user_id"       INTEGER,
    "category_id"   INTEGER,
    "account_id"    INTEGER,
    "name"          VARCHAR(80),
    "description"   VARCHAR(80),
    "created_at"    TIMESTAMP,
    "updated_at"    TIMESTAMP,
    "expected_date" TIMESTAMP,
    "amount"        NUMERIC,
    "currency_code" VARCHAR(3),
    "exchange_rate" NUMERIC,
    "kind"          VARCHAR(80),
    "source"        VARCHAR(80),
    "destination"   VARCHAR(80),
    "balance"       NUMERIC
);

DROP TABLE IF EXISTS "planned_charges";

CREATE TABLE "planned_charges" (
    "id"            SERIAL PRIMARY KEY,
    "user_id"       INTEGER,
    "category_id"   INTEGER,
    "account_id"    INTEGER,
    "name"          VARCHAR(80),
    "description"   VARCHAR(80),
    "created_at"    TIMESTAMP,
    "updated_at"    TIMESTAMP,
    "expected_date" TIMESTAMP,
    "amount"        NUMERIC,
    "currency_code" VARCHAR(3),
    "exchange_rate" NUMERIC,
    "kind"          VARCHAR(80),
    "source"        VARCHAR(80),
    "destination"   VARCHAR(80),
    "periodicity"   INTEGER,
    "since"         TIMESTAMP,
    "until"         TIMESTAMP
);


DROP TABLE IF EXISTS "category";

CREATE TABLE "category" (
    "category_id" SERIAL PRIMARY KEY,
    "user_id"     INTEGER,
    "name"        VARCHAR(80),
    "created_at"  TIMESTAMP,
    "updated_at"  TIMESTAMP,
    "kind"        VARCHAR(80),
    "total"       NUMERIC
);

DROP TABLE IF EXISTS "account";

CREATE TABLE "account" (
    "account_id"  SERIAL PRIMARY KEY,
    "user_id"     INTEGER,
    "name"        VARCHAR(80),
    "bank"        VARCHAR(80),
    "description" VARCHAR(80),
    "created_at"  TIMESTAMP,
    "updated_at"  TIMESTAMP,
    "total"       NUMERIC
);

DROP TABLE IF EXISTS "sessions";

CREATE TABLE "sessions" (
    "user_id" INTEGER,
    "token"   VARCHAR(80) UNIQUE
);
