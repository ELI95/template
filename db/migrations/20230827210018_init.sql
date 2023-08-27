-- +goose Up
-- +goose StatementBegin
CREATE TABLE user (
	id INTEGER PRIMARY KEY,
	email TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL,
	created_at TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
Drop TABLE user;
-- +goose StatementEnd
