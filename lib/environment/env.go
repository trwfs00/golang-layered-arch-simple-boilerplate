package environment

import (
	"log"

	"github.com/spf13/viper"
)

var (
	EnvKey      = "ENV"
	ServicePort = "SERVICE_PORT"
	DsnKey      = "DB_DSN"
)

func New(dirDepth uint) {
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	var configDir string
	if dirDepth == 0 {
		configDir = "."
	} else {
		configDir = ".."
		for i := uint(1); i < dirDepth; i++ {
			configDir += "/.."
		}
	}

	viper.AddConfigPath(configDir)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
}

func GetString(key string) string {
	if !viper.IsSet(key) {
		panic("failed to get environment key: " + key)
	}

	return viper.GetString(key)
}

func GetInt(key string) int {
	if !viper.IsSet(key) {
		panic("failed to get environment key: " + key)
	}

	return viper.GetInt(key)
}

func GetBool(key string) bool {
	if !viper.IsSet(key) {
		panic("failed to get environment key: " + key)
	}

	return viper.GetBool(key)
}
