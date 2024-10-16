-- +goose Up
-- +goose StatementBegin
CREATE TABLE courts
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE courts;
-- +goose StatementEnd
