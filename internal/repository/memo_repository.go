package repository

import (
	"database/sql"
	"fmt"

	"github.com/royxu/simplegin/v2/internal/model"
)

type MemoRepository struct {
	DB *sql.DB
}

func (mr *MemoRepository) CreateMemo(memo *model.Memo) (*model.Memo, error) {
	rows, err := mr.DB.Query(
		`INSERT INTO memos (
			title,
			content,
			is_public,
			user_id
		) VALUES ($1, $2, $3, $4)
		 RETURNING id, created_at, updated_at;`,
		memo.Title, memo.Content, memo.IsPublic, memo.UserID,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create memo: %w", err)
	}

	var newMemo model.Memo
	rows.Next()
	rows.Scan(
		&newMemo.ID,
		&newMemo.CreatedAt,
		&newMemo.UpdatedAt,
	)
	newMemo.Title = memo.Title
	newMemo.Content = memo.Content
	newMemo.UserID = memo.UserID
	newMemo.IsPublic = memo.IsPublic

	return &newMemo, nil
}
