CREATE TABLE IF NOT EXISTS "schema_migrations" (version varchar(128) primary key);
CREATE TABLE users (
  id integer primary key,
  name text not null
);
-- Dbmate schema migrations
INSERT INTO "schema_migrations" (version) VALUES
  ('20250215170014');
