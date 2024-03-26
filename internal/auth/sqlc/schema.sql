CREATE TABLE users (
    "id" 	            INTEGER PRIMARY KEY,
	"uuid"				TEXT NOT NULL UNIQUE,
	"email"				TEXT NOT NULL UNIQUE,
	"password"			TEXT NOT NULL,
	"created_at"  		TEXT NOT NULL,
	"updated_at"	    TEXT NOT NULL
);