-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
	"id"			INTEGER NOT NULL UNIQUE,
	"username"		TEXT NOT NULL UNIQUE,
	"email"			TEXT NOT NULL UNIQUE,
	"password"		TEXT NOT NULL,
	"created_at"	TEXT,
	"updated_at"	TEXT,
	PRIMARY KEY("id" AUTOINCREMENT)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
