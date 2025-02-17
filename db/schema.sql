CREATE TABLE IF NOT EXISTS "schema_migrations" (version varchar(128) primary key);
CREATE TABLE users (
  name text primary key,
  salt blob not null,
  hashpassword blob not null
);
CREATE TABLE categories (
  category text primary key
);
CREATE TABLE expenses (
   id integer primary key,
   user text references users(name),
   description text not null,
   category text references categories(category),
   amount decimal not null,
   created_date date not null default current_date
 );
-- Dbmate schema migrations
INSERT INTO "schema_migrations" (version) VALUES
  ('20250215170014'),
  ('20250217135647');
