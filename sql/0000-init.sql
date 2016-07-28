--' init.sql

-- cat sql/0000-init.sql | PGPASSWORD=pass psql -h localhost -Ucc cc_users

DROP VIEW IF EXISTS "leaders_public_footprint";
DROP VIEW IF EXISTS "leaders_public_food_footprint";
DROP VIEW IF EXISTS "leaders_public_housing_footprint";
DROP VIEW IF EXISTS "leaders_public_shopping_footprint";
DROP VIEW IF EXISTS "leaders_public_transport_footprint";
DROP INDEX IF EXISTS "leaders_public_footprint_index";
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
    "city"             VARCHAR(80),
    "state"            VARCHAR(80),
    "county"           VARCHAR(80),
    "total_footprint"  JSONB,
    "reset_hash"       BYTEA,
    "reset_expiration" TIMESTAMP
);

CREATE INDEX leaders_public_footprint_index
ON users(first_name, last_name, total_footprint, city, state, county)
WHERE public IS TRUE AND (total_footprint->'result_grand_total') IS NOT NULL;

CREATE VIEW leaders_public_footprint AS
SELECT first_name, last_name, total_footprint->'result_grand_total' as footprint, city, state, county
FROM users
WHERE public IS TRUE AND (total_footprint->'result_grand_total') IS NOT NULL
ORDER BY total_footprint->'result_grand_total' ASC;

CREATE VIEW leaders_public_food_footprint AS
SELECT first_name, last_name, total_footprint->'result_food_total' as footprint, city, state, county
FROM users
WHERE public IS TRUE AND (total_footprint->'result_food_total') IS NOT NULL
ORDER BY total_footprint->'result_food_total' ASC;

CREATE VIEW leaders_public_housing_footprint AS
SELECT first_name, last_name, total_footprint->'result_housing_total' as footprint, city, state, county
FROM users
WHERE public IS TRUE AND (total_footprint->'result_housing_total') IS NOT NULL
ORDER BY total_footprint->'result_housing_total' ASC;

CREATE VIEW leaders_public_shopping_footprint AS
SELECT first_name, last_name, (total_footprint->'result_shopping_total') as footprint, city, state, county
FROM users
WHERE public IS TRUE AND (total_footprint->'result_shopping_total') IS NOT NULL
ORDER BY total_footprint->'result_shopping_total' ASC;

CREATE VIEW leaders_public_transport_footprint AS
SELECT first_name, last_name, (total_footprint->'result_transport_total') as footprint, city, state, county
FROM users
WHERE public IS TRUE AND (total_footprint->'result_transport_total') IS NOT NULL
ORDER BY total_footprint->'result_transport_total' ASC;
