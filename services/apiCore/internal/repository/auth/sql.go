package auth

const (
	createUserSQL      = "INSERT INTO users (email, password, id, created_at) VALUES ($1, $2, $3, $4)"
	findUserByEmailSQL = "SELECT id, email, password, created_at FROM users WHERE email = $1"
)
