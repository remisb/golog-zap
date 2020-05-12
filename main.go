package main

import (
	"fmt"
	"go.uber.org/zap"
	"os"
	"path/filepath"
)

func main() {
	logFileName, err := checkLogFile("test.log")
	if err != nil {
		fmt.Printf("Error on log file create/open %s", err)
		os.Exit(1)
	}

	logConfig :=  zap.NewDevelopmentConfig()
	logConfig.OutputPaths = append(logConfig.OutputPaths, logFileName)
	logger, err := logConfig.Build()
	if err != nil {
		fmt.Printf("", "Zap logger setup error")
		fmt.Printf("Application exeution stopped ")

		os.Exit(1)
		return
	}

	logger.Info("Starting golog-zap application.")
	logger.Info("Info level demo log message")
	logger.Info("This log will be shown in the terminal only.")
	logger.Debug("App execution is finished.")
}

func checkLogFile(logFileName string) (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	fName, err := filepath.Abs(filepath.Join(wd, logFileName))
	if err != nil {
		return "", err
	}

	return fName, nil
}
