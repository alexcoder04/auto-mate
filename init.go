package main

import (
	"os"

	"go.uber.org/zap"
)

func init() {
	var err error

	Logger, err = zap.NewProduction()
	if err != nil {
		println("failed to create logger:", err.Error())
		os.Exit(1)
	}

	defer Logger.Sync()

	Sugar = Logger.Sugar()
}
