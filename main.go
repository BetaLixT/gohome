package main

import (
	"fmt"

	"github.com/spf13/viper"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	err := loadConfig("dev", "BLT_GOHK", "./configs", false)
	if err != nil {
		fmt.Println(err)
		panic(fmt.Sprintf("Failed to load config: %v", err))
	}

	app := fx.New(
		fx.Provide(
			NewLogger,
		),
		fx.Invoke(Start),
		fx.WithLogger(
			func() fxevent.Logger {
				return fxevent.NopLogger
			},
		),
	)

	app.Run()
	app.Done()
}

func NewLogger() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(fmt.Sprintf("Failed to create logger: %v", err))
	}
	return logger
}

func loadConfig(
	env string,
	prfix string,
	loc string,
	fileRequired bool,
) error {
	// TODO Mechanism for defaults

	// - Loading config from file
	viper.SetConfigName(fmt.Sprintf("config.%s", env))
	viper.AddConfigPath(loc)
	err := viper.ReadInConfig()
	if err != nil {
		if fileRequired {
			return err
		}
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	// - Loading config from env
	viper.SetEnvPrefix(prfix)
	viper.AutomaticEnv()
	return nil
}
