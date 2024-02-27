-- +goose Up
-- +goose StatementBegin
CREATE TABLE Users (
	"id"	INTEGER NOT NULL UNIQUE,
	"email"	TEXT NOT NULL UNIQUE,
	"password"	TEXT NOT NULL,
	PRIMARY KEY("id" AUTOINCREMENT)
);

INSERT INTO users VALUES
(0, 'test1@gmail.com', '123456'),
(1, 'test2@gmail.com', '123456');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE Users;
-- +goose StatementEnd
