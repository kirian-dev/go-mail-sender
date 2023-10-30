package file

const (
	CreateFileSQL   = "INSERT INTO files (id, user_id, name, success_accounts, fail_accounts, loading_accounts, created_at, end_time, status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id;"
	FindFileByIDSQL = "SELECT * FROM files WHERE id = $1, user_id = $2"
	UpdateFileSQL   = "UPDATE files SET success_accounts = $2, fail_accounts = $3, loading_accounts = $4, end_time = $5, status = $6 WHERE id = $1"
	GetFilesSQL     = "SELECT * FROM files WHERE  user_id = $1"
)
