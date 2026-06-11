package model

type Memo struct {
	ID        int64  `json:"id"`
	Title     string `json:"username"`
	Content   string `json:"content"`
	IsPublic  bool   `json:"is_public"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
	UserID    int64  `json:"user_id"`
}
