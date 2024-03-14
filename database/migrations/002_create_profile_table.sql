



-- +goose Up
-- +goose StatementBegin
CREATE TABLE profiles (
    "id" 	            INTEGER NOT NULL UNIQUE,
    "name"				TEXT,
	"bio"				TEXT,
	"link"				TEXT,
	"avatar"			TEXT,
	"user_id"			INTEGER UNIQUE NOT NULL,
	"created_at"  		TEXT NOT NULL,
	"updated_at"	    TEXT NOT NULL,
	PRIMARY KEY("id" AUTOINCREMENT),
	FOREIGN KEY("user_id") REFERENCES "users"("id")
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE profiles;
-- +goose StatementEnd
