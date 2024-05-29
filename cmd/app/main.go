package main

import (
	"context"
	"os"
	"os/signal"
	"standard/cmd/app/server"
	"standard/internal/config"
	"standard/internal/handler"
	"standard/internal/repository"
	"standard/internal/service"
	"standard/pkg/logger"
	"standard/pkg/utils"
	"syscall"

	"github.com/sirupsen/logrus"
)

// @title Simple App API
// @version 1.0
// @description API Server for Application

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	cfg := config.GetConfig()
	log := logger.GetLogger()

	db, err := utils.SetupPostgresConnection()
	if err != nil {
		log.Fatal(err.Error())
	}

	repos := repository.NewRepository(db, log)
	services := service.NewService(repos, *log)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	go func() {
		if err := srv.Run(cfg.Port, handlers.InitRoutes()); err != nil {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	log.Info("App started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Warn("App shutting down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		log.Errorf("error occured on db connection close: %s", err.Error())
	}
}
