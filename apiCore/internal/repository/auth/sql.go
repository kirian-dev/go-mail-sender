package auth

const (
	createUserSQL   = "INSERT INTO users (email, password, id) VALUES ($1, $2, $3)"
	findUserByIdSQL = "SELECT id, email, password FROM users WHERE email = $1"
)
