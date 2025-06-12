package main

import (
	"tasker/internal/app"
	"tasker/internal/config"
	"tasker/internal/logger"
)

func main() {
	//here you could construct your app
	cfg := config.MustLoad()
	log := logger.NewLogger(cfg.Env)

	ap := app.NewApp(log, cfg)

	ap.MustRun()
}
