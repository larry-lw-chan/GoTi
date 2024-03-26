-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
	"id"			INTEGER NOT NULL UNIQUE,
	"uuid"			TEXT NOT NULL UNIQUE,
	"email"			TEXT NOT NULL UNIQUE,
	"password"		TEXT NOT NULL,
	"created_at"	TEXT NOT NULL,
	"updated_at"	TEXT NOT NULL,
	PRIMARY KEY("id" AUTOINCREMENT)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
