-- +goose Up
CREATE TABLE "teams" (
  "id" text NOT NULL,
  "name" character varying(100) NULL,
  PRIMARY KEY ("id")
);

-- +goose Down
DROP TABLE "teams";