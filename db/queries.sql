-- name: InsertUser :one
insert into users (name)
values (?)
returning id;

-- name: GetUsers :many
select * from users;
