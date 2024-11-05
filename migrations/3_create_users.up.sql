-- Create "users" table
CREATE TABLE "users" (
  "id" text NOT NULL,
  "name" character varying(100) NULL,
  "email" character varying(100) NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_users_email" to table: "users"
CREATE UNIQUE INDEX "idx_users_email" ON "users" ("email");