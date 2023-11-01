package packets

const (
	CreatePacket = "INSERT INTO packets (id, message, created_at) VALUES ($1, $2, $3)"
)
