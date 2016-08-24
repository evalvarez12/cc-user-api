
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

ALTER TABLE "users" ADD COLUMN "public" BOOLEAN,
                    ADD COLUMN "city" VARCHAR(80),
                    ADD COLUMN "state" VARCHAR(80),
                    ADD COLUMN "county" VARCHAR(80),
                    ADD COLUMN "total_footprint" JSONB;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE "users" DROP COLUMN "public",
                    DROP COLUMN "city",
                    DROP COLUMN "state",
                    DROP COLUMN "county",
                    DROP COLUMN "total_footprint";
