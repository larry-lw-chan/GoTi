CREATE TABLE relationships (
    "id" 	            INTEGER PRIMARY KEY,
    "follower_id"		INTEGER NOT NULL,
	"followee_id"		INTEGER NOT NULL,
	"created_at"  		TEXT NOT NULL,
	"updated_at"	    TEXT NOT NULL,
	FOREIGN KEY("follower_id") REFERENCES "users"("id")
	FOREIGN KEY("followee_id") REFERENCES "users"("id")
);