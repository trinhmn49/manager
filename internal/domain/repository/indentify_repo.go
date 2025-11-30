package repository

import (
	"context"
	"manager/internal/domain/entity"
)

type UserRepo interface {
	FindUserByPhone(ctx context.Context, phone string) (*entity.UserEntity, error)
	Create(ctx context.Context, customer entity.UserEntity) error
}
