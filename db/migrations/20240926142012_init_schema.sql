-- +goose Up
-- +goose StatementBegin
CREATE TABLE todo(
    id SERIAL8 PRIMARY KEY,
    title varchar(255) NOT NULL,
    text STRING NOT NULL,
    iscompleted bool DEFAULT FALSE NOT NULL,
    category varchar(255),
    deadline date
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE todo;

