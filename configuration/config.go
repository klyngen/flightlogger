package configuration

import (
	"fmt"

	"github.com/spf13/viper"
)

// Add more good places to hide configs right here ;)
var configpaths = [...]string{"/etc/flightlogger", "$HOME/.flightlogger", "."}

// ApplicationConfig contains the config parameters used for this excellent app
type ApplicationConfig struct {
	Serverport            string
	DatabaseConfiguration DatabaseConfig `mapstructure:"database"`
}

// DatabaseConfig describes parameters used to connect to a database
type DatabaseConfig struct {
	Hostname string
	Password string
	Database string
	Port     string
	Username string
}

// GetConfiguration - well... it gets the configuration
func GetConfiguration() ApplicationConfig {
	initializeConfig()

	var c ApplicationConfig

	viper.Unmarshal(&c)

	return c
}

func initializeConfig() {
	// what is the name of the config-file
	// Can be any of many configuration parameters
	viper.SetConfigName("flightlog")

	// Where can configs be located
	for _, path := range configpaths {
		viper.AddConfigPath(path)
	}

	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("Unable to find flightlog-config in any of "))
	}

}
