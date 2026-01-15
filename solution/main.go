package main

import (
	"fmt"
	"log/slog"
	"os"
)

func main() {
	logger := slog.Default()

	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		logger.Error("missed SERVER_PORT env (export smth like '8080')")
		os.Exit(1)
	}

	serverAddress := fmt.Sprintf("0.0.0.0:%s", serverPort)

	s := NewServer(serverAddress, logger)

	err := s.Start()
	if err != nil {
		logger.Error("server has been stopped", "error", err)
	}
}
