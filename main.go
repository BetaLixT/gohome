package main

import (
	"fmt"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	err := initLogger()
	if err != nil {
		fmt.Println(err)
		panic("Failed to initialize logger")
	}

	err = loadConfig("dev", "BLT_GOHK", "./configs", false)
	if err != nil {
		fmt.Println(err)
		panic("Failed to load config")
	}
}

func initLogger () error {
	logger, err := zap.NewProduction()
	if err != nil {
		return err
	}
	defer logger.Sync()


	return nil
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

