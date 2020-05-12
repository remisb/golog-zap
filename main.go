package main

import (
	"fmt"
	"go.uber.org/zap"
	"os"
	"path/filepath"
)

func main() {
	fname, err := checkLogFile()
	if err != nil {
		fmt.Printf("Error on log file create/open %s", err)
		os.Exit(1)
	}
	logger, err := NewCustomConfig(fname).Build()
	//logger, err := zap.NewDevelopment(zap.Option())
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

func checkLogFile() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	fname := filepath.Join(wd, "test.log")
	fname, err = filepath.Abs(fname)
	if err != nil {
		return "", err
	}

	if _, err := os.Stat(fname); os.IsNotExist(err) {
		// create file
		f, err := os.OpenFile(fname, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
		if err != nil {
			fmt.Printf("error creating file: %v", err)
		}
		defer f.Close()
	}
	return fname, nil
}

func NewCustomConfig(logfname string) zap.Config {
	outputPaths := []string{"stderr", logfname}
	return zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:      true,
		Encoding:         "console",
		EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
		OutputPaths:      outputPaths,
		ErrorOutputPaths: []string{"stderr"},
	}
}
