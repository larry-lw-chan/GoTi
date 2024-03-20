CREATE TABLE threads (
    "id" 	            INTEGER PRIMARY KEY,
    "content"	  		TEXT NOT NULL,
	"thread_id"			INTEGER,
	"user_id"			INTEGER NOT NULL,
	"created_at"  		TEXT NOT NULL,
	"updated_at"	    TEXT NOT NULL,
	
	FOREIGN KEY("thread_id") 	REFERENCES "threads"("id")
	FOREIGN KEY("user_id") 		REFERENCES "users"("id")
);

CREATE TABLE hashtags (
    "id" 	            INTEGER PRIMARY KEY,
    "hashtag"		  	TEXT NOT NULL,
	"created_at"  		TEXT NOT NULL,
	"updated_at"	    TEXT NOT NULL
);

CREATE TABLE threads_hashtags (
    "id" 	            INTEGER PRIMARY KEY,
	"thread_id"			INTEGER,
	"hashtag_id"		INTEGER NOT NULL,
	"created_at"  		TEXT NOT NULL,
	"updated_at"	    TEXT NOT NULL,
	
	FOREIGN KEY("thread_id") 	REFERENCES "threads"("id"),
	FOREIGN KEY("hashtag_id") 	REFERENCES "hashtags"("id")
);

CREATE TABLE likes (
    "id" 	            INTEGER PRIMARY KEY,
	"thread_id"			INTEGER,
	"user_id"			INTEGER NOT NULL,
	"created_at"  		TEXT NOT NULL,
	"updated_at"	    TEXT NOT NULL,
	
	FOREIGN KEY("thread_id") 	REFERENCES "threads"("id")
	FOREIGN KEY("user_id") 		REFERENCES "users"("id")
);




-- Schema from other packages to allow for foreign key constraints
CREATE TABLE users (
    "id" 	            INTEGER PRIMARY KEY,
    "username"			TEXT NOT NULL UNIQUE,
	"email"				TEXT NOT NULL UNIQUE,
	"password"			TEXT NOT NULL,
	"created_at"  		TEXT NOT NULL,
	"updated_at"	    TEXT NOT NULL
);