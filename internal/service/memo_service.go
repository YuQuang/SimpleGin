package service

import (
	"github.com/royxu/simplegin/v2/internal/model"
	"github.com/royxu/simplegin/v2/internal/repository"
)

type MemoService struct {
	MemoRepository *repository.MemoRepository
}

func (ms *MemoService) CreateMemo(
	title string,
	content string,
	userID int,
	isPublic bool,
) (*model.Memo, error) {
	memo, err := ms.MemoRepository.CreateMemo(
		&model.Memo{
			Title:    title,
			Content:  content,
			UserID:   int64(userID),
			IsPublic: isPublic,
		},
	)
	if err != nil {
		return nil, err
	}

	return memo, nil
}
