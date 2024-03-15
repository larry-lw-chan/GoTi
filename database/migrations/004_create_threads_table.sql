-- +goose Up
-- +goose StatementBegin
CREATE TABLE threads (
    "id" 	            INTEGER NOT NULL UNIQUE,
    "content"	  		TEXT NOT NULL,
	"thread_id"			INTEGER,
	"user_id"			INTEGER NOT NULL,
	"created_at"  		TEXT NOT NULL,
	"updated_at"	    TEXT NOT NULL,
	PRIMARY KEY("id" AUTOINCREMENT),
	FOREIGN KEY("thread_id") 	REFERENCES "threads"("id")
	FOREIGN KEY("user_id") 		REFERENCES "users"("id")
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE threads;
-- +goose StatementEnd