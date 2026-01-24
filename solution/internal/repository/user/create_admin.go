package user

import (
	"context"
	"solution/internal/model"
)

func (r *repo) CreateAdmin(ctx context.Context, email, passwordHash, fullName string) error {
	admin := model.NewUser(email, passwordHash, fullName, model.AdminRole, nil, nil, nil, nil, nil)
	return r.Create(ctx, admin)
}
