package subscribers

const (
	CreateSubscriber      = "INSERT INTO subscribers (id, email, first_name, last_name, user_id, created_at) VALUES ($1, $2, $3, $4, $5, $6)"
	GetSubscriberCount    = "SELECT COUNT(*) FROM subscribers"
	FindAccountByEmailSQL = "SELECT id, email FROM subscribers WHERE email = $1 AND user_id = $2"
	GetSubscribers        = "SELECT id, email, first_name, last_name, user_id FROM subscribers WHERE user_id = $1 OFFSET $2 LIMIT $3"
)
