package provider

import (
	"manager/pkg/config"
	"manager/pkg/database"
	"manager/pkg/hash"
)

func ProvidePostgres(appConf *config.AppConfig) *database.DbInstance { 
	return database.NewPostgres(appConf)
}

func ProvideHasher() hash.PasswordHasher {
	return hash.NewBcryptPasswordHasher()
}
