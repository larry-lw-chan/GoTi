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

CREATE TABLE likes (
    "id" 	            INTEGER NOT NULL UNIQUE,
	"thread_id"			INTEGER,
	"user_id"			INTEGER NOT NULL,
	"created_at"  		TEXT NOT NULL,
	"updated_at"	    TEXT NOT NULL,

	PRIMARY KEY("id" AUTOINCREMENT),
	FOREIGN KEY("thread_id") 	REFERENCES "threads"("id")
	FOREIGN KEY("user_id") 		REFERENCES "users"("id")
);

CREATE TABLE hashtags (
    "id" 	            INTEGER NOT NULL UNIQUE,
    "hashtag"		  	TEXT NOT NULL,
	"created_at"  		TEXT NOT NULL,
	"updated_at"	    TEXT NOT NULL,

	PRIMARY KEY("id" AUTOINCREMENT)
);

CREATE TABLE threads_hashtags (
    "id" 	            INTEGER NOT NULL UNIQUE,
	"thread_id"			INTEGER NOT NULL,
	"hashtag_id"		INTEGER NOT NULL,
	"created_at"  		TEXT NOT NULL,
	"updated_at"	    TEXT NOT NULL,
	
	PRIMARY KEY("id" AUTOINCREMENT)
	FOREIGN KEY("thread_id") 	REFERENCES "threads"("id")
	FOREIGN KEY("hashtag_id") 	REFERENCES "hashtags"("id")
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE threads_hashtags;
DROP TABLE hashtags;
DROP TABLE likes;
DROP TABLE threads;
-- +goose StatementEnd