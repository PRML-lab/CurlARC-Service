-- Create "user_teams" table
CREATE TABLE "user_teams" (
  "user_id" text NOT NULL,
  "team_id" text NOT NULL,
  "state" character varying(100) NULL,
  PRIMARY KEY ("user_id", "team_id"),
  CONSTRAINT "fk_user_teams_team" FOREIGN KEY ("team_id") REFERENCES "teams" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_user_teams_user" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);