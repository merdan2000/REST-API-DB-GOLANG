-- +goose Up
-- +goose StatementBegin
CREATE TABLE if not exists users(
    id serial not null primary key,
    first_name varchar(255) not null,
    last_name varchar(255) not null,
    email varchar(255) not null unique,
    age int,
    password varchar(255) not null,
    created timestamp not null default now()
);


-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE if exists users;
-- +goose StatementEnd
