package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/Nattakornn/cache/config"
	"github.com/shopspring/decimal"
	"github.com/spf13/viper"
)

var configFile string

func initConfig() {
	// Init viper
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath("./config")
		viper.SetConfigType("yaml")
	}

	// Enable Viper to automatically read env vars
	viper.AutomaticEnv()

	// Read Config
	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintf(os.Stderr, "unable to read config: %v\n", err)
		os.Exit(1)
	}

	// Parse Json Decimal to Float
	decimal.MarshalJSONWithoutQuotes = true

	// config timezone
	timeZone, err := time.LoadLocation(config.LoadConfig().Utils().TimeZone())
	if err != nil {
		panic("Not found timezone: " + config.LoadConfig().Utils().TimeZone() + " please check config SystemConfig.TimeZone\n" + err.Error())
	}
	time.Local = timeZone
}
