-- cat sql/0000-reset.sql | PGPASSWORD=pass psql -h localhost -Ucc cc_users

DROP VIEW IF EXISTS "leaders_public_footprint";
DROP VIEW IF EXISTS "leaders_public_food_footprint";
DROP VIEW IF EXISTS "leaders_public_housing_footprint";
DROP VIEW IF EXISTS "leaders_public_shopping_footprint";
DROP VIEW IF EXISTS "leaders_public_transport_footprint";
DROP INDEX IF EXISTS "leaders_public_footprint_index";
DROP TABLE IF EXISTS "users";
