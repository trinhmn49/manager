package persistence

import (
	"context"
	"manager/internal/domain/entity"
)

func (u *userRepoImpl) FindUserByPhone(ctx context.Context, phone string) (*entity.UserEntity, error) {
	var user entity.UserEntity
	result := u.executor.WithContext(ctx).Debug().Where("phone = ?", phone).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
