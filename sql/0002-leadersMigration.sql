-- cat sql/0002-leadersMigration.sql | PGPASSWORD=pass psql -h localhost -Ucc cc_users

ALTER TABLE "users" ADD COLUMN "public" BOOLEAN,
                    ADD COLUMN "city" VARCHAR(80),
                    ADD COLUMN "state" VARCHAR(80),
                    ADD COLUMN "county" VARCHAR(80),
                    ADD COLUMN "total_footprint" JSONB;

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
