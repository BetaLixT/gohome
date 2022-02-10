package main

import "go.uber.org/zap"

// Starts up the application
func Start(logger *zap.Logger){
	logger.Info("Booting server...")
}