-- migrate:up
create table users (
  id integer primary key,
  name text not null
);

-- migrate:down
drop table users;
