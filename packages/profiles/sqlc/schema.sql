CREATE TABLE profiles (
    "id" 	            INTEGER PRIMARY KEY,
    "name"				TEXT,
	"bio"				TEXT,
	"link"				TEXT,
	"avatar"			TEXT,
	"private"			INTEGER DEFAULT 0,
	"user_id"			INTEGER UNIQUE NOT NULL,
	"created_at"  		TEXT NOT NULL,
	"updated_at"	    TEXT NOT NULL,
	FOREIGN KEY("user_id") REFERENCES "users"("id")
);