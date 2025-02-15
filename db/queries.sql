-- name: InsertUser :one
insert into users (name)
values (?)
returning id;
