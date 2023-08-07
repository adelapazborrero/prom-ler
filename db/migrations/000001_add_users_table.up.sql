create table if not exists users (
    id serial primary key,
    username varchar(250),
    email varchar(250)
);
