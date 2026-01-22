package user

import (
	"context"
	"solution/internal/model"
)

func (s *userService) GetAll(ctx context.Context, page, size int) ([]*model.User, int, error) {
	if page < 0 {
		page = 0
	}
	if size <= 0 || size > 100 {
		size = 20 // значение по умолчанию
	}

	return s.userRepo.GetAll(ctx, page, size)
}
