package persistence

import (
	"manager/internal/domain/repository"

	"gorm.io/gorm"
)

type userRepoImpl struct {
	executor *gorm.DB
}

func NewUserRepo(gorm *gorm.DB) repository.UserRepo {
	return &userRepoImpl{
		executor: gorm,
	}
}
