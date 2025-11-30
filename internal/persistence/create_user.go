package persistence

import (
	"context"
	"manager/internal/domain/entity"
)

func (c *userRepoImpl) Create(ctx context.Context, customer entity.UserEntity) error {
	resutl := c.executor.WithContext(ctx).Create(&customer)
	if resutl.Error != nil {
		return resutl.Error
	}

	return nil
}
