-- name: InsertUser :one
insert into users (name, salt, hashpassword)
values (?, ?, ?)
returning id;

-- name: GetUsers :many
select * from users;

-- name: GetCredentials :one
select salt, hashpassword from users
where name = ?;
