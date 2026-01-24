package user

import (
	"context"
	"solution/internal/config"
	"solution/internal/model"
	"solution/internal/repository"
	"solution/pkg/jwt"

	"go.uber.org/zap"
)

type UserService interface {
	Register(ctx context.Context, req model.RegisterRequest) (*model.RegisterResponse, error)
	Login(ctx context.Context, req model.LoginRequest) (*model.LoginResponse, error)
	CreateAdmin(ctx context.Context, email, passwordHash, fullName string) error

	GetMe(ctx context.Context, userID string) (*model.User, error)
	UpdateMe(ctx context.Context, userID string, req model.UserUpdateRequest) (*model.User, error)

	CreateByAdmin(ctx context.Context, req model.UserCreateRequest) (*model.User, error)
	GetByID(ctx context.Context, userID string) (*model.User, error)
	GetByIDIncludingInactive(ctx context.Context, userID string) (*model.User, error)
	UpdateByAdmin(ctx context.Context, userID string, req model.UserUpdateRequest) (*model.User, error)
	SoftDelete(ctx context.Context, userID string) error
	GetAll(ctx context.Context, page, size int) ([]*model.User, int, error)
	ValidateTokenAndGetUserID(token string) (string, error)
	ValidateTokenAndGetUserRole(token string) (string, error)
}

type userService struct {
	userRepo   repository.UserRepository
	jwtService *jwt.JWT
	logger     *zap.Logger
	jwtSecret  string
}

func NewUserService(userRepo repository.UserRepository, logger *zap.Logger) UserService {
	jwtSecret := config.AppConfig().RandomSecret.RANDOM_SECRET()
	jwtService := jwt.NewJWT(jwtSecret)

	return &userService{
		userRepo:   userRepo,
		jwtService: jwtService,
		logger:     logger,
		jwtSecret:  jwtSecret,
	}
}
