package main

import (
	"flag"
	"fmt"
	"manager/internal/delivery"
	"manager/internal/persistence"
	"manager/internal/usecase"
	"manager/pkg/config"
	"manager/pkg/database"
	"manager/pkg/hash"
	"manager/pkg/logger"
	"manager/pkg/server"
	"manager/shared/provider"
)

type AppConn struct {
	DB     *database.DbInstance
	Hasher hash.PasswordHasher
}

type Application struct {
	appConn *AppConn
	server  *server.Server
}

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

func buildRoutes(appConn *AppConn) []server.RouteRegister {
	userRepo := persistence.NewUserRepo(appConn.DB.Db)

	loginUseCase := usecase.NewLoginUseCase(userRepo, appConn.Hasher)
	registerUseCase := usecase.NewRegisterUseCase(userRepo, appConn.Hasher)
	return []server.RouteRegister{
		delivery.NewCustomerHandler(loginUseCase, registerUseCase),
	}
}

func (a *Application) Start() {
	routes := buildRoutes(a.appConn)
	a.server.SetupRoute(routes)
	if err := a.server.Start(); err != nil {
		panic(err)
	}
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

	confPath := flag.String("conf", "", "application config path")
	flag.Parse()
	if *confPath == "" {
		*confPath = "conf.env" //default value
	}
	logger.Debugf("Load config from path: %v", *confPath)
	var appConf config.AppConfig
	config.LoadConfig(*confPath, &appConf)
	app := Application{
		appConn: &AppConn{
			DB:     provider.ProvidePostgres(&appConf),
			Hasher: provider.ProvideHasher(),
		},
		server: server.New(appConf),
	}
	app.Start()
}
