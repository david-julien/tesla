package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/davidjulien/tesla/config"
	"github.com/davidjulien/tesla/controller"
	"github.com/davidjulien/tesla/data"
	"github.com/davidjulien/tesla/server"
)

func main() {
	cfg := config.LocalTest()
	dal := data.New(cfg)
	srv := server.New(cfg, dal, log.WithField("service", "server"))
	controller.Init(srv)

	srv.Start()
}
