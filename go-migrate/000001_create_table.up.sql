create table products(
    id serial primary key,
    name varchar(255) not null
);

create table users(
    id serial primary key,
    number_phone varchar(255) unique not null
);