package main

import (
	"fmt"
	"manager/config"
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

//init db

func main() {
	// init logger
	fmt.Println("Init logger")
	if err := InitLogger(); err != nil {
		fmt.Errorf("Err: %v", err)
		return
	}

	fmt.Println("Hello manager")
}
