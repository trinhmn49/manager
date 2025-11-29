package main

import (
	"fmt"
	"manager/pkg/config"
	"manager/pkg/database"
	"manager/pkg/logger"
)

func InitLogger() error {
	// init logger
	logCfg, err := config.LoadLogConfig()
	if err != nil {
		logger.Error("failed to load logger config", logger.E(err))
		return fmt.Errorf("load log config: %s", err)
	}

	// set logger
	cfg := logger.Config{
		Type:     logCfg.Type,
		Level:    logCfg.Level,
		Output:   logCfg.Output,
		Filename: logCfg.Filename,
		UseJSON:  logCfg.UseJSON,
	}
	logger.SetLogger(cfg)

	logger.Info("init logger successfully", logger.F("cfg", cfg))
	return nil
}

func InitDatabase() error {
	//init db
	dbCfg, err := config.LoadDbConfig()
	if err != nil {
		logger.Error("failed to load database config", logger.E(err))
		return fmt.Errorf("load db config: %s", err)
	}
	logger.Debugf("database config: %v", dbCfg)
	database.NewDb(&config.DbConfig{
		UserName: dbCfg.UserName,
		Password: dbCfg.Password,
		Host:     dbCfg.Host,
		Port:     dbCfg.Port,
		DbName:   dbCfg.DbName,
	})

	logger.Info("init db successfully", logger.F("dbCfg", dbCfg))
	return nil
}

func main() {
	// init logger
	fmt.Println("Init logger")
	if err := InitLogger(); err != nil {
		logger.Errorf("Err: %v", err)
		return
	}

	//init db
	if err := InitDatabase(); err != nil {
		logger.Errorf("Err: %v", err)
		return
	}

	fmt.Println("Hello manager")
}
