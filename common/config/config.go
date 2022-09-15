package config

import "os"

var workingDir string
var configFile string

func init() {
	workingDir = os.Getenv("COURIER_MANAGEMENT_WORKING_DIR")
	if workingDir == "" {
		workingDir = "."
	}

	configFile = os.Getenv("COURIER_MANAGEMENT_CONFIG_FILE")
	if configFile == "" {
		configFile = "config"
	}
}

func InitConfig() {
	setupViper()
}

func InitTestConfig() {
	setupViper()
}
