-- Create "records" table
CREATE TABLE "records" (
  "id" uuid NOT NULL,
  "team_id" text NULL,
  "result" character varying(10) NULL,
  "enemy_team_name" character varying(255) NULL,
  "place" character varying(255) NULL,
  "date" timestamp NULL,
  "ends_data_json" jsonb NULL,
  "is_red" boolean NULL,
  "is_public" boolean NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_teams_records" FOREIGN KEY ("team_id") REFERENCES "teams" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
