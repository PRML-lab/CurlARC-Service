-- +goose Up
ALTER TABLE "public"."records" ADD COLUMN "is_red" boolean NULL;

-- +goose Down
ALTER TABLE "public"."records" DROP COLUMN IF EXISTS "is_red";