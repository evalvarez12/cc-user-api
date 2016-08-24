
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE INDEX leaders_public_footprint_index
ON users(first_name, last_name, total_footprint, city, state, county)
WHERE public IS TRUE AND (total_footprint->'result_grand_total') IS NOT NULL;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP INDEX IF EXISTS "leaders_public_footprint_index";
