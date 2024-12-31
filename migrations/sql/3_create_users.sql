-- +goose Up
CREATE TABLE "users" (
  "id" text NOT NULL,
  "name" character varying(100) NULL,
  "email" character varying(100) NULL,
  PRIMARY KEY ("id")
);

CREATE UNIQUE INDEX "idx_users_email" ON "users" ("email");

-- +goose Down
DROP TABLE "users";