package app

import (
	"github.com/sirupsen/logrus"
	"log"
	"www.github.com/Wolves-from-MISIS/internal/config"
	http_delivery "www.github.com/Wolves-from-MISIS/internal/delivery/http"
	"www.github.com/Wolves-from-MISIS/internal/repository"
	"www.github.com/Wolves-from-MISIS/internal/server"
	"www.github.com/Wolves-from-MISIS/internal/service"
	database "www.github.com/Wolves-from-MISIS/pkg/db/postgres"
)

func Run() {
	// initialize extensions

	cfg, err := config.ReadConfig()
	if err != nil {
		log.Panic(err)
	}

	dbClient, err := database.NewSqlClient(cfg.DB.Dsn)

	repo := repository.NewHackLabRepository(dbClient)

	userService := service.NewUserService(repo)

	handler := http_delivery.NewHandler(userService, nil)

	srv := server.NewServer(cfg, handler.Init(cfg.SwaggerAddr))
	logrus.Println("Server is listening..." + cfg.ServerAddr)
	err = srv.Run()
	if err != nil {
		logrus.Panic("Couldn't run server")
	}
}
