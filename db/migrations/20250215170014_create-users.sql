-- migrate:up
create table users (
  name text primary key,
  salt blob not null,
  hashpassword blob not null
);

-- migrate:down
drop table users;
