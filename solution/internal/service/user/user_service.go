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

// userService реализует всю бизнес-логику работы с пользователями
// Использую dependency injection для тестируемости и чистоты архитектуры
type userService struct {
	userRepo repository.UserRepository
	jwtSecret string
}

// NewUserService создает новый экземпляр сервиса пользователей
// TODO: подумать о добавлении логгера для трейсинга операций
func NewUserService(userRepo repository.UserRepository, jwtSecret string) *userService {
	return &userService{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

// Register регистрирует нового пользователя с валидацией и хешированием пароля
// Из прошлого проекта: всегда валидируем на нескольких уровнях для безопасности
func (s *userService) Register(ctx context.Context, req model.RegisterRequest) (model.AuthResponse, error) {
	// Валидация email на уникальность - критично для безопасности
	exists, err := s.userRepo.ExistsByEmailAndActive(ctx, req.Email)
	if err != nil {
		return model.AuthResponse{}, fmt.Errorf("failed to check email existence: %w", err)
	}
	if exists {
		return model.AuthResponse{}, fmt.Errorf("user with email %s already exists", req.Email)
	}

	// Хеширование пароля с bcrypt - проверенная временем практика
	// Cost factor 12 хороший баланс между скоростью и безопасностью
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.AuthResponse{}, fmt.Errorf("failed to hash password: %w", err)
	}

	// Создание пользователя с ролью USER по умолчанию
	user := model.NewUser(req.Email, string(passwordHash), req.FullName, model.UserRoleConst)
	
	// Установка опциональных полей если они переданы
	user.Age = req.Age
	user.Region = req.Region
	user.Gender = req.Gender
	user.MaritalStatus = req.MaritalStatus

	// Сохранение в базу данных
	if err := s.userRepo.Create(ctx, user); err != nil {
		return model.AuthResponse{}, fmt.Errorf("failed to create user: %w", err)
	}

	// Генерация JWT токена
	token, err := s.generateJWT(user.ID, user.Role)
	if err != nil {
		return model.AuthResponse{}, fmt.Errorf("failed to generate token: %w", err)
	}

	return model.AuthResponse{
		AccessToken: token,
		ExpiresIn:   3600, // 1 час в секундах
		User:        user,
	}, nil
}

// Login аутентифицирует пользователя по email и паролю
// Важно: проверяем включая неактивных пользователей для корректных ошибок
func (s *userService) Login(ctx context.Context, req model.LoginRequest) (model.AuthResponse, error) {
	// Ищем пользователя включая неактивных - нужно для проверки статуса
	user, err := s.userRepo.FindByEmailIncludingInactive(ctx, req.Email)
	if err != nil {
		return model.AuthResponse{}, fmt.Errorf("failed to find user: %w", err)
	}

	// Проверяем пароль с постоянным временем для защиты от timing attacks
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return model.AuthResponse{}, fmt.Errorf("invalid credentials")
	}

	// Проверяем активность пользователя
	if !user.IsActive {
		return model.AuthResponse{}, fmt.Errorf("user is deactivated")
	}

	// Генерация JWT токена
	token, err := s.generateJWT(user.ID, user.Role)
	if err != nil {
		return model.AuthResponse{}, fmt.Errorf("failed to generate token: %w", err)
	}

	return model.AuthResponse{
		AccessToken: token,
		ExpiresIn:   3600, // 1 час в секундах
		User:        user,
	}, nil
}

// GetMe возвращает профиль текущего пользователя
// Простая операция - пользователь уже аутентифицирован через JWT
func (s *userService) GetMe(ctx context.Context, userID string) (model.User, error) {
	// TODO: можно добавить кэширование для часто запрашиваемых профилей
	user, err := s.userRepo.FindByID(ctx, userID)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to get user profile: %w", err)
	}

	return user, nil
}

// UpdateMe полностью обновляет профиль пользователя
// Важно: все поля обязательны, для очистки нужно передать null
func (s *userService) UpdateMe(ctx context.Context, userID string, req model.UserUpdateRequest) (model.User, error) {
	// Сначала получаем текущего пользователя для сохранения неизменяемых полей
	existingUser, err := s.userRepo.FindByID(ctx, userID)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to find existing user: %w", err)
	}

	// Создаем обновленного пользователя с сохранением неизменяемых полей
	updatedUser := model.User{
		ID:           existingUser.ID,
		Email:        existingUser.Email,        // email нельзя менять
		PasswordHash: existingUser.PasswordHash, // пароль нельзя менять через этот метод
		FullName:     req.FullName,
		Age:          req.Age,
		Region:       req.Region,
		Gender:       req.Gender,
		MaritalStatus: req.MaritalStatus,
		Role:         existingUser.Role,     // обычный пользователь не может менять роль
		IsActive:     existingUser.IsActive, // обычный пользователь не может деактивировать себя
		CreatedAt:    existingUser.CreatedAt,
		UpdatedAt:    time.Now().UTC(), // обновляем время изменения
	}

	// Сохраняем в базу данных
	if err := s.userRepo.Update(ctx, updatedUser); err != nil {
		return model.User{}, fmt.Errorf("failed to update user: %w", err)
	}

	// Возвращаем обновленного пользователя
	return updatedUser, nil
}

// CreateByAdmin создает нового пользователя с правами администратора
// Важно: позволяет указывать роль пользователя
func (s *userService) CreateByAdmin(ctx context.Context, req model.UserCreateRequest) (model.User, error) {
	// Валидация email на уникальность - критично для безопасности
	exists, err := s.userRepo.ExistsByEmail(ctx, req.Email)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to check email existence: %w", err)
	}
	if exists {
		return model.User{}, fmt.Errorf("user with email %s already exists", req.Email)
	}

	// Хеширование пароля с bcrypt - проверенная временем практика
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to hash password: %w", err)
	}

	// Создание пользователя с указанной ролью
	user := model.NewUser(req.Email, string(passwordHash), req.FullName, req.Role)
	
	// Установка опциональных полей если они переданы
	user.Age = req.Age
	user.Region = req.Region
	user.Gender = req.Gender
	user.MaritalStatus = req.MaritalStatus

	// Сохранение в базу данных
	if err := s.userRepo.Create(ctx, user); err != nil {
		return model.User{}, fmt.Errorf("failed to create user: %w", err)
	}

	return user, nil
}

// GetByID возвращает пользователя по ID (только для админов)
func (s *userService) GetByID(ctx context.Context, userID string) (model.User, error) {
	user, err := s.userRepo.FindByIDIncludingInactive(ctx, userID)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to get user by ID: %w", err)
	}

	return user, nil
}

// UpdateByAdmin обновляет пользователя с полными правами администратора
func (s *userService) UpdateByAdmin(ctx context.Context, userID string, req model.UserUpdateRequest) (model.User, error) {
	// Сначала получаем текущего пользователя
	existingUser, err := s.userRepo.FindByIDIncludingInactive(ctx, userID)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to find existing user: %w", err)
	}

	// Создаем обновленного пользователя
	updatedUser := model.User{
		ID:           existingUser.ID,
		Email:        existingUser.Email,        // email нельзя менять
		PasswordHash: existingUser.PasswordHash, // пароль нельзя менять через этот метод
		FullName:     req.FullName,
		Age:          req.Age,
		Region:       req.Region,
		Gender:       req.Gender,
		MaritalStatus: req.MaritalStatus,
		Role:         existingUser.Role,     // по умолчанию сохраняем текущую роль
		IsActive:     existingUser.IsActive, // по умолчанию сохраняем текущий статус
		CreatedAt:    existingUser.CreatedAt,
		UpdatedAt:    time.Now().UTC(), // обновляем время изменения
	}

	// Применяем административные изменения если они указаны
	if req.Role != nil {
		updatedUser.Role = *req.Role
	}
	if req.IsActive != nil {
		updatedUser.IsActive = *req.IsActive
	}

	// Сохраняем в базу данных через административный метод
	if err := s.userRepo.UpdateByAdmin(ctx, updatedUser); err != nil {
		return model.User{}, fmt.Errorf("failed to update user by admin: %w", err)
	}

	// Возвращаем обновленного пользователя
	return updatedUser, nil
}

// SoftDelete деактивирует пользователя (только для админов)
func (s *userService) SoftDelete(ctx context.Context, userID string) error {
	if err := s.userRepo.SoftDelete(ctx, userID); err != nil {
		return fmt.Errorf("failed to soft delete user: %w", err)
	}

	return nil
}

// generateJWT создает JWT токен для пользователя
// Использую стандартные claim'ы и 1 час жизни
func (s *userService) generateJWT(userID string, role model.UserRole) (string, error) {
	// Воспроизводимость: всегда используем одно и то же время жизни
	expiresIn := time.Hour
	
	// Создаем токен с базовыми claim'ами
	token, err := jwt.GenerateToken(userID, string(role), s.jwtSecret, expiresIn)
	if err != nil {
		return "", fmt.Errorf("failed to generate JWT token: %w", err)
	}

	return token, nil
}

// validatePassword выполняет базовую валидацию пароля
// Из опыта: лучше проверять на нескольких уровнях
func (s *userService) validatePassword(password string) error {
	if len(password) < 8 {
		return fmt.Errorf("password must be at least 8 characters long")
	}
	if len(password) > 72 {
		return fmt.Errorf("password must be at most 72 characters long")
	}
	
	// Проверяем наличие цифры и буквы - требование из readme.txt
	hasLetter := false
	hasDigit := false
	
	for _, char := range password {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			hasLetter = true
		}
		if char >= '0' && char <= '9' {
			hasDigit = true
		}
	}
	
	if !hasLetter || !hasDigit {
		return fmt.Errorf("password must contain at least one letter and one digit")
	}
	
	return nil
}

// TODO: добавить методы для административных функций
// CreateByAdmin, GetByID, UpdateByAdmin, GetAll, SoftDelete
