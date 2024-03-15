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

CREATE TABLE likes (
    "id" 	            INTEGER PRIMARY KEY,
	"thread_id"			INTEGER,
	"user_id"			INTEGER NOT NULL,
	"created_at"  		TEXT NOT NULL,
	"updated_at"	    TEXT NOT NULL,
	
	FOREIGN KEY("thread_id") 	REFERENCES "threads"("id")
	FOREIGN KEY("user_id") 		REFERENCES "users"("id")
);