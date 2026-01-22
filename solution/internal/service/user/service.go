package user

import "solution/internal/repository"

type userService struct {
	userRepo  repository.UserRepository
	jwtSecret string
}

func NewUserService(userRepo repository.UserRepository, jwtSecret string) *userService {
	return &userService{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}
