package user

import (
	"context"
	"fmt"
	"solution/internal/model"
	"solution/internal/repository"
	"solution/pkg/jwt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

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

func (s *userService) Register(ctx context.Context, req model.RegisterRequest) (*model.RegisterResponse, error) {
	exists, err := s.userRepo.ExistsByEmail(ctx, req.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to check email existence: %w", err)
	}
	if exists {
		return nil, fmt.Errorf("user with email %s already exists", req.Email)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	user := model.NewUser(req.Email, string(hashedPassword), req.FullName, model.UserRoleConst)

	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	token, err := s.generateJWT(user.ID, user.Role)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	return &model.RegisterResponse{
		User:  user,
		Token: token,
	}, nil
}

func (s *userService) Login(ctx context.Context, req model.LoginRequest) (*model.LoginResponse, error) {
	user, err := s.userRepo.FindByEmailIncludingInactive(ctx, req.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %w", err)
	}
	if user == nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	if !user.IsActive {
		return nil, fmt.Errorf("user account is deactivated")
	}

	token, err := s.generateJWT(user.ID, user.Role)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	return &model.LoginResponse{
		User:  *user,
		Token: token,
	}, nil
}

func (s *userService) GetMe(ctx context.Context, userID string) (*model.User, error) {
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user profile: %w", err)
	}
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

func (s *userService) UpdateMe(ctx context.Context, userID string, req model.UserUpdateRequest) (*model.User, error) {
	existingUser, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get existing user: %w", err)
	}
	if existingUser == nil {
		return nil, fmt.Errorf("user not found")
	}

	if req.FullName != nil {
		existingUser.FullName = *req.FullName
	}

	if req.Age != nil {
		existingUser.Age = req.Age
	}
	if req.Region != nil {
		existingUser.Region = req.Region
	}
	if req.Gender != nil {
		existingUser.Gender = req.Gender
	}
	if req.MaritalStatus != nil {
		existingUser.MaritalStatus = req.MaritalStatus
	}

	updatedUser, err := s.userRepo.Update(ctx, userID, req)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return updatedUser, nil
}

func (s *userService) CreateByAdmin(ctx context.Context, req model.UserCreateRequest) (*model.User, error) {
	exists, err := s.userRepo.ExistsByEmail(ctx, req.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to check email existence: %w", err)
	}
	if exists {
		return nil, fmt.Errorf("user with email %s already exists", req.Email)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	user := model.NewUser(req.Email, string(hashedPassword), req.FullName, req.Role)

	user.Age = req.Age
	user.Region = req.Region
	user.Gender = req.Gender
	user.MaritalStatus = req.MaritalStatus

	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return &user, nil
}

func (s *userService) GetByID(ctx context.Context, userID string) (*model.User, error) {
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by ID: %w", err)
	}
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

func (s *userService) UpdateByAdmin(ctx context.Context, userID string, req model.UserUpdateRequest) (*model.User, error) {
	return s.userRepo.UpdateByAdmin(ctx, userID, req)
}

func (s *userService) SoftDelete(ctx context.Context, userID string) error {
	if err := s.userRepo.SoftDelete(ctx, userID); err != nil {
		return fmt.Errorf("failed to soft delete user: %w", err)
	}

	return nil
}

func (s *userService) GetAll(ctx context.Context, page, size int) ([]*model.User, int, error) {
	if page < 0 {
		page = 0
	}
	if size <= 0 || size > 100 {
		size = 20 // значение по умолчанию
	}

	return s.userRepo.GetAll(ctx, page, size)
}

func (s *userService) generateJWT(userID string, role model.UserRole) (string, error) {
	expiresIn := time.Hour

	token, err := jwt.GenerateToken(userID, string(role), s.jwtSecret, expiresIn)
	if err != nil {
		return "", fmt.Errorf("failed to generate JWT token: %w", err)
	}

	return token, nil
}

func (s *userService) validatePassword(password string) error {
	if len(password) < 8 {
		return fmt.Errorf("password must be at least 8 characters long")
	}
	if len(password) > 72 {
		return fmt.Errorf("password must be at most 72 characters long")
	}

	hasDigit := false
	hasLetter := false
	for _, char := range password {
		if char >= '0' && char <= '9' {
			hasDigit = true
		}
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			hasLetter = true
		}
		if hasDigit && hasLetter {
			break
		}
	}

	if !hasDigit || !hasLetter {
		return fmt.Errorf("password must contain at least one digit and one letter")
	}

	return nil
}
