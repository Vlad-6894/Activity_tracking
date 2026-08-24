package main

import (
	"fmt"
	"os"

	core_logger "github.com/Vlad-6894/Activity_tracking/internal/core/logger"
)

func main() {

	logger, err := core_logger.NewLogger(core_logger.NewLoggerConfigMust())
	if err != nil {
		fmt.Println("init logger error")
		os.Exit(1)
	}
	defer logger.Close()

}
