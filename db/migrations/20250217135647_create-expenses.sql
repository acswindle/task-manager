-- migrate:up
create table categories (
  category text primary key
);

insert into categories (category) values
('Housing'),
('Utilities'),
('Insurance'),
('Financing '),
('Groceries'),
('Medical'),
('Entertainment'),
('Travel'),
('Other');

 create table expenses (
   id integer primary key,
   user text references users(name),
   description text not null,
   category text references categories(category),
   amount decimal not null,
   created_date date not null default current_date
 );

-- migrate:down
drop table expenses;
drop table categories;
