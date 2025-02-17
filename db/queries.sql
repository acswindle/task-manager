-- name: InsertUser :one
insert into users (name, salt, hashpassword)
values (?, ?, ?)
returning name;

-- name: GetUsers :many
select * from users;

-- name: GetCredentials :one
select salt, hashpassword from users
where name = ?;

-- name: InsertExpense :one
insert into expenses (user, category, amount, description)
values (?, ?, ?, ?)
returning id;

-- name: GetExpenses :many
select * from expenses
where user = ?;

-- name: GetExpensesByCategory :many
select * from expenses
where category = ? and user = ?;

-- name: GetExpensesByDate :many
select * from expenses
where created_date >= ? and user = ?;

-- name: GetExpensesByDateAndCategory :many
select * from expenses
where created_date >= ? and category = ? and user = ?;

-- name: DeleteExpense :exec
delete from expenses
where id = ?;

-- name: UpdateExpense :exec
update expenses
set category = ?, amount = ?, description = ?
where id = ?;
