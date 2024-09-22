-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
DROP CONSTRAINT users_email_key,
DROP CONSTRAINT users_username_key;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users ADD CONSTRAINT users_email_key UNIQUE (email),
ADD CONSTRAINT users_username_key UNIQUE (username);
-- +goose StatementEnd
