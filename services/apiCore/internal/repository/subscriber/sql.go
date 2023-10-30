package subscriber

const (
	CreateSubscriber      = "INSERT INTO subscribers (id, email, first_name, last_name, user_id, created_at) VALUES ($1, $2, $3, $4, $5, $6)"
	GetSubscriberCount    = "SELECT COUNT(*) FROM subscribers"
	FindAccountByEmailSQL = "SELECT id, email FROM subscribers WHERE email = $1"
)
