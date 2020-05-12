package main

import (
	"fmt"
	"go.uber.org/zap"
	"os"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		fmt.Printf("", "Zap logger setup error")
		fmt.Printf("Application exeution stopped ")

		// Linux and Unix exit code tutorial with examples https://shapeshed.com/unix-exit-codes/
		os.Exit(1)
		return
	}

	logger.Info("Starting golog-zap application.")
	logger.Info("Info level demo log message")
	logger.Info("This log will be shown in the terminal only.")
	logger.Debug("App execution is finished.")
}
