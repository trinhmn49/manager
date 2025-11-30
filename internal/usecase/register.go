package usecase

import (
	"context"
	"manager/internal/domain/entity"
	"manager/internal/domain/repository"
	"manager/pkg/hash"
	"manager/pkg/logger"
)

type RegisterUseCase interface {
	Execute(ctx context.Context, input RegisterInput) error
}

type RegisterInput struct {
	DisplayName string  `json:"display_name"`
	Avatar      string  `json:"avatar"`
	Phone       string  `json:"phone"`
	Password    string  `json:"password"`
	Email       *string `json:"email"`
}

type registerUseCaseImpl struct {
	userRepo repository.UserRepo
	hasher   hash.PasswordHasher
}

func (r *registerUseCaseImpl) Execute(ctx context.Context, input RegisterInput) error {
	logger.Debug("start execute register usecase")
	hashPwd, err := r.hasher.Hash(input.Password)
	if err != nil {
		return err
	}

	if err := r.userRepo.Create(ctx, entity.UserEntity{
		DisplayName: input.DisplayName,
		Avatar:      input.Avatar,
		Phone:       input.Phone,
		Password:    hashPwd,
		Email:       input.Email,
	}); err != nil {
		logger.Debugf("err: %v", err)
		return err
	}

	return nil
}

func NewRegisterUseCase(userRepo repository.UserRepo, hasher hash.PasswordHasher) RegisterUseCase {
	return &registerUseCaseImpl{
		userRepo: userRepo,
		hasher:   hasher,
	}
}
