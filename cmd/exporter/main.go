package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/hhk7734/custom_exporter.go/internal/pkg/logger"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName(".env")
	viper.SetConfigType("dotenv")

	workDir, _ := os.Getwd()
	for {
		if _, err := os.Stat(filepath.Join(workDir, ".env")); err == nil {
			viper.AddConfigPath(workDir)
			break
		}
		if workDir == "/" {
			break
		}
		workDir = filepath.Dir(workDir)
	}

	if err := viper.ReadInConfig(); err != nil && !errors.As(err, &viper.ConfigFileNotFoundError{}) {
		panic(fmt.Errorf("failed to read config file: %w", err))
	}

	// env
	viper.AutomaticEnv()

	// flag
	pflag.CommandLine.AddFlagSet(logger.LogPFlags())
	pflag.Parse()

	viper.BindPFlags(pflag.CommandLine)

	logger.SetGlobalZapLogger(logger.LogConfigFromViper())
}
