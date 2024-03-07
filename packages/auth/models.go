package auth

// User Model
type User struct {
	ID        int64
	Username  string
	Email     string
	Password  string
	CreatedAt string
	UpdatedAt string
}

// User Session
type UserSession struct {
	Username      string
	Authenticated bool
}
