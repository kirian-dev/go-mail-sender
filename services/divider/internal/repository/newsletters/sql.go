package newsletters

const (
	CreateNewsletter = "INSERT INTO newsletters (id, message, user_id, created_at) VALUES ($1, $2, $3, $4)"
)
