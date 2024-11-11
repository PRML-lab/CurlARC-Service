-- Modify "records" table
ALTER TABLE "records" DROP CONSTRAINT "fk_teams_records", ADD
 CONSTRAINT "fk_teams_records" FOREIGN KEY ("team_id") REFERENCES "teams" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
-- Modify "user_teams" table
ALTER TABLE "user_teams" DROP CONSTRAINT "fk_user_teams_team", ADD
 CONSTRAINT "fk_user_teams_team" FOREIGN KEY ("team_id") REFERENCES "teams" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
