-- +goose Up
ALTER TABLE "public"."records" ADD COLUMN "is_red" boolean NOT NULL DEFAULT false;

-- +goose Down
ALTER TABLE "public"."records" DROP COLUMN IF EXISTS "is_red";