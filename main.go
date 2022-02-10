package main

import (
	"fmt"

	"github.com/betalixt/gohome/services"
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
			services.NewLogger,
		),
		fx.Invoke(Start),
		fx.WithLogger(
			func(logger *zap.Logger) fxevent.Logger {
				return &fxevent.ZapLogger{Logger: logger}
			},
		),
	)

	app.Run()
	app.Done()
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
