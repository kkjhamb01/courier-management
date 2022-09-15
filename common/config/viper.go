package config

import (
	"github.com/spf13/viper"
	"log"
	"sync"
)

// to make sure viper would be setup only once
var setupOnce sync.Once

func setupViper() {
	setupOnce.Do(func() {
		viper.AddConfigPath(workingDir)
		viper.SetConfigName(configFile)

		err := viper.ReadInConfig()
		if err != nil {
			log.Fatalf("failed to read the config file: %v", err)
		}

		data = &Data{}
		err = viper.Unmarshal(data)
		if err != nil {
			panic("unable to decode into config struct")
		}
	})
}
