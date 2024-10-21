-- +goose Up
-- +goose StatementBegin
CREATE TABLE booked_court
(
    id       SERIAL PRIMARY KEY,
    court_id INT       NOT NULL,
    user_id  INT       NOT NULL,
    date     TIMESTAMP NOT NULL,
    duration BIGINT    NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (court_id) REFERENCES courts (id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE booked_court;
-- +goose StatementEnd
