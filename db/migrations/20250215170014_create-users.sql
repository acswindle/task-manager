-- migrate:up
create table users (
  id integer primary key,
  name text not null,
  salt blob not null,
  hashpassword blob not null
);

-- migrate:down
drop table users;
