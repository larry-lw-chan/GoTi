-- +goose Up
-- +goose StatementBegin
CREATE TABLE followers (
    "id" 	            INTEGER NOT NULL UNIQUE,
    "follower_id"		INTEGER NOT NULL,
	"followee_id"		INTEGER NOT NULL,  
	"created_at"  		TEXT NOT NULL,
	"updated_at"	    TEXT NOT NULL,
	PRIMARY KEY("id" AUTOINCREMENT),
	FOREIGN KEY("follower_id") REFERENCES "users"("id")
	FOREIGN KEY("followee_id") REFERENCES "users"("id")
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE followers;
-- +goose StatementEnd
