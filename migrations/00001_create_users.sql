-- +goose Up
create table users (
    id bigserial primary key,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp,
    email      text unique
);

-- +goose Down
drop table users;