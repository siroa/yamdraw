package main

import (
	"generate/app"
	"generate/utils/logger"
)

func main() {
	logger.Info("Starting our application...")
	app.Run()
	logger.Info("finished our application...")
}
