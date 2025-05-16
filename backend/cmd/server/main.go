package main

import (
	"log"
	"net/http"
	"packs/internal/api"
	"packs/internal/service"
	"packs/internal/utils"
)

func main() {
	logger := utils.NewLogger()

	packConfig, err := utils.LoadPackConfig("./packs.json")
	if err != nil {
		logger.Error("cannot load pack config", "error", err)
		return
	}

	calculator := service.NewPackCalculator(packConfig.Packs, logger)

	httpHandler := api.NewHandler(calculator, logger)

	logger.Info("starting server", "port", 8080)

	if err := http.ListenAndServe(":8080", httpHandler); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
