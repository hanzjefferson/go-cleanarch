package usecase

import (
	"github.com/gofiber/fiber/v3"
	"github.com/hanzjefferson/go-cleanarch/internal/entity"
	"github.com/hanzjefferson/go-cleanarch/internal/model"
	"github.com/hanzjefferson/go-cleanarch/internal/repository"
	"github.com/hanzjefferson/go-cleanarch/pkg/hash"
	"github.com/hanzjefferson/go-cleanarch/pkg/jwt"
	"github.com/sirupsen/logrus"
)

type AuthUseCase struct {
	Log      *logrus.Logger
	JWT      *jwt.Provider
	UserRepo *repository.UserRepo
}

func NewAuthUseCase(
	log *logrus.Logger,
	jwt      *jwt.Provider,
	userRepo *repository.UserRepo,
) *AuthUseCase {
	return &AuthUseCase{
		Log:      log,
		JWT: jwt,
		UserRepo: userRepo,
	}
}

func (uc *AuthUseCase) Login(login *model.LoginRequest) (*model.AuthResponseData, error) {
	user := new(entity.User)
	err := uc.UserRepo.FindByUsername(user, login.Username)
	if err != nil {
		uc.Log.Warnf("user not found:%+v", err)
		return nil, fiber.ErrNotFound
	}

	if !hash.Compare(user.Password, login.Password) {
		uc.Log.Warnf("incorrect password:%+v", err)
		return nil, fiber.ErrBadRequest
	}

	token, err := uc.JWT.Generate(user.ID)
	if err != nil {
		uc.Log.Warnf("failed to generate jwt:%+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return &model.AuthResponseData{
		Token: token,
	}, nil
}

func (uc *AuthUseCase) Create(register *model.RegisterRequest) (*model.UserResponseData, error) {
	exist, err := uc.UserRepo.HasExist(register.Username)
	if err != nil {
		uc.Log.Warnf("username check query error:%+v", err)
		return nil, fiber.ErrInternalServerError
	}
	if exist {
		uc.Log.Warnf("user has exist:%+v", err)
		return nil, fiber.ErrConflict
	}

	hashedPassword, err := hash.Hash(register.Password)
	if err != nil {
		uc.Log.Warnf("failed to hashing password:%+v", err)
		return nil, fiber.ErrInternalServerError
	}

	user := &entity.User{
		Email:    register.Email,
		Username: register.Username,
		Password: hashedPassword,
	}
	if err := uc.UserRepo.Create(user); err != nil {
		return nil, fiber.ErrInternalServerError
	}

	return model.UserResponseDTO(user), nil
}
